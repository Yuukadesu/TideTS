// TideTS DataNode：gRPC 服务端入口。
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/hanami/tidets/core/datanode/grpcserver"
	"github.com/hanami/tidets/core/metrics"
	"github.com/hanami/tidets/core/queryengine"
	"github.com/hanami/tidets/core/storageengine"
	"github.com/hanami/tidets/core/storageengine/wal"
	pb "github.com/hanami/tidets/protocol/grpc-datanode/pb"

	"google.golang.org/grpc"
)

const defaultListenAddr = ":5556"

func main() {
	addr := flag.String("addr", defaultListenAddr, "gRPC listen address")
	metricsAddr := flag.String("metrics-addr", ":9090", "metrics HTTP listen address (empty to disable)")
	dataDir := flag.String("data-dir", "./data", "storage data directory")
	flushAt := flag.Int("flush-at", 0, "flush memtable to segment when point count reaches this (0 = default)")
	asyncFlush := flag.Bool("async-flush", true, "enable async flush worker")
	sealAfter := flag.Int("seal-after-flushes", 0, "seal active.seg after N flushes (0 = default)")
	compactThreshold := flag.Int("compact-threshold", 0, "trigger compaction when sealed segments reach this (0 = default)")
	compactMerge := flag.Int("compact-merge", 0, "merge N oldest segments per compaction (0 = default)")
	walSync := flag.String("wal-sync", "always", "WAL sync mode: always|onflush")
	walTruncate := flag.Bool("wal-truncate", false, "truncate wal.log to 0 when idle (requires checkpoint)")
	flag.Parse()

	asyncFlushVal := *asyncFlush
	opts := storageengine.Options{
		DataDir:          *dataDir,
		FlushAt:          *flushAt,
		AsyncFlush:       &asyncFlushVal,
		WALTruncate:      *walTruncate,
		SealAfterFlushes: *sealAfter,
		CompactThreshold: *compactThreshold,
		CompactMerge:     *compactMerge,
		WALSync:          wal.SyncAlways,
	}
	switch *walSync {
	case "always":
		opts.WALSync = wal.SyncAlways
	case "onflush":
		opts.WALSync = wal.SyncOnFlush
	default:
		log.Fatalf("invalid -wal-sync %q, must be always|onflush", *walSync)
	}

	engine, err := storageengine.OpenWithOptions(opts)
	if err != nil {
		log.Fatalf("open storage: %v", err)
	}
	defer func() {
		if err := engine.Close(); err != nil {
			log.Printf("close storage: %v", err)
		}
	}()

	lis, err := net.Listen("tcp", *addr)
	if err != nil {
		log.Fatalf("listen %s: %v", *addr, err)
	}

	srv, err := grpcserver.New(engine)
	if err != nil {
		log.Fatalf("open catalog: %v", err)
	}

	var (
		reg        *metrics.Registry
		grpcServer *grpc.Server
		metricsSrv *http.Server
	)
	if *metricsAddr != "" {
		reg = metrics.NewRegistry()
		reg.RegisterStorageCollector(engine)
		reg.RegisterSessionCollector(srv.SessionManager().ActiveCount)
		engine.SetHooks(reg.StorageHooks())
		srv.SQLService().SetHooks(queryengine.Hooks{OnPlanExecuted: reg.ObserveSQL})
		srv.SetHooks(grpcserver.Hooks{
			OnSQL:      reg.ObserveSQL,
			OnRPCItems: reg.ObserveRPCItems,
		})
		grpcServer = grpc.NewServer(grpc.UnaryInterceptor(reg.UnaryServerInterceptor()))
		metricsSrv = metrics.NewHTTPServer(*metricsAddr, reg)
		go func() {
			log.Printf("TideTS metrics listening on %s", *metricsAddr)
			if err := metricsSrv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				log.Fatalf("metrics serve: %v", err)
			}
		}()
	} else {
		grpcServer = grpc.NewServer()
	}
	pb.RegisterDataNodeSessionServiceServer(grpcServer, srv)

	go func() {
		log.Printf(
			"TideTS DataNode listening on %s, data-dir=%s, flush-at=%d, async-flush=%v, wal-sync=%s, wal-truncate=%v, metrics-addr=%s",
			*addr, *dataDir, *flushAt, *asyncFlush, *walSync, *walTruncate, *metricsAddr,
		)
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("serve: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("shutting down...")
	grpcServer.GracefulStop()
	if metricsSrv != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := metricsSrv.Shutdown(ctx); err != nil {
			log.Printf("shutdown metrics: %v", err)
		}
	}
	fmt.Println("bye")
}
