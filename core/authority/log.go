package authority

import (
	"bufio"
	"encoding/binary"
	"io"
	"os"
	"path/filepath"

	commons "github.com/hanami/tidets/commons/errors"
	"github.com/hanami/tidets/core/tsmodel/codec"
)

const (
	authMagic uint32 = 0x54534155 // "TSAU"

	opCreateUser  uint8 = 1
	opSetPassword uint8 = 2
	opGrant       uint8 = 3
	opRevoke      uint8 = 4
)

const authLogName = "authlog.bin"

type authLog struct {
	file *os.File
	w    *bufio.Writer
}

func authLogPath(dir string) string {
	return filepath.Join(dir, authLogName)
}

func openAuthLog(dir string) (*authLog, error) {
	if err := os.MkdirAll(dir, 0o700); err != nil {
		return nil, err
	}
	if err := os.Chmod(dir, 0o700); err != nil {
		return nil, err
	}
	f, err := os.OpenFile(authLogPath(dir), os.O_APPEND|os.O_CREATE|os.O_RDWR, 0o600)
	if err != nil {
		return nil, err
	}
	if err := f.Chmod(0o600); err != nil {
		_ = f.Close()
		return nil, err
	}
	return &authLog{file: f, w: bufio.NewWriter(f)}, nil
}

func (l *authLog) offset() (int64, error) {
	if l == nil || l.file == nil {
		return 0, nil
	}
	if err := l.w.Flush(); err != nil {
		return 0, err
	}
	fi, err := l.file.Stat()
	if err != nil {
		return 0, err
	}
	return fi.Size(), nil
}

func (l *authLog) appendCreate(u User) error {
	if err := writeAuthHeader(l.w, opCreateUser); err != nil {
		return syncErr(err)
	}
	if err := writeUserPayload(l.w, u); err != nil {
		return syncErr(err)
	}
	return l.sync()
}

func (l *authLog) appendPassword(id int64, hash string) error {
	if err := writeAuthHeader(l.w, opSetPassword); err != nil {
		return syncErr(err)
	}
	if err := binary.Write(l.w, binary.LittleEndian, id); err != nil {
		return syncErr(err)
	}
	if err := codec.WriteString(l.w, hash); err != nil {
		return syncErr(err)
	}
	return l.sync()
}

func (l *authLog) appendGrant(op uint8, g Grant) error {
	if err := writeAuthHeader(l.w, op); err != nil {
		return syncErr(err)
	}
	if err := binary.Write(l.w, binary.LittleEndian, g.UserID); err != nil {
		return syncErr(err)
	}
	if err := codec.WriteString(l.w, g.Path); err != nil {
		return syncErr(err)
	}
	if err := binary.Write(l.w, binary.LittleEndian, uint8(g.Privilege)); err != nil {
		return syncErr(err)
	}
	return l.sync()
}

func (l *authLog) sync() error {
	if err := l.w.Flush(); err != nil {
		return syncErr(err)
	}
	if err := l.file.Sync(); err != nil {
		return syncErr(err)
	}
	return nil
}

func (l *authLog) close() error {
	if l == nil {
		return nil
	}
	if l.w != nil {
		if err := l.w.Flush(); err != nil {
			return err
		}
	}
	if l.file != nil {
		return l.file.Close()
	}
	return nil
}

func syncErr(err error) error {
	if err == nil {
		return nil
	}
	if _, ok := commons.As(err); ok {
		return err
	}
	return commons.Wrap("auth", commons.CodeUnavailable, "auth log sync failed", err)
}

func writeAuthHeader(w io.Writer, op uint8) error {
	if err := binary.Write(w, binary.LittleEndian, authMagic); err != nil {
		return err
	}
	return binary.Write(w, binary.LittleEndian, op)
}

