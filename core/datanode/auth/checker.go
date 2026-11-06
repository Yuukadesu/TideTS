package auth

import (
	"github.com/hanami/tidets/commons/errors"
	"github.com/hanami/tidets/core/authority"
)

const (
	RootUsername = "root"
	// RootUserID 是接入持久化账户之前的内存 root id。持久化 root 的 id 由账户日志分配，从 1 开始。
	RootUserID = int64(0)
)

// Privilege 数据操作权限。
type Privilege int

const (
	PrivilegeRead Privilege = iota
	PrivilegeWrite
)

// 对外别名，便于 auth 包调用方使用 errors.Is。
var (
	ErrInvalidPath      = commons.ErrAuthInvalidPath
	ErrPermissionDenied = commons.ErrAuthPermissionDenied
)

// Checker 鉴权中心：登录 + 路径权限。账户数据来自 authority.Service。
type Checker struct {
	accounts *authority.Service
}

// NewChecker 绑定已打开的账户存储。
func NewChecker(accounts *authority.Service) Checker {
	return Checker{accounts: accounts}
}

// DefaultChecker 没有账户存储的鉴权器。登录失败，路径访问拒绝。
func DefaultChecker() Checker {
	return Checker{}
}

// Authenticate 校验用户名密码，返回 userID。
func (c Checker) Authenticate(username, password string) (int64, bool, error) {
	if c.accounts == nil {
		return -1, false, nil
	}
	return c.accounts.Authenticate(username, password)
}

// RootPassword 默认 root 密码，与客户端 session 默认配置一致。
func RootPassword() string { return "root" }

// CheckPrivilege 校验用户对路径的操作权限。超级用户放行，其它用户按路径授权。
func (c Checker) CheckPrivilege(userID int64, username, path string, priv Privilege) error {
	_ = username
	if c.accounts == nil {
		return ErrPermissionDenied
	}
	mapped, ok := mapPrivilege(priv)
	if !ok {
		return ErrPermissionDenied
	}
	return c.accounts.CheckPrivilege(userID, path, mapped)
}

// IsSuper 报告用户是否为超级用户。
func (c Checker) IsSuper(userID int64) bool {
	return c.accounts != nil && c.accounts.IsSuper(userID)
}

func mapPrivilege(priv Privilege) (authority.Privilege, bool) {
	switch priv {
	case PrivilegeRead:
		return authority.PrivilegeRead, true
	case PrivilegeWrite:
		return authority.PrivilegeWrite, true
	default:
		return 0, false
	}
}
