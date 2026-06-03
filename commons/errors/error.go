// Package commons 提供跨模块共用的错误类型（对齐 IoTDB commons 子集）。
package commons

import (
	"errors"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Code 错误分类，用于映射 gRPC status 与 errors.Is。
type Code int

const (
	CodeUnknown Code = iota
	CodeInvalidArgument
	CodeNotFound
	CodePermissionDenied
	CodeUnauthenticated
	CodeCorrupt
	CodeUnavailable
	CodeInternal
)

// Error 项目统一错误类型。
type Error struct {
	Code  Code
	Op    string // 子系统，如 storage、auth、segment
	Msg   string
	Cause error
}

// New 构造无底层 cause 的错误。
func New(op string, code Code, msg string) *Error {
	return &Error{Code: code, Op: op, Msg: msg}
}

// Wrap 包装底层错误。
func Wrap(op string, code Code, msg string, cause error) *Error {
	return &Error{Code: code, Op: op, Msg: msg, Cause: cause}
}

// Errorf 构造带格式化消息的错误。
func Errorf(op string, code Code, format string, args ...any) *Error {
	return &Error{Code: code, Op: op, Msg: fmt.Sprintf(format, args...)}
}

func (e *Error) Error() string {
	if e == nil {
		return ""
	}
	if e.Op != "" {
		if e.Cause != nil {
			return fmt.Sprintf("%s: %s: %v", e.Op, e.Msg, e.Cause)
		}
		return fmt.Sprintf("%s: %s", e.Op, e.Msg)
	}
	if e.Cause != nil {
		return fmt.Sprintf("%s: %v", e.Msg, e.Cause)
	}
	return e.Msg
}

func (e *Error) Unwrap() error { return e.Cause }

// Is 判断 err 是否为目标 *Error 且 Code 一致（也支持 errors.Is 到同一哨兵指针）。
func Is(err error, target *Error) bool {
	if target == nil {
		return err == nil
	}
	var e *Error
	if errors.As(err, &e) {
		return e.Code == target.Code && e.Op == target.Op && e.Msg == target.Msg
	}
	return errors.Is(err, target)
}

// As 提取 *Error。
func As(err error) (*Error, bool) {
	var e *Error
	if errors.As(err, &e) {
		return e, true
	}
	return nil, false
}

func (e *Error) GRPCCode() codes.Code {
	switch e.Code {
	case CodeInvalidArgument:
		return codes.InvalidArgument
	case CodeNotFound:
		return codes.NotFound
	case CodePermissionDenied:
		return codes.PermissionDenied
	case CodeUnauthenticated:
		return codes.Unauthenticated
	case CodeUnavailable:
		return codes.Unavailable
	case CodeCorrupt, CodeInternal, CodeUnknown:
		return codes.Internal
	default:
		return codes.Internal
	}
}

// ToGRPCStatus 将 error 转为 gRPC status；*Error 保留分类，其它视为 Internal。
func ToGRPCStatus(err error) error {
	if err == nil {
		return nil
	}
	if e, ok := As(err); ok {
		return status.Error(e.GRPCCode(), e.Error())
	}
	return status.Errorf(codes.Internal, "%v", err)
}
