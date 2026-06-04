package queryengine

import (
	"context"

	"github.com/hanami/tidets/core/dataplane"
	"github.com/hanami/tidets/core/queryengine/backend"
	"github.com/hanami/tidets/core/queryengine/executor"
	"github.com/hanami/tidets/core/queryengine/plan"
	"github.com/hanami/tidets/core/queryengine/result"
	sqlparser "github.com/hanami/tidets/core/sql/parser"
	"github.com/hanami/tidets/core/storageengine"
	"github.com/hanami/tidets/core/tsmodel"
)

// Service 解析 SQL → 计划 → 执行（INSERT / SELECT）。
type Service struct {
	exec *executor.Executor
}

// NewService 创建 SQL 查询服务。
func NewService(engine *storageengine.Engine, gw *dataplane.Gateway, catalog backend.CatalogBackend) *Service {
	return &Service{
		exec: &executor.Executor{
			Store:        &backend.EngineBackend{Writer: gw, Reader: engine},
			Catalog:      catalog,
			DefaultLimit: executor.DefaultQueryLimit,
		},
	}
}

// NewServiceWithBackend 注入自定义 Backend（测试用）。
func NewServiceWithBackend(store backend.Backend, catalog backend.CatalogBackend, defaultLimit int) *Service {
	if defaultLimit <= 0 {
		defaultLimit = executor.DefaultQueryLimit
	}
	return &Service{
		exec: &executor.Executor{Store: store, Catalog: catalog, DefaultLimit: defaultLimit},
	}
}

// Plan 解析并生成执行计划（不访问存储）。
func Plan(sqlText string) (plan.Plan, error) {
	stmt, err := sqlparser.Parse(sqlText)
	if err != nil {
		return plan.Plan{}, err
	}
	return plan.Build(stmt)
}

// Execute 执行一条 SQL。
func (s *Service) Execute(ctx context.Context, sqlText string) (*result.Result, error) {
	p, err := Plan(sqlText)
	if err != nil {
		return nil, err
	}
	return s.exec.Execute(ctx, p)
}

// ExecutePlan 执行已有计划。
func (s *Service) ExecutePlan(ctx context.Context, p plan.Plan) (*result.Result, error) {
	return s.exec.Execute(ctx, p)
}

// QueryRange 统一范围查询入口，供非 SQL RPC 复用 queryengine。
func (s *Service) QueryRange(ctx context.Context, key tsmodel.SeriesKey, start, end int64, limit int) (*result.Result, error) {
	return s.exec.Execute(ctx, plan.Plan{
		Kind: plan.KindSelect,
		Select: &plan.Select{
			Key:   key,
			Start: start,
			End:   end,
			Limit: limit,
		},
	})
}

// ResolveQueryLimit 合并请求 limit、session fetchSize 与默认值。
func ResolveQueryLimit(requestLimit, sessionFetchSize int) int {
	limit := requestLimit
	if limit <= 0 {
		limit = sessionFetchSize
	}
	if limit <= 0 {
		limit = executor.DefaultQueryLimit
	}
	return limit
}
