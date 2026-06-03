package session

import "fmt"

// IClientSession 客户端会话抽象（对齐 IoTDB IClientSession）。
// 一条连接对应一个会话身份，后续权限校验依赖 UserID / Username。
type IClientSession interface {
	ID() int64
	SetID(id int64)

	UserID() int64
	SetUserID(userID int64)

	Username() string
	SetUsername(username string)

	IsLogin() bool
	SetLogin(login bool)

	LogInTime() int64
	SetLogInTime(ms int64)

	LastActiveTime() int64
	TouchActiveTime()

	ClientVersion() ClientVersion
	SetClientVersion(v ClientVersion)

	Zone() string
	SetZone(zone string)

	FetchSize() int64
	SetFetchSize(n int64)

	EnableRPCCompression() bool
	SetEnableRPCCompression(enable bool)

	ClientAddress() string
	ClientPort() int
	ConnectionType() ConnectionType
	ConnectionID() string

	String() string
	SessionInfo() SessionInfo
}

// baseSession IClientSession 公共字段（对齐 IoTDB IClientSession 成员）。
type baseSession struct {
	id                   int64
	userID               int64
	username             string
	login                bool
	logInTime            int64
	lastActiveTime       int64
	clientVersion        ClientVersion
	zone                 string
	fetchSize            int64
	enableRPCCompression bool
}

func (s *baseSession) ID() int64                        { return s.id }
func (s *baseSession) SetID(id int64)                   { s.id = id }
func (s *baseSession) UserID() int64                    { return s.userID }
func (s *baseSession) SetUserID(userID int64)           { s.userID = userID }
func (s *baseSession) Username() string                 { return s.username }
func (s *baseSession) SetUsername(username string)      { s.username = username }
func (s *baseSession) IsLogin() bool                    { return s.login }
func (s *baseSession) SetLogin(login bool)              { s.login = login }
func (s *baseSession) LogInTime() int64                 { return s.logInTime }
func (s *baseSession) SetLogInTime(ms int64)            { s.logInTime = ms }
func (s *baseSession) LastActiveTime() int64            { return s.lastActiveTime }
func (s *baseSession) TouchActiveTime()                 { s.lastActiveTime = nowMillis() }
func (s *baseSession) ClientVersion() ClientVersion     { return s.clientVersion }
func (s *baseSession) SetClientVersion(v ClientVersion) { s.clientVersion = v }
func (s *baseSession) Zone() string                     { return s.zone }
func (s *baseSession) SetZone(zone string)              { s.zone = zone }
func (s *baseSession) FetchSize() int64                 { return s.fetchSize }
func (s *baseSession) SetFetchSize(n int64)             { s.fetchSize = n }
func (s *baseSession) EnableRPCCompression() bool       { return s.enableRPCCompression }
func (s *baseSession) SetEnableRPCCompression(e bool)   { s.enableRPCCompression = e }

func sessionInfoFrom(s IClientSession) SessionInfo {
	return SessionInfo{
		SessionID: s.ID(),
		User: UserEntity{
			UserID:        s.UserID(),
			Username:      s.Username(),
			ClientAddress: s.ClientAddress(),
		},
		Zone:      s.Zone(),
		Version:   s.ClientVersion(),
		FetchSize: s.FetchSize(),
	}
}

func formatSessionString(id int64, username, connectionID string) string {
	return fmt.Sprintf("%d-%s:%s", id, username, connectionID)
}
