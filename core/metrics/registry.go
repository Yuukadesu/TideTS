package metrics

import (
	"net/http"
	"runtime/debug"
	"time"

	"github.com/hanami/tidets/core/queryengine/plan"
	"github.com/hanami/tidets/core/storageengine"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc/codes"
)

const namespace = "tidets"

// Registry 持有 TideTS 全部 Prometheus 指标与 collector。
type Registry struct {
	reg *prometheus.Registry

	rpcRequests          *prometheus.CounterVec
	rpcRequestDuration   *prometheus.HistogramVec
	rpcRequestItems      *prometheus.CounterVec
	rpcResponseItems     *prometheus.CounterVec
	sqlRequests          *prometheus.CounterVec
	sqlRequestDuration   *prometheus.HistogramVec
	storageWriteRequests *prometheus.CounterVec
	storageWritePoints   *prometheus.CounterVec
	storageWriteDuration *prometheus.HistogramVec
	storageReadRequests  *prometheus.CounterVec
	storageReadPoints    *prometheus.CounterVec
	storageReadDuration  *prometheus.HistogramVec
	storageWALEvents     *prometheus.CounterVec
	storageWALRecords    *prometheus.CounterVec
	storageTombstoneOps  *prometheus.CounterVec
	storageTombstoneRg   *prometheus.CounterVec
	storageFlushTotal    prometheus.Counter
	storageFlushPoints   prometheus.Counter
	storageFlushDuration prometheus.Histogram
	storageCompactTotal  prometheus.Counter
	storageCompactInput  prometheus.Counter
	storageCompactOutput prometheus.Counter
	storageCompactSecond prometheus.Histogram
	nodeCollector        *nodeCollector
}

