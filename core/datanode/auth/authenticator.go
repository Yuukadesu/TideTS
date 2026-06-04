package auth

// Authenticator 登录认证（对齐 IoTDB AuthorityChecker.checkUser 入口）。
type Authenticator interface {
	Authenticate(username, password string) (userID int64, ok bool, err error)
}

// DefaultAuthenticator 默认认证器，供 DataNode SessionManager 使用。
func DefaultAuthenticator() Authenticator {
	c := DefaultChecker()
	return &c
}
