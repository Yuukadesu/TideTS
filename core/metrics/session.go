package metrics

import "github.com/prometheus/client_golang/prometheus"

// SessionCollector 在 scrape 时读取当前活跃 session 数。
type SessionCollector struct {
	active func() int
	desc   *prometheus.Desc
}

// NewSessionCollector 创建 session gauge collector。
func NewSessionCollector(active func() int) *SessionCollector {
	return &SessionCollector{
		active: active,
		desc: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "session_active"),
			"Current number of active client sessions.",
			nil, nil,
		),
	}
}

// Describe 实现 prometheus.Collector。
func (c *SessionCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.desc
}

// Collect 实现 prometheus.Collector。
func (c *SessionCollector) Collect(ch chan<- prometheus.Metric) {
	if c == nil || c.active == nil {
		return
	}
	ch <- prometheus.MustNewConstMetric(c.desc, prometheus.GaugeValue, float64(c.active()))
}