// NewRegistry 创建 TideTS metrics 注册中心。
func NewRegistry() *Registry {
	reg := prometheus.NewRegistry()
	r := &Registry{
		reg: reg,
		rpcRequests: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Namespace: namespace,
				Subsystem: "grpc",
				Name:      "requests_total",
				Help:      "Total number of gRPC requests by method and status code.",
			},
			[]string{"method", "code"},
		),
		rpcRequestDuration: prometheus.NewHistogramVec(
			prometheus.HistogramOpts{
				Namespace: namespace,
				Subsystem: "grpc",
				Name:      "request_duration_seconds",
				Help:      "Duration of gRPC requests by method.",
				Buckets:   prometheus.DefBuckets,
			},
			[]string{"method"},
		),
		rpcRequestItems: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Namespace: namespace,
				Subsystem: "grpc",
				Name:      "request_items_total",
				Help:      "Total number of logical request items by gRPC method.",
			},
			[]string{"method"},
		),
		rpcResponseItems: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Namespace: namespace,
				Subsystem: "grpc",
				Name:      "response_items_total",
				Help:      "Total number of logical response items by gRPC method.",
			},
			[]string{"method"},
		),
		sqlRequests: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Namespace: namespace,
				Subsystem: "sql",
				Name:      "requests_total",
				Help:      "Total number of SQL requests by plan kind, success and error class.",
			},
			[]string{"kind", "success", "error_class"},
		),
		sqlRequestDuration: prometheus.NewHistogramVec(
			prometheus.HistogramOpts{
				Namespace: namespace,
				Subsystem: "sql",
				Name:      "request_duration_seconds",
				Help:      "Duration of SQL requests by plan kind, success and error class.",
				Buckets:   prometheus.DefBuckets,
			},
			[]string{"kind", "success", "error_class"},
		),
		storageWriteRequests: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Namespace: namespace,
				Subsystem: "storage",
				Name:      "write_requests_total",
				Help:      "Total number of storage write requests by operation.",
			},
			[]string{"op"},
		),
		storageWritePoints: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Namespace: namespace,
				Subsystem: "storage",
				Name:      "write_points_total",
				Help:      "Total number of storage write points by operation.",
			},
			[]string{"op"},
		),
		storageWriteDuration: prometheus.NewHistogramVec(
			prometheus.HistogramOpts{
				Namespace: namespace,
				Subsystem: "storage",
				Name:      "write_duration_seconds",
				Help:      "Duration of storage write requests by operation.",
				Buckets:   prometheus.DefBuckets,
			},
			[]string{"op"},
		),
		storageReadRequests: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Namespace: namespace,
				Subsystem: "storage",
				Name:      "read_requests_total",
				Help:      "Total number of storage read requests by operation.",
			},
			[]string{"op"},
		),
		storageReadPoints: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Namespace: namespace,
				Subsystem: "storage",
				Name:      "read_points_total",
				Help:      "Total number of storage points returned by read operation.",
			},
			[]string{"op"},
		),
		storageReadDuration: prometheus.NewHistogramVec(
			prometheus.HistogramOpts{
				Namespace: namespace,
				Subsystem: "storage",
				Name:      "read_duration_seconds",
				Help:      "Duration of storage read requests by operation.",
				Buckets:   prometheus.DefBuckets,
			},
			[]string{"op"},
		),
		storageWALEvents: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Namespace: namespace,
				Subsystem: "storage",
				Name:      "wal_events_total",
				Help:      "Total number of WAL events by operation.",
			},
			[]string{"op"},
		),
		storageWALRecords: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Namespace: namespace,
				Subsystem: "storage",
				Name:      "wal_records_total",
				Help:      "Total number of WAL records touched by operation.",
			},
			[]string{"op"},
		),
		storageTombstoneOps: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Namespace: namespace,
				Subsystem: "storage",
				Name:      "tombstone_events_total",
				Help:      "Total number of tombstone events by operation.",
			},
			[]string{"op"},
		),
		storageTombstoneRg: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Namespace: namespace,
				Subsystem: "storage",
				Name:      "tombstone_ranges_total",
				Help:      "Total number of tombstone ranges marked or pruned by operation.",
			},
			[]string{"op"},
		),
		storageFlushTotal: prometheus.NewCounter(
			prometheus.CounterOpts{
				Namespace: namespace,
				Subsystem: "storage",
				Name:      "flush_total",
				Help:      "Total number of successful storage flushes.",
			},
		),
		storageFlushPoints: prometheus.NewCounter(
			prometheus.CounterOpts{
				Namespace: namespace,
				Subsystem: "storage",
				Name:      "flush_points_total",
				Help:      "Total number of points flushed to segment files.",
			},
		),
		storageFlushDuration: prometheus.NewHistogram(
			prometheus.HistogramOpts{
				Namespace: namespace,
				Subsystem: "storage",
				Name:      "flush_duration_seconds",
				Help:      "Duration of successful storage flushes.",
				Buckets:   prometheus.DefBuckets,
			},
		),
		storageCompactTotal: prometheus.NewCounter(
			prometheus.CounterOpts{
				Namespace: namespace,
				Subsystem: "storage",
				Name:      "compact_total",
				Help:      "Total number of successful storage compactions.",
			},
		),
		storageCompactInput: prometheus.NewCounter(
			prometheus.CounterOpts{
				Namespace: namespace,
				Subsystem: "storage",
				Name:      "compact_input_files_total",
				Help:      "Total number of compact input files.",
			},
		),
		storageCompactOutput: prometheus.NewCounter(
			prometheus.CounterOpts{
				Namespace: namespace,
				Subsystem: "storage",
				Name:      "compact_output_files_total",
				Help:      "Total number of compact output files.",
			},
		),
		storageCompactSecond: prometheus.NewHistogram(
			prometheus.HistogramOpts{
				Namespace: namespace,
				Subsystem: "storage",
				Name:      "compact_duration_seconds",
				Help:      "Duration of successful storage compactions.",
				Buckets:   prometheus.DefBuckets,
			},
		),
	}
	r.nodeCollector = newNodeCollector()

	reg.MustRegister(
		collectors.NewGoCollector(),
		collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}),
		r.nodeCollector,
		r.rpcRequests,
		r.rpcRequestDuration,
		r.rpcRequestItems,
		r.rpcResponseItems,
		r.sqlRequests,
		r.sqlRequestDuration,
		r.storageWriteRequests,
		r.storageWritePoints,
		r.storageWriteDuration,
		r.storageReadRequests,
		r.storageReadPoints,
		r.storageReadDuration,
		r.storageWALEvents,
		r.storageWALRecords,
		r.storageTombstoneOps,
		r.storageTombstoneRg,
		r.storageFlushTotal,
		r.storageFlushPoints,
		r.storageFlushDuration,
		r.storageCompactTotal,
		r.storageCompactInput,
		r.storageCompactOutput,
		r.storageCompactSecond,
	)
	return r
}

// PrometheusRegistry 返回底层 Prometheus registry。
func (r *Registry) PrometheusRegistry() *prometheus.Registry {
	if r == nil {
		return nil
	}
	return r.reg
}

// Handler 返回 /metrics HTTP handler。
func (r *Registry) Handler() http.Handler {
	if r == nil || r.reg == nil {
		return promhttp.Handler()
	}
	return promhttp.HandlerFor(r.reg, promhttp.HandlerOpts{})
}

// RegisterStorageCollector 注册存储层 collector。
func (r *Registry) RegisterStorageCollector(engine *storageengine.Engine) {
	if r == nil || r.reg == nil || engine == nil {
		return
	}
	r.reg.MustRegister(NewStorageCollector(engine))
}

// RegisterSessionCollector 注册活跃会话 collector。
func (r *Registry) RegisterSessionCollector(active func() int) {
	if r == nil || r.reg == nil || active == nil {
		return
	}
	r.reg.MustRegister(NewSessionCollector(active))
}

