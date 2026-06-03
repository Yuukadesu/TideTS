package queryengine

import (
	"context"
	"github.com/hanami/tidets/core/queryengine/executor"
	"github.com/hanami/tidets/core/queryengine/plan"
	"github.com/hanami/tidets/core/queryengine/result"
	querystore "github.com/hanami/tidets/core/queryengine/storage"
	sqlparser "github.com/hanami/tidets/core/sql/parser"
	"github.com/hanami/tidets/core/sql/planner"
)

// Service 解析 SQL → 计划 → 执行（INSERT / SELECT）。
type Service struct {
	exec *executor.Executor
}

// NewService 创建 SQL 查询服务。
func NewService(store querystore.Backend) *Service {
	return &Service{
		exec: &executor.Executor{Store: store, DefaultLimit: 10000},
	}
}

// NewServiceWithLimit 指定 SELECT 默认 limit（SQL 未写 LIMIT 时使用）。
func NewServiceWithLimit(store querystore.Backend, defaultLimit int) *Service {
	return &Service{
		exec: &executor.Executor{Store: store, DefaultLimit: defaultLimit},
	}
}

// Plan 解析并生成执行计划（不访问存储）。
func Plan(sqlText string) (plan.Plan, error) {
	stmt, err := sqlparser.Parse(sqlText)
	if err != nil {
		return plan.Plan{}, err
	}
	return planner.Build(stmt)
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
