package queryengine

import (
	"time"

	"github.com/hanami/tidets/core/queryengine/plan"
)

// Hooks 允许外层注入 SQL 观测回调，避免 queryengine 直接依赖具体 metrics 框架。
type Hooks struct {
	OnPlanExecuted func(kind plan.Kind, success bool, errorClass string, duration time.Duration)
}