// StorageHooks 返回供 storageengine 注入的观测 hooks。
func (r *Registry) StorageHooks() storageengine.Hooks {
	if r == nil {
		return storageengine.Hooks{}
	}
	return storageengine.Hooks{
		OnWrite: func(op string, points int, duration time.Duration) {
			r.storageWriteRequests.WithLabelValues(op).Inc()
			r.storageWritePoints.WithLabelValues(op).Add(float64(points))
			r.storageWriteDuration.WithLabelValues(op).Observe(duration.Seconds())
		},
		OnRead: func(op string, points int, duration time.Duration) {
			r.storageReadRequests.WithLabelValues(op).Inc()
			r.storageReadPoints.WithLabelValues(op).Add(float64(points))
			r.storageReadDuration.WithLabelValues(op).Observe(duration.Seconds())
		},
		OnWAL: func(op string, records int) {
			r.storageWALEvents.WithLabelValues(op).Inc()
			r.storageWALRecords.WithLabelValues(op).Add(float64(records))
		},
		OnTombstone: func(op string, ranges int) {
			r.storageTombstoneOps.WithLabelValues(op).Inc()
			r.storageTombstoneRg.WithLabelValues(op).Add(float64(ranges))
		},
		OnFlush: func(points int, duration time.Duration) {
			r.storageFlushTotal.Inc()
			r.storageFlushPoints.Add(float64(points))
			r.storageFlushDuration.Observe(duration.Seconds())
		},
		OnCompact: func(duration time.Duration, inputFiles, outputFiles int) {
			r.storageCompactTotal.Inc()
			r.storageCompactInput.Add(float64(inputFiles))
			r.storageCompactOutput.Add(float64(outputFiles))
			r.storageCompactSecond.Observe(duration.Seconds())
		},
	}
}

// ObserveRPC 记录一次 gRPC 请求。
func (r *Registry) ObserveRPC(method string, code codes.Code, duration time.Duration) {
	if r == nil {
		return
	}
	method = trimMethod(method)
	r.rpcRequests.WithLabelValues(method, code.String()).Inc()
	r.rpcRequestDuration.WithLabelValues(method).Observe(duration.Seconds())
}

// ObserveRPCItems 记录一次 gRPC 请求/响应业务量。
func (r *Registry) ObserveRPCItems(method string, requestItems, responseItems int) {
	if r == nil {
		return
	}
	method = trimMethod(method)
	if requestItems > 0 {
		r.rpcRequestItems.WithLabelValues(method).Add(float64(requestItems))
	}
	if responseItems > 0 {
		r.rpcResponseItems.WithLabelValues(method).Add(float64(responseItems))
	}
}

// ObserveSQL 记录一次 SQL 执行。
func (r *Registry) ObserveSQL(kind plan.Kind, success bool, errorClass string, duration time.Duration) {
	if r == nil {
		return
	}
	label := planKindLabel(kind)
	if errorClass == "" {
		errorClass = "none"
	}
	successLabel := "false"
	if success {
		successLabel = "true"
	}
	r.sqlRequests.WithLabelValues(label, successLabel, errorClass).Inc()
	r.sqlRequestDuration.WithLabelValues(label, successLabel, errorClass).Observe(duration.Seconds())
}

type nodeCollector struct {
	startedAt  time.Time
	startDesc  *prometheus.Desc
	uptimeDesc *prometheus.Desc
	buildInfo  *prometheus.Desc
	version    string
	goVersion  string
}

func newNodeCollector() *nodeCollector {
	version := "dev"
	goVersion := "unknown"
	if info, ok := debug.ReadBuildInfo(); ok && info != nil {
		if info.Main.Version != "" && info.Main.Version != "(devel)" {
			version = info.Main.Version
		}
		if info.GoVersion != "" {
			goVersion = info.GoVersion
		}
	}
	return &nodeCollector{
		startedAt: time.Now(),
		startDesc: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "node", "start_time_seconds"),
			"Unix start time of the TideTS node process.",
			nil, nil,
		),
		uptimeDesc: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "node", "uptime_seconds"),
			"Uptime of the TideTS node process in seconds.",
			nil, nil,
		),
		buildInfo: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "build_info"),
			"Build information about the TideTS binary.",
			[]string{"version", "go_version"},
			nil,
		),
		version:   version,
		goVersion: goVersion,
	}
}

func (c *nodeCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.startDesc
	ch <- c.uptimeDesc
	ch <- c.buildInfo
}

func (c *nodeCollector) Collect(ch chan<- prometheus.Metric) {
	if c == nil {
		return
	}
	ch <- prometheus.MustNewConstMetric(c.startDesc, prometheus.GaugeValue, float64(c.startedAt.Unix()))
	ch <- prometheus.MustNewConstMetric(c.uptimeDesc, prometheus.GaugeValue, time.Since(c.startedAt).Seconds())
	ch <- prometheus.MustNewConstMetric(c.buildInfo, prometheus.GaugeValue, 1, c.version, c.goVersion)
}