func writeUserPayload(w io.Writer, u User) error {
	if err := binary.Write(w, binary.LittleEndian, u.ID); err != nil {
		return err
	}
	if err := codec.WriteString(w, u.Username); err != nil {
		return err
	}
	if err := codec.WriteString(w, u.PasswordHash); err != nil {
		return err
	}
	flag := uint8(0)
	if u.IsSuper {
		flag = 1
	}
	return binary.Write(w, binary.LittleEndian, flag)
}

type replayApplier struct {
	create   func(User) error
	password func(id int64, hash string) error
	grant    func(Grant) error
	revoke   func(Grant) error
}

func replayAuthLog(dir string, from int64, apply replayApplier) error {
	f, err := os.Open(authLogPath(dir))
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		return err
	}
	if from > fi.Size() {
		return commons.ErrAuthLogCorrupt
	}
	if from > 0 {
		if _, err := f.Seek(from, io.SeekStart); err != nil {
			return commons.ErrAuthLogCorrupt
		}
	}
	r := bufio.NewReader(f)
	for {
		op, err := readAuthHeader(r)
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
		switch op {
		case opCreateUser:
			u, err := readUserPayload(r)
			if err != nil {
				return err
			}
			if err := apply.create(u); err != nil {
				return err
			}
		case opSetPassword:
			var id int64
			if err := binary.Read(r, binary.LittleEndian, &id); err != nil {
				return corruptRead(err)
			}
			hash, err := codec.ReadString(r)
			if err != nil {
				return corruptRead(err)
			}
			if err := apply.password(id, hash); err != nil {
				return err
			}
		case opGrant, opRevoke:
			g, err := readGrantPayload(r)
			if err != nil {
				return err
			}
			if op == opGrant {
				if err := apply.grant(g); err != nil {
					return err
				}
			} else if err := apply.revoke(g); err != nil {
				return err
			}
		default:
			return commons.ErrAuthLogCorrupt
		}
	}
}

func readAuthHeader(r io.Reader) (uint8, error) {
	var magic uint32
	if err := binary.Read(r, binary.LittleEndian, &magic); err != nil {
		if err == io.EOF {
			return 0, io.EOF
		}
		return 0, commons.ErrAuthLogCorrupt
	}
	if magic != authMagic {
		return 0, commons.ErrAuthLogCorrupt
	}
	var op uint8
	if err := binary.Read(r, binary.LittleEndian, &op); err != nil {
		return 0, commons.ErrAuthLogCorrupt
	}
	return op, nil
}

func readUserPayload(r io.Reader) (User, error) {
	var id int64
	if err := binary.Read(r, binary.LittleEndian, &id); err != nil {
		return User{}, corruptRead(err)
	}
	name, err := codec.ReadString(r)
	if err != nil {
		return User{}, corruptRead(err)
	}
	hash, err := codec.ReadString(r)
	if err != nil {
		return User{}, corruptRead(err)
	}
	var flag uint8
	if err := binary.Read(r, binary.LittleEndian, &flag); err != nil {
		return User{}, corruptRead(err)
	}
	if flag > 1 {
		return User{}, commons.ErrAuthLogCorrupt
	}
	return User{ID: id, Username: name, PasswordHash: hash, IsSuper: flag == 1}, nil
}

func readGrantPayload(r io.Reader) (Grant, error) {
	var id int64
	if err := binary.Read(r, binary.LittleEndian, &id); err != nil {
		return Grant{}, corruptRead(err)
	}
	path, err := codec.ReadString(r)
	if err != nil {
		return Grant{}, corruptRead(err)
	}
	var priv uint8
	if err := binary.Read(r, binary.LittleEndian, &priv); err != nil {
		return Grant{}, corruptRead(err)
	}
	if priv != uint8(PrivilegeRead) && priv != uint8(PrivilegeWrite) {
		return Grant{}, commons.ErrAuthLogCorrupt
	}
	return Grant{UserID: id, Path: path, Privilege: Privilege(priv)}, nil
}

func corruptRead(err error) error {
	if err == nil {
		return nil
	}
	if _, ok := commons.As(err); ok {
		return err
	}
	return commons.ErrAuthLogCorrupt
}
