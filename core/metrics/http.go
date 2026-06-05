package metrics

import "net/http"

// NewHTTPServer 创建暴露 /metrics 的 HTTP 服务。
func NewHTTPServer(addr string, reg *Registry) *http.Server {
	mux := http.NewServeMux()
	if reg != nil {
		mux.Handle("/metrics", reg.Handler())
	}
	return &http.Server{
		Addr:    addr,
		Handler: mux,
	}
}
