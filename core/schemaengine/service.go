package schemaengine

import (
	"sync"

	"github.com/hanami/tidets/commons/errors"
	"github.com/hanami/tidets/core/tsmodel"
)

const snapshotEveryOps = 100

// Service 管理时间序列 schema：MTree + mlog + snapshot（对齐 IoTDB Schema Region 子集）。
type Service struct {
	dir     string
	tree    *MTree
	mlog    *mlog
	mu      sync.Mutex
	mlogOps int
}

// Open 从 system/schema 加载 snapshot 并 replay mlog。
func Open(dataDir string) (*Service, error) {
	dir := schemaDir(dataDir)

	snap, err := loadSnapshot(dir)
	if err != nil {
		return nil, err
	}
	tree := newMTree()
	var replayFrom int64
	if snap != nil {
		tree = newMTreeFromSeries(snap.Series)
		replayFrom = snap.MlogOffset
	}

	ml, err := openMlog(dir)
	if err != nil {
		return nil, err
	}
	s := &Service{
		dir:  dir,
		tree: tree,
		mlog: ml,
	}
	if err := replayMlog(dir, replayFrom, func(ts Timeseries) error {
		return s.tree.put(ts)
	}); err != nil {
		_ = s.Close()
		return nil, err
	}
	return s, nil
}

// Close 关闭前写 snapshot。
func (s *Service) Close() error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if err := s.snapshotLocked(); err != nil {
		return err
	}
	return s.mlog.close()
}

// CreateTimeseries 显式创建序列（IoTDB CREATE TIMESERIES 子集）。
func (s *Service) CreateTimeseries(devicePath, measurement string, dt tsmodel.DataType) (Timeseries, error) {
	if devicePath == "" || measurement == "" {
		return Timeseries{}, commons.ErrSchemaPathRequired
	}
	if dt == tsmodel.DataTypeUnknown {
		return Timeseries{}, commons.ErrSchemaDataTypeRequired
	}
	ts := NewTimeseries(devicePath, measurement, dt)

	s.mu.Lock()
	defer s.mu.Unlock()
	if s.tree.hasLocked(ts.Key()) {
		return Timeseries{}, commons.ErrSchemaTimeseriesExists
	}
	if err := s.mlog.appendCreate(ts); err != nil {
		return Timeseries{}, err
	}
	if err := s.tree.putLocked(ts); err != nil {
		return Timeseries{}, err
	}
	s.mlogOps++
	if err := s.snapshotIfNeededLocked(); err != nil {
		return Timeseries{}, err
	}
	return ts, nil
}

// Get 查询 schema。
func (s *Service) Get(key tsmodel.SeriesKey) (Timeseries, bool) {
	return s.tree.get(key)
}

// Has 序列是否已注册。
func (s *Service) Has(key tsmodel.SeriesKey) bool {
	return s.tree.has(key)
}

// ValidateInsert 写入前校验；若序列不存在则自动注册 schema。
func (s *Service) ValidateInsert(key tsmodel.SeriesKey, value tsmodel.Value) (Timeseries, error) {
	if key.DevicePath == "" || key.Measurement == "" {
		return Timeseries{}, commons.ErrSchemaPathRequired
	}
	if err := value.Validate(); err != nil {
		return Timeseries{}, err
	}
	if ts, ok := s.tree.get(key); ok {
		if ts.DataType != value.Type {
			return Timeseries{}, commons.ErrSchemaDataTypeMismatch
		}
		return ts, nil
	}

	ts := NewTimeseries(key.DevicePath, key.Measurement, value.Type)
	s.mu.Lock()
	defer s.mu.Unlock()
	if existing, ok := s.tree.getLocked(key); ok {
		if existing.DataType != value.Type {
			return Timeseries{}, commons.ErrSchemaDataTypeMismatch
		}
		return existing, nil
	}
	if err := s.mlog.appendCreate(ts); err != nil {
		return Timeseries{}, err
	}
	if err := s.tree.putLocked(ts); err != nil {
		return Timeseries{}, err
	}
	s.mlogOps++
	if err := s.snapshotIfNeededLocked(); err != nil {
		return Timeseries{}, err
	}
	return ts, nil
}

// RegisterBatch 批量注册（bootstrap，不覆盖已有类型）。
func (s *Service) RegisterBatch(keys []tsmodel.SeriesKey, dataTypeOf func(key tsmodel.SeriesKey) tsmodel.DataType) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	added := make([]Timeseries, 0)
	for _, key := range keys {
		if key.DevicePath == "" || key.Measurement == "" {
			continue
		}
		if s.tree.hasLocked(key) {
			continue
		}
		dt := dataTypeOf(key)
		if dt == tsmodel.DataTypeUnknown {
			continue
		}
		added = append(added, NewTimeseries(key.DevicePath, key.Measurement, dt))
	}
	if len(added) == 0 {
		return nil
	}
	if err := s.mlog.appendBatch(added); err != nil {
		return err
	}
	for _, ts := range added {
		if err := s.tree.putLocked(ts); err != nil {
			return err
		}
	}
	s.mlogOps += len(added)
	return s.snapshotIfNeededLocked()
}

// ListDevices 列举设备路径（供 metadata 层委托）。
func (s *Service) ListDevices(prefix string) []string {
	return s.tree.ListDevices(prefix)
}

// ListMeasurements 列举设备下测点 schema。
func (s *Service) ListMeasurements(devicePath string) []Timeseries {
	return s.tree.ListMeasurements(devicePath)
}

// ChildPaths 返回 prefix 下直接子路径。
func (s *Service) ChildPaths(prefix string) []ChildPath {
	return s.tree.ChildPaths(prefix)
}

func (s *Service) snapshotIfNeededLocked() error {
	if s.mlogOps < snapshotEveryOps {
		return nil
	}
	if err := s.snapshotLocked(); err != nil {
		return err
	}
	s.mlogOps = 0
	return nil
}

func (s *Service) snapshotLocked() error {
	off, err := s.mlog.offset()
	if err != nil {
		return err
	}
	return saveSnapshot(s.dir, snapshotMeta{
		MlogOffset: off,
		Series:     s.tree.snapshotSeries(),
	})
}
