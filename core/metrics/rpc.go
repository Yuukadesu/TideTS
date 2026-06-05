package metrics

import (
	"context"
	"strings"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

// UnaryServerInterceptor 记录 gRPC 请求总数与耗时。
func (r *Registry) UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	if r == nil {
		return func(ctx context.Context, req any, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
			return handler(ctx, req)
		}
	}
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
		start := time.Now()
		resp, err := handler(ctx, req)
		r.ObserveRPC(info.FullMethod, status.Code(err), time.Since(start))
		return resp, err
	}
}

func trimMethod(fullMethod string) string {
	if idx := strings.LastIndex(fullMethod, "/"); idx >= 0 && idx+1 < len(fullMethod) {
		return fullMethod[idx+1:]
	}
	return fullMethod
}
