package authority

import (
	"os"
	"path/filepath"
	"sync"

	commons "github.com/hanami/tidets/commons/errors"
	"github.com/hanami/tidets/core/tsmodel"
)

const snapshotEveryOps = 100

// Service 账户与路径授权。热路径只读内存，变更先 sync 日志再改内存。
type Service struct {
	mu     sync.RWMutex
	dir    string
	log    *authLog
	users  map[string]*User
	byID   map[int64]*User
	grants map[int64][]Grant
	nextID int64
	ops    int
}

// Open 从 dataDir/system/auth 加载账户。目录为空时种下 root。
func Open(dataDir string) (*Service, error) {
	dir := filepath.Join(dataDir, "system", "auth")
	if err := os.MkdirAll(dir, 0o700); err != nil {
		return nil, err
	}
	if err := os.Chmod(dir, 0o700); err != nil {
		return nil, err
	}

	snap, err := loadSnapshot(dir)
	if err != nil {
		return nil, err
	}
	logExists, logSize, err := fileSize(authLogPath(dir))
	if err != nil {
		return nil, err
	}
	if snap == nil && !logExists {
		return seed(dir)
	}

	s := newService(dir)
	from := int64(0)
	if snap != nil {
		if err := s.installSnapshot(snap); err != nil {
			return nil, err
		}
		from = snap.Offset
	}
	if from > 0 && !logExists {
		return nil, commons.ErrAuthLogCorrupt
	}
	if logExists && from > logSize {
		return nil, commons.ErrAuthLogCorrupt
	}
	lg, err := openAuthLog(dir)
	if err != nil {
		return nil, err
	}
	s.log = lg
	if logExists && logSize > from {
		if err := replayAuthLog(dir, from, s.replayer()); err != nil {
			_ = lg.close()
			return nil, err
		}
	}
	if !s.hasSuper() {
		_ = lg.close()
		return nil, commons.ErrAuthNoSuperuser
	}
	return s, nil
}

func seed(dir string) (*Service, error) {
	hash, err := hashPassword(rootPassword)
	if err != nil {
		return nil, err
	}
	lg, err := openAuthLog(dir)
	if err != nil {
		return nil, err
	}
	root := User{ID: 1, Username: rootUsername, PasswordHash: hash, IsSuper: true}
	if err := lg.appendCreate(root); err != nil {
		_ = lg.close()
		return nil, err
	}
	s := newService(dir)
	s.log = lg
	s.addUser(root)
	return s, nil
}

func newService(dir string) *Service {
	return &Service{
		dir:    dir,
		users:  make(map[string]*User),
		byID:   make(map[int64]*User),
		grants: make(map[int64][]Grant),
		nextID: 1,
	}
}

