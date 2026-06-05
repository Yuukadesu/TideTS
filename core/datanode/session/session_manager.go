package session

import (
	"sync"
	"sync/atomic"

	"github.com/hanami/tidets/commons/errors"
	"github.com/hanami/tidets/core/datanode/auth"
)

// Manager 管理 DataNode 上所有活跃会话（对齐 IoTDB SessionManager 子集）。
type Manager struct {
	mu            sync.RWMutex
	sessions      map[int64]IClientSession
	nextSessionID int64

	auth auth.Authenticator
}

// NewManager 创建会话管理器，auth 为空时使用 datanode/auth 默认实现。
func NewManager(authenticator auth.Authenticator) *Manager {
	if authenticator == nil {
		authenticator = auth.DefaultAuthenticator()
	}
	return &Manager{
		sessions: make(map[int64]IClientSession),
		auth:     authenticator,
	}
}

// Register 在 login 前登记连接（对齐 registerSession）。
// gRPC 单连接模型下，失败登录的会话不会进入 sessions 映射。
func (m *Manager) Register(_ IClientSession) bool {
	return true
}

// SupplySession login 成功后填充会话身份（对齐 supplySession）。
func (m *Manager) SupplySession(sess IClientSession, userID int64, username, zone string, version ClientVersion, fetchSize int64) int64 {
	id := atomic.AddInt64(&m.nextSessionID, 1)
	sess.SetID(id)
	sess.SetUserID(userID)
	sess.SetUsername(username)
	if zone != "" {
		sess.SetZone(zone)
	}
	if version != "" {
		sess.SetClientVersion(version)
	}
	if fetchSize > 0 {
		sess.SetFetchSize(fetchSize)
	}
	sess.SetLogInTime(nowMillis())
	sess.SetLogin(true)
	sess.TouchActiveTime()

	m.mu.Lock()
	m.sessions[id] = sess
	m.mu.Unlock()
	return id
}

// Login 认证并开通会话。
func (m *Manager) Login(sess IClientSession, params LoginParams) LoginResult {
	userID, ok, err := m.auth.Authenticate(params.Username, params.Password)
	if err != nil {
		return LoginResult{OK: false, Message: err.Error()}
	}
	if !ok {
		return LoginResult{OK: false, Message: "invalid username or password"}
	}

	version := params.ClientVersion
	if version == "" {
		version = ClientVersionV10
	}
	sess.SetEnableRPCCompression(params.EnableRPCCompression)
	if params.FetchSize > 0 {
		sess.SetFetchSize(params.FetchSize)
	}

	id := m.SupplySession(sess, userID, params.Username, params.ZoneID, version, params.FetchSize)
	return LoginResult{
		SessionID: id,
		UserID:    userID,
		OK:        true,
		Message:   "login successfully",
	}
}

// OpenGRPCSession 创建 ClientSession 并完成 login，供 OpenSession RPC 使用。
func (m *Manager) OpenGRPCSession(params OpenParams) (IClientSession, LoginResult) {
	sess := NewClientSession(params.ClientAddr, params.ClientPort)
	if !m.Register(sess) {
		return sess, LoginResult{OK: false, Message: "session already registered on this connection"}
	}
	result := m.Login(sess, params.Login)
	if !result.OK {
		return sess, result
	}
	return sess, result
}

// Get 按 sessionID 获取已登录会话。
func (m *Manager) Get(sessionID int64) (IClientSession, bool) {
	if sessionID <= 0 {
		return nil, false
	}
	m.mu.RLock()
	sess, ok := m.sessions[sessionID]
	m.mu.RUnlock()
	if !ok || !sess.IsLogin() {
		return nil, false
	}
	return sess, true
}

// Require 获取会话并刷新活跃时间，供数据面 RPC 使用。
func (m *Manager) Require(sessionID int64) (IClientSession, error) {
	sess, ok := m.Get(sessionID)
	if !ok {
		return nil, commons.ErrSessionNotFound
	}
	sess.TouchActiveTime()
	return sess, nil
}

// CloseSession 关闭并移除会话（对齐 closeSession 资源释放子集）。
func (m *Manager) CloseSession(sessionID int64) error {
	m.mu.Lock()
	sess, ok := m.sessions[sessionID]
	if ok {
		delete(m.sessions, sessionID)
	}
	m.mu.Unlock()
	if !ok {
		return commons.ErrSessionNotFound
	}
	sess.SetLogin(false)
	return nil
}

// SessionInfo 构造执行层 SessionInfo。
func (m *Manager) SessionInfo(sessionID int64) (SessionInfo, bool) {
	sess, ok := m.Get(sessionID)
	if !ok {
		return SessionInfo{}, false
	}
	return sess.SessionInfo(), true
}

// ActiveCount 当前活跃会话数。
func (m *Manager) ActiveCount() int {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return len(m.sessions)
}

// Count 当前活跃会话数。
func (m *Manager) Count() int {
	return m.ActiveCount()
}
