package session

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/hanami/tidets/commons/errors"
	"github.com/hanami/tidets/protocol/convert"
	pb "github.com/hanami/tidets/protocol/grpc-datanode/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const defaultDialTimeout = 10 * time.Second

type sessionConnection struct {
	host     string
	port     int
	username string
	password string
	version  Version

	enableRPCCompression bool
	fetchSize            int64

	mu        sync.Mutex
	conn      *grpc.ClientConn
	client    pb.DataNodeSessionServiceClient
	sessionID int64
}

type dialParams struct {
	host                 string
	port                 int
	username             string
	password             string
	version              Version
	enableRPCCompression bool
	fetchSize            int64
}

func openSessionConnection(ctx context.Context, p dialParams) (*sessionConnection, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	sc := &sessionConnection{
		host:                 p.host,
		port:                 p.port,
		username:             p.username,
		password:             p.password,
		version:              p.version,
		enableRPCCompression: p.enableRPCCompression,
		fetchSize:            p.fetchSize,
	}

	if err := sc.connect(ctx); err != nil {
		return nil, err
	}
	if err := sc.openSession(ctx); err != nil {
		_ = sc.close()
		return nil, err
	}
	return sc, nil
}

func (c *sessionConnection) connect(ctx context.Context) error {
	addr := fmt.Sprintf("%s:%d", c.host, c.port)

	dialCtx, cancel := context.WithTimeout(ctx, defaultDialTimeout)
	defer cancel()

	conn, err := grpc.DialContext(
		dialCtx,
		addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		return commons.Wrap("session", commons.CodeUnavailable, fmt.Sprintf("grpc dial %s", addr), err)
	}

	c.mu.Lock()
	defer c.mu.Unlock()
	if c.conn != nil {
		_ = conn.Close()
		return commons.ErrClientSessionConnEstablished
	}
	c.conn = conn
	c.client = pb.NewDataNodeSessionServiceClient(conn)
	return nil
}

func (c *sessionConnection) openSession(ctx context.Context) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	c.mu.Lock()
	client := c.client
	c.mu.Unlock()
	if client == nil {
		return commons.ErrClientSessionGRPCNotReady
	}

	resp, err := client.OpenSession(ctx, &pb.OpenSessionRequest{
		Username:             c.username,
		Password:             c.password,
		Version:              string(c.version),
		EnableRpcCompression: c.enableRPCCompression,
		FetchSize:            c.fetchSize,
	})
	if err != nil {
		return commons.Wrap("session", commons.CodeInternal, "OpenSession rpc", err)
	}

	c.mu.Lock()
	c.sessionID = resp.GetSessionId()
	c.mu.Unlock()
	return nil
}

func (c *sessionConnection) close() error {
	c.mu.Lock()
	client := c.client
	sessionID := c.sessionID
	conn := c.conn

	c.client = nil
	c.sessionID = 0
	c.conn = nil
	c.mu.Unlock()

	if client != nil && sessionID != 0 {
		closeCtx, cancel := context.WithTimeout(context.Background(), defaultDialTimeout)
		_, _ = client.CloseSession(closeCtx, &pb.CloseSessionRequest{SessionId: sessionID})
		cancel()
	}
	if conn != nil {
		return conn.Close()
	}
	return nil
}

func (c *sessionConnection) insertBatch(ctx context.Context, devicePath string, points []BatchPoint) error {
	c.mu.Lock()
	client := c.client
	sessionID := c.sessionID
	c.mu.Unlock()
	if client == nil || sessionID == 0 {
		return commons.ErrClientSessionNotOpened
	}
	batch := make([]*pb.BatchPoint, 0, len(points))
	for _, p := range points {
		pbVal, err := convert.ToPB(p.Value)
		if err != nil {
			return commons.Wrap("session", commons.CodeInvalidArgument, "batch point value", err)
		}
		batch = append(batch, &pb.BatchPoint{
			Measurement: p.Measurement,
			Timestamp:   p.Timestamp,
			Value:       pbVal,
		})
	}
	_, err := client.InsertBatch(ctx, &pb.InsertBatchRequest{
		SessionId:  sessionID,
		DevicePath: devicePath,
		Points:     batch,
	})
	if err != nil {
		return commons.Wrap("session", commons.CodeInternal, "InsertBatch rpc", err)
	}
	return nil
}

func (c *sessionConnection) insertPoint(ctx context.Context, devicePath, measurement string, timestamp int64, value Value) error {
	c.mu.Lock()
	client := c.client
	sessionID := c.sessionID
	c.mu.Unlock()
	if client == nil || sessionID == 0 {
		return commons.ErrClientSessionNotOpened
	}
	pbVal, err := convert.ToPB(value)
	if err != nil {
		return commons.Wrap("session", commons.CodeInvalidArgument, "point value", err)
	}
	_, err = client.InsertPoint(ctx, &pb.InsertPointRequest{
		SessionId:   sessionID,
		DevicePath:  devicePath,
		Measurement: measurement,
		Timestamp:   timestamp,
		Value:       pbVal,
	})
	if err != nil {
		return commons.Wrap("session", commons.CodeInternal, "InsertPoint rpc", err)
	}
	return nil
}

func (c *sessionConnection) getSessionID() int64 {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.sessionID
}

func (c *sessionConnection) queryRange(ctx context.Context, devicePath, measurement string, startTime, endTime int64, limit int) ([]Point, error) {
	c.mu.Lock()
	client := c.client
	sessionID := c.sessionID
	c.mu.Unlock()
	if client == nil || sessionID == 0 {
		return nil, commons.ErrClientSessionNotOpened
	}
	resp, err := client.QueryRange(ctx, &pb.QueryRangeRequest{
		SessionId:   sessionID,
		DevicePath:  devicePath,
		Measurement: measurement,
		StartTime:   startTime,
		EndTime:     endTime,
		Limit:       int32(limit),
	})
	if err != nil {
		return nil, commons.Wrap("session", commons.CodeInternal, "QueryRange rpc", err)
	}
	out := make([]Point, 0, len(resp.GetPoints()))
	for _, p := range resp.GetPoints() {
		val, err := convert.FromPB(p.GetValue())
		if err != nil {
			return nil, commons.Wrap("session", commons.CodeInternal, "query point value", err)
		}
		out = append(out, Point{Timestamp: p.GetTimestamp(), Value: val})
	}
	return out, nil
}

func (c *sessionConnection) executeSQL(ctx context.Context, sql string) (*SQLResult, error) {
	c.mu.Lock()
	client := c.client
	sessionID := c.sessionID
	c.mu.Unlock()
	if client == nil || sessionID == 0 {
		return nil, commons.ErrClientSessionNotOpened
	}
	resp, err := client.ExecuteSQL(ctx, &pb.ExecuteSQLRequest{
		SessionId: sessionID,
		Sql:       sql,
	})
	if err != nil {
		return nil, commons.Wrap("session", commons.CodeInternal, "ExecuteSQL rpc", err)
	}
	out := &SQLResult{
		AffectedRows: int(resp.GetAffectedRows()),
		ColumnNames:  resp.GetColumnNames(),
	}
	for _, row := range resp.GetRows() {
		val, err := convert.FromPB(row.GetValue())
		if err != nil {
			return nil, commons.Wrap("session", commons.CodeInternal, "sql row value", err)
		}
		out.Rows = append(out.Rows, Point{Timestamp: row.GetTimestamp(), Value: val})
	}
	for _, row := range resp.GetCatalogRows() {
		out.CatalogRows = append(out.CatalogRows, CatalogRow{Columns: row.GetColumns()})
	}
	return out, nil
}