// Close 写快照并关闭日志。
func (s *Service) Close() error {
	if s == nil {
		return nil
	}
	payload, err := s.snapshotPayload()
	if err != nil {
		return err
	}
	if err := persistSnapshot(s.dir, payload); err != nil {
		return err
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.log.close()
}

// Authenticate 校验密码。未知用户、空密码和错误密码都返回 ok=false。
func (s *Service) Authenticate(username, password string) (int64, bool, error) {
	if s == nil || username == "" || password == "" {
		return -1, false, nil
	}
	s.mu.RLock()
	u := s.users[username]
	var id int64
	var hash string
	if u != nil {
		id = u.ID
		hash = u.PasswordHash
	}
	s.mu.RUnlock()
	if hash == "" || !passwordMatches(hash, password) {
		return -1, false, nil
	}
	return id, true, nil
}

// CreateUser 创建非超级用户。
func (s *Service) CreateUser(username, password string) error {
	if err := validateUsername(username); err != nil {
		return err
	}
	if username == rootUsername {
		return commons.ErrAuthReservedUser
	}
	hash, err := hashPassword(password)
	if err != nil {
		return err
	}
	return s.mutate(func() error {
		if _, ok := s.users[username]; ok {
			return commons.ErrAuthUserExists
		}
		u := User{ID: s.nextID, Username: username, PasswordHash: hash, IsSuper: false}
		if err := s.log.appendCreate(u); err != nil {
			return err
		}
		s.addUser(u)
		return nil
	})
}

// SetPassword 修改已有用户的密码，包括 root。
func (s *Service) SetPassword(username, password string) error {
	if err := validateUsername(username); err != nil {
		return err
	}
	hash, err := hashPassword(password)
	if err != nil {
		return err
	}
	return s.mutate(func() error {
		u := s.users[username]
		if u == nil {
			return commons.ErrAuthUserNotFound
		}
		if err := s.log.appendPassword(u.ID, hash); err != nil {
			return err
		}
		u.PasswordHash = hash
		return nil
	})
}

// Grant 授予路径权限。已存在的组合不追加日志。
func (s *Service) Grant(username, path string, privs ...Privilege) error {
	if err := validateUsername(username); err != nil {
		return err
	}
	if err := validateGrantPath(path); err != nil {
		return err
	}
	if len(privs) == 0 {
		return commons.ErrAuthInvalidGrantPath
	}
	for _, priv := range privs {
		if priv != PrivilegeRead && priv != PrivilegeWrite {
			return commons.ErrAuthInvalidGrantPath
		}
	}
	return s.mutate(func() error {
		u := s.users[username]
		if u == nil {
			return commons.ErrAuthUserNotFound
		}
		for _, priv := range privs {
			if s.hasGrant(u.ID, path, priv) {
				continue
			}
			g := Grant{UserID: u.ID, Path: path, Privilege: priv}
			if err := s.log.appendGrant(opGrant, g); err != nil {
				return err
			}
			s.grants[u.ID] = append(s.grants[u.ID], g)
		}
		return nil
	})
}

// Revoke 回收路径权限。任一授权不存在则整批不写日志。
func (s *Service) Revoke(username, path string, privs ...Privilege) error {
	if err := validateUsername(username); err != nil {
		return err
	}
	if err := validateGrantPath(path); err != nil {
		return err
	}
	if len(privs) == 0 {
		return commons.ErrAuthGrantNotFound
	}
	return s.mutate(func() error {
		u := s.users[username]
		if u == nil {
			return commons.ErrAuthUserNotFound
		}
		for _, priv := range privs {
			if !s.hasGrant(u.ID, path, priv) {
				return commons.ErrAuthGrantNotFound
			}
		}
		for _, priv := range privs {
			g := Grant{UserID: u.ID, Path: path, Privilege: priv}
			if err := s.log.appendGrant(opRevoke, g); err != nil {
				return err
			}
			s.removeGrant(u.ID, path, priv)
		}
		return nil
	})
}

// CheckPrivilege 按内存快照判断路径权限。超级用户放行，包括空路径。
func (s *Service) CheckPrivilege(userID int64, path string, priv Privilege) error {
	if s == nil {
		return commons.ErrAuthPermissionDenied
	}
	s.mu.RLock()
	defer s.mu.RUnlock()
	u, ok := s.byID[userID]
	if !ok {
		return commons.ErrAuthPermissionDenied
	}
	if u.IsSuper {
		return nil
	}
	if path == "" {
		return commons.ErrAuthInvalidPath
	}
	for _, g := range s.grants[userID] {
		if g.Privilege == priv && tsmodel.MatchDevicePattern(path, g.Path) {
			return nil
		}
	}
	return commons.ErrAuthPermissionDenied
}

// IsSuper 报告 userID 是否为超级用户。
func (s *Service) IsSuper(userID int64) bool {
	if s == nil {
		return false
	}
	s.mu.RLock()
	defer s.mu.RUnlock()
	u, ok := s.byID[userID]
	return ok && u.IsSuper
}

func (s *Service) mutate(fn func() error) error {
	s.mu.Lock()
	err := fn()
	var payload []byte
	snap := false
	if err == nil {
		s.ops++
		if s.ops >= snapshotEveryOps {
			payload, err = s.encodeLocked()
			if err == nil {
				snap = true
				s.ops = 0
			}
		}
	}
	s.mu.Unlock()
	if err != nil {
		return err
	}
	if !snap {
		return nil
	}
	if werr := persistSnapshot(s.dir, payload); werr != nil {
		s.mu.Lock()
		s.ops = snapshotEveryOps
		s.mu.Unlock()
		return werr
	}
	return nil
}

func (s *Service) snapshotPayload() ([]byte, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.encodeLocked()
}

func (s *Service) encodeLocked() ([]byte, error) {
	off, err := s.log.offset()
	if err != nil {
		return nil, err
	}
	st := snapshotState{Offset: off, Users: make([]User, 0, len(s.byID)), Grants: make([]Grant, 0)}
	for _, u := range s.byID {
		st.Users = append(st.Users, *u)
	}
	for _, list := range s.grants {
		st.Grants = append(st.Grants, list...)
	}
	return encodeSnapshot(st)
}

func (s *Service) installSnapshot(snap *snapshotState) error {
	for _, u := range snap.Users {
		if _, ok := s.users[u.Username]; ok {
			return commons.ErrAuthSnapshotCorrupt
		}
		if _, ok := s.byID[u.ID]; ok {
			return commons.ErrAuthSnapshotCorrupt
		}
		cp := u
		s.addUser(cp)
	}
	for _, g := range snap.Grants {
		if s.byID[g.UserID] == nil || s.hasGrant(g.UserID, g.Path, g.Privilege) {
			return commons.ErrAuthSnapshotCorrupt
		}
		s.grants[g.UserID] = append(s.grants[g.UserID], g)
	}
	return nil
}

func (s *Service) replayer() replayApplier {
	return replayApplier{
		create: func(u User) error {
			if _, ok := s.users[u.Username]; ok {
				return commons.ErrAuthLogCorrupt
			}
			if _, ok := s.byID[u.ID]; ok {
				return commons.ErrAuthLogCorrupt
			}
			s.addUser(u)
			return nil
		},
		password: func(id int64, hash string) error {
			u := s.byID[id]
			if u == nil || hash == "" {
				return commons.ErrAuthLogCorrupt
			}
			u.PasswordHash = hash
			return nil
		},
		grant: func(g Grant) error {
			if s.byID[g.UserID] == nil || s.hasGrant(g.UserID, g.Path, g.Privilege) {
				return commons.ErrAuthLogCorrupt
			}
			s.grants[g.UserID] = append(s.grants[g.UserID], g)
			return nil
		},
		revoke: func(g Grant) error {
			if !s.hasGrant(g.UserID, g.Path, g.Privilege) {
				return commons.ErrAuthLogCorrupt
			}
			s.removeGrant(g.UserID, g.Path, g.Privilege)
			return nil
		},
	}
}

func (s *Service) addUser(u User) {
	cp := u
	s.users[cp.Username] = &cp
	s.byID[cp.ID] = &cp
	if cp.ID >= s.nextID {
		s.nextID = cp.ID + 1
	}
}

func (s *Service) hasGrant(userID int64, path string, priv Privilege) bool {
	for _, g := range s.grants[userID] {
		if g.same(userID, path, priv) {
			return true
		}
	}
	return false
}

func (s *Service) removeGrant(userID int64, path string, priv Privilege) {
	list := s.grants[userID]
	out := list[:0]
	for _, g := range list {
		if g.same(userID, path, priv) {
			continue
		}
		out = append(out, g)
	}
	s.grants[userID] = out
}

func (s *Service) hasSuper() bool {
	for _, u := range s.byID {
		if u.IsSuper {
			return true
		}
	}
	return false
}

func fileSize(path string) (exists bool, size int64, err error) {
	fi, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, 0, nil
		}
		return false, 0, err
	}
	return true, fi.Size(), nil
}
