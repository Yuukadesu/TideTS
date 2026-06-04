package auth

import (
	"github.com/hanami/tidets/commons/errors"
)

const (
	RootUsername = "root"
	RootUserID   = int64(0)
)

// Privilege 数据操作权限（后续与用户-路径授权表对齐）。
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

// Checker 鉴权中心：登录 + 路径权限（对齐 IoTDB AuthorityChecker 子集）。
type Checker struct{}

// DefaultChecker 默认鉴权器（登录 + 路径权限）。
func DefaultChecker() Checker {
	return Checker{}
}

// Authenticate 校验用户名密码，返回 userID。
func (Checker) Authenticate(username, password string) (int64, bool, error) {
	if username == "" || password == "" {
		return -1, false, nil
	}
	if username == RootUsername && password == RootPassword() {
		return RootUserID, true, nil
	}
	return -1, false, nil
}

// RootPassword 默认 root 密码，与客户端 session 默认配置一致。
func RootPassword() string { return "root" }

// CheckPrivilege 校验用户对路径的操作权限；root 放行，其它用户待接元数据。
func (Checker) CheckPrivilege(userID int64, username, path string, priv Privilege) error {
	if userID == RootUserID && username == RootUsername {
		return nil
	}
	if path == "" {
		return ErrInvalidPath
	}
	_ = priv
	return ErrPermissionDenied
}
