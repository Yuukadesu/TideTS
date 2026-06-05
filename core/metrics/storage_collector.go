package metrics

import (
	"github.com/hanami/tidets/core/storageengine"
	"github.com/prometheus/client_golang/prometheus"
)

// StorageCollector 在 scrape 时读取 Engine.Stats() 暴露存储层 gauge。
type StorageCollector struct {
	engine *storageengine.Engine

	walBytes              *prometheus.Desc
	segmentCount          *prometheus.Desc
	sealedSegmentCount    *prometheus.Desc
	activeSegmentBytes    *prometheus.Desc
	memTablePoints        *prometheus.Desc
	delayedMemTablePoints *prometheus.Desc
	pendingFlushBatches   *prometheus.Desc
	lastFlushTS           *prometheus.Desc
}

// NewStorageCollector 创建存储层 collector。
func NewStorageCollector(engine *storageengine.Engine) *StorageCollector {
	return &StorageCollector{
		engine: engine,
		walBytes: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "storage", "wal_bytes"),
			"Current wal.log size in bytes.",
			nil, nil,
		),
		segmentCount: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "storage", "segment_count"),
			"Current number of segment files including active segment.",
			nil, nil,
		),
		sealedSegmentCount: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "storage", "sealed_segment_count"),
			"Current number of sealed segment files.",
			nil, nil,
		),
		activeSegmentBytes: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "storage", "active_segment_bytes"),
			"Current active segment file size in bytes.",
			nil, nil,
		),
		memTablePoints: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "storage", "memtable_points"),
			"Current number of points in the normal memtable.",
			nil, nil,
		),
		delayedMemTablePoints: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "storage", "delayed_memtable_points"),
			"Current number of points in the delayed memtable.",
			nil, nil,
		),
		pendingFlushBatches: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "storage", "pending_flush_batches"),
			"Current number of pending flush batches.",
			nil, nil,
		),
		lastFlushTS: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "storage", "last_flush_timestamp_seconds"),
			"Unix timestamp in seconds of the latest successful flush.",
			nil, nil,
		),
	}
}

// Describe 实现 prometheus.Collector。
func (c *StorageCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.walBytes
	ch <- c.segmentCount
	ch <- c.sealedSegmentCount
	ch <- c.activeSegmentBytes
	ch <- c.memTablePoints
	ch <- c.delayedMemTablePoints
	ch <- c.pendingFlushBatches
	ch <- c.lastFlushTS
}

// Collect 实现 prometheus.Collector。
func (c *StorageCollector) Collect(ch chan<- prometheus.Metric) {
	if c == nil || c.engine == nil {
		return
	}
	st := c.engine.Stats()
	ch <- prometheus.MustNewConstMetric(c.walBytes, prometheus.GaugeValue, float64(st.WALBytes))
	ch <- prometheus.MustNewConstMetric(c.segmentCount, prometheus.GaugeValue, float64(st.SegmentCount))
	ch <- prometheus.MustNewConstMetric(c.sealedSegmentCount, prometheus.GaugeValue, float64(st.SealedSegmentCount))
	ch <- prometheus.MustNewConstMetric(c.activeSegmentBytes, prometheus.GaugeValue, float64(st.ActiveSegmentBytes))
	ch <- prometheus.MustNewConstMetric(c.memTablePoints, prometheus.GaugeValue, float64(st.MemTablePoints))
	ch <- prometheus.MustNewConstMetric(c.delayedMemTablePoints, prometheus.GaugeValue, float64(st.DelayedMemTablePoints))
	ch <- prometheus.MustNewConstMetric(c.pendingFlushBatches, prometheus.GaugeValue, float64(st.PendingFlushBatches))

	lastFlush := 0.0
	if !st.LastFlushAt.IsZero() {
		lastFlush = float64(st.LastFlushAt.Unix())
	}
	ch <- prometheus.MustNewConstMetric(c.lastFlushTS, prometheus.GaugeValue, lastFlush)
}
