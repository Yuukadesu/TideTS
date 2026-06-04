package grpcserver

import (
	"github.com/hanami/tidets/core/datanode/metadata"
	"github.com/hanami/tidets/core/dataplane"
	"github.com/hanami/tidets/core/schemaengine"
	"github.com/hanami/tidets/core/storageengine"
)

// Bootstrap 启动时打开的 schema、metadata 与数据面网关。
type Bootstrap struct {
	Schema   *schemaengine.Service
	Metadata *metadata.Manager
	Gateway  *dataplane.Gateway
}

// OpenBootstrap 打开 schema + metadata，并与存储层已有序列对齐。
func OpenBootstrap(engine *storageengine.Engine) (*Bootstrap, error) {
	schemaSvc, err := schemaengine.Open(engine.DataDir())
	if err != nil {
		return nil, err
	}
	metaMgr := metadata.New(schemaSvc)
	if err := metaMgr.ReconcileFromStorage(engine.KnownSeriesTypes()); err != nil {
		_ = schemaSvc.Close()
		return nil, err
	}
	engine.BindSchemaGuard(schemaSvc.Has)
	return &Bootstrap{
		Schema:   schemaSvc,
		Metadata: metaMgr,
		Gateway:  dataplane.New(engine, schemaSvc),
	}, nil
}
