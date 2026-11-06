package authority

// Privilege 路径授权类型。0 保留为非法值，避免把未初始化记录当成 READ。
type Privilege uint8

const (
	PrivilegeRead Privilege = 1 + iota
	PrivilegeWrite
)

const (
	rootUsername = "root"
	rootPassword = "root"
)

// User 内存中的账户。密码只保留 bcrypt 哈希。
type User struct {
	ID           int64
	Username     string
	PasswordHash string
	IsSuper      bool
}

// Grant 一条路径授权。
type Grant struct {
	UserID    int64
	Path      string
	Privilege Privilege
}

func (g Grant) same(userID int64, path string, priv Privilege) bool {
	return g.UserID == userID && g.Path == path && g.Privilege == priv
}
