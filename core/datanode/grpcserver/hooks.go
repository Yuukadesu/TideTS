package grpcserver

import (
	"time"

	"github.com/hanami/tidets/core/queryengine/plan"
)

// Hooks 允许外层注入 gRPC 与 SQL 早期失败的观测回调。
type Hooks struct {
	OnSQL      func(kind plan.Kind, success bool, errorClass string, duration time.Duration)
	OnRPCItems func(method string, requestItems, responseItems int)
}

// SetHooks 设置服务观测 hooks。
func (s *Server) SetHooks(h Hooks) {
	s.hooks = h
}
