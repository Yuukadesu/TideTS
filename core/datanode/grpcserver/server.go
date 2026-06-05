// Package grpcserver 实现 DataNode 面向客户端的 gRPC 会话服务。
package grpcserver

import (
	"context"
	"log"
	"net"
	"strconv"
	"strings"
	"time"

	"github.com/hanami/tidets/commons/errors"
	"github.com/hanami/tidets/core/datanode/auth"
	"github.com/hanami/tidets/core/datanode/metadata"
	"github.com/hanami/tidets/core/datanode/session"
	"github.com/hanami/tidets/core/datanode/sink"
	"github.com/hanami/tidets/core/dataplane"
	"github.com/hanami/tidets/core/queryengine"
	"github.com/hanami/tidets/core/queryengine/backend"
	"github.com/hanami/tidets/core/queryengine/plan"
	"github.com/hanami/tidets/core/schemaengine"
	"github.com/hanami/tidets/core/storageengine"
	"github.com/hanami/tidets/core/tsmodel"
	"github.com/hanami/tidets/protocol/convert"
	pb "github.com/hanami/tidets/protocol/grpc-datanode/pb"

	"google.golang.org/grpc/peer"
)

// Server 实现 DataNodeSessionService。
type Server struct {
	pb.UnimplementedDataNodeSessionServiceServer

	gateway  *dataplane.Gateway
	schema   *schemaengine.Service
	metadata *metadata.Manager
	sql      *queryengine.Service
	mgr      *session.Manager
	auth     auth.Checker
	hooks    Hooks
}

// New 创建 gRPC 服务实现；engine 由调用方负责 Open/Close。
func New(engine *storageengine.Engine) (*Server, error) {
	boot, err := OpenBootstrap(engine)
	if err != nil {
		return nil, err
	}
	authChecker := auth.DefaultChecker()
	catalog := &backend.MetadataCatalog{Meta: boot.Metadata}
	return &Server{
		gateway:  boot.Gateway,
		schema:   boot.Schema,
		metadata: boot.Metadata,
		sql:      queryengine.NewService(engine, boot.Gateway, catalog),
		mgr:      session.NewManager(&authChecker),
		auth:     authChecker,
	}, nil
}

// Schema 返回 schema 服务（DDL / 写入校验）。
func (s *Server) Schema() *schemaengine.Service { return s.schema }

// Metadata 返回元数据目录（SHOW DEVICES / SHOW TIMESERIES 等扩展用）。
func (s *Server) Metadata() *metadata.Manager { return s.metadata }

// Gateway 返回统一数据读写入口。
func (s *Server) Gateway() *dataplane.Gateway { return s.gateway }

// SQLService 返回 SQL 服务，供外层注入 hooks。
func (s *Server) SQLService() *queryengine.Service { return s.sql }

// SessionManager 返回会话管理器，供外层暴露活跃会话等 metrics。
func (s *Server) SessionManager() *session.Manager { return s.mgr }

func (s *Server) OpenSession(ctx context.Context, req *pb.OpenSessionRequest) (*pb.OpenSessionResponse, error) {
	if req.GetUsername() == "" {
		return nil, commons.ToGRPCStatus(commons.ErrGRPCUsernameRequired)
	}

	addr, port := peerHostPort(ctx)
	sess, result := s.mgr.OpenGRPCSession(session.OpenParams{
		Login: session.LoginParams{
			Username:             req.GetUsername(),
			Password:             req.GetPassword(),
			ZoneID:               "",
			ClientVersion:        session.ClientVersion(req.GetVersion()),
			FetchSize:            req.GetFetchSize(),
			EnableRPCCompression: req.GetEnableRpcCompression(),
		},
		ClientAddr: addr,
		ClientPort: port,
	})
	if !result.OK {
		return nil, commons.ToGRPCStatus(commons.Errorf("grpc", commons.CodeUnauthenticated, "%s", result.Message))
	}

	log.Printf("OpenSession: user=%s session_id=%d user_id=%d fetch_size=%d client=%s",
		req.GetUsername(), result.SessionID, result.UserID, req.GetFetchSize(), sess.ConnectionID())

	return &pb.OpenSessionResponse{SessionId: result.SessionID}, nil
}

func (s *Server) CloseSession(_ context.Context, req *pb.CloseSessionRequest) (*pb.CloseSessionResponse, error) {
	sessionID := req.GetSessionId()
	if sessionID <= 0 {
		return nil, commons.ToGRPCStatus(commons.ErrGRPCSessionIDInvalid)
	}

	if err := s.mgr.CloseSession(sessionID); err != nil {
		return nil, commons.ToGRPCStatus(err)
	}

	log.Printf("CloseSession: session_id=%d", sessionID)
	return &pb.CloseSessionResponse{}, nil
}

func (s *Server) InsertPoint(ctx context.Context, req *pb.InsertPointRequest) (*pb.InsertPointResponse, error) {
	sess, err := s.requireSession(ctx, req.GetSessionId())
	if err != nil {
		return nil, err
	}
	if req.GetDevicePath() == "" || req.GetMeasurement() == "" {
		return nil, commons.ToGRPCStatus(commons.ErrGRPCDeviceMeasurementRequired)
	}
	if req.GetTimestamp() <= 0 {
		return nil, commons.ToGRPCStatus(commons.ErrGRPCTimestampInvalid)
	}

	info := sess.SessionInfo()
	if err := s.auth.CheckPrivilege(info.User.UserID, info.User.Username, req.GetDevicePath(), auth.PrivilegeWrite); err != nil {
		return nil, commons.ToGRPCStatus(err)
	}

	val, err := convert.FromPB(req.GetValue())
	if err != nil {
		return nil, commons.ToGRPCStatus(err)
	}
	key := tsmodel.SeriesKey{
		DevicePath:  req.GetDevicePath(),
		Measurement: req.GetMeasurement(),
	}
	if err := s.gateway.Insert(key, tsmodel.Point{
		Timestamp: req.GetTimestamp(),
		Value:     val,
	}); err != nil {
		return nil, commons.ToGRPCStatus(commons.Wrap("grpc", commons.CodeInternal, "insert", err))
	}
	if s.hooks.OnRPCItems != nil {
		s.hooks.OnRPCItems("InsertPoint", 1, 0)
	}
	return &pb.InsertPointResponse{}, nil
}

func (s *Server) InsertBatch(ctx context.Context, req *pb.InsertBatchRequest) (*pb.InsertBatchResponse, error) {
	sess, err := s.requireSession(ctx, req.GetSessionId())
	if err != nil {
		return nil, err
	}
	if req.GetDevicePath() == "" {
		return nil, commons.ToGRPCStatus(commons.ErrGRPCDevicePathRequired)
	}
	if len(req.GetPoints()) == 0 {
		return nil, commons.ToGRPCStatus(commons.ErrGRPCPointsEmpty)
	}

	info := sess.SessionInfo()
	if err := s.auth.CheckPrivilege(info.User.UserID, info.User.Username, req.GetDevicePath(), auth.PrivilegeWrite); err != nil {
		return nil, commons.ToGRPCStatus(err)
	}

	records := make([]storageengine.Record, 0, len(req.GetPoints()))
	for _, pt := range req.GetPoints() {
		if pt.GetMeasurement() == "" {
			return nil, commons.ToGRPCStatus(commons.ErrGRPCMeasurementRequired)
		}
		if pt.GetTimestamp() <= 0 {
			return nil, commons.ToGRPCStatus(commons.ErrGRPCTimestampInvalid)
		}
		val, err := convert.FromPB(pt.GetValue())
		if err != nil {
			return nil, commons.ToGRPCStatus(err)
		}
		key := tsmodel.SeriesKey{
			DevicePath:  req.GetDevicePath(),
			Measurement: pt.GetMeasurement(),
		}
		records = append(records, storageengine.Record{
			Key: key,
			Point: tsmodel.Point{
				Timestamp: pt.GetTimestamp(),
				Value:     val,
			},
		})
	}

	if err := s.gateway.InsertBatch(records); err != nil {
		return nil, commons.ToGRPCStatus(commons.Wrap("grpc", commons.CodeInternal, "insert batch", err))
	}
	if s.hooks.OnRPCItems != nil {
		s.hooks.OnRPCItems("InsertBatch", len(records), int(len(records)))
	}
	return &pb.InsertBatchResponse{Inserted: int32(len(records))}, nil
}

func (s *Server) QueryRange(ctx context.Context, req *pb.QueryRangeRequest) (*pb.QueryRangeResponse, error) {
	sess, err := s.requireSession(ctx, req.GetSessionId())
	if err != nil {
		return nil, err
	}
	info := sess.SessionInfo()
	if err := s.auth.CheckPrivilege(info.User.UserID, info.User.Username, req.GetDevicePath(), auth.PrivilegeRead); err != nil {
		return nil, commons.ToGRPCStatus(err)
	}
	if req.GetDevicePath() == "" || req.GetMeasurement() == "" {
		return nil, commons.ToGRPCStatus(commons.ErrGRPCDeviceMeasurementRequired)
	}
	if req.GetStartTime() > req.GetEndTime() {
		return nil, commons.ToGRPCStatus(commons.ErrGRPCTimeRangeInvalid)
	}

	limit := queryengine.ResolveQueryLimit(int(req.GetLimit()), int(info.FetchSize))
	res, err := s.sql.QueryRange(ctx, tsmodel.SeriesKey{
		DevicePath:  req.GetDevicePath(),
		Measurement: req.GetMeasurement(),
	}, req.GetStartTime(), req.GetEndTime(), limit)
	if err != nil {
		return nil, commons.ToGRPCStatus(commons.Wrap("grpc", commons.CodeInternal, "query", err))
	}

	resp := &pb.QueryRangeResponse{Points: make([]*pb.PointData, 0, len(res.Rows))}
	for _, row := range res.Rows {
		pbVal, err := convert.ToPB(row.Value)
		if err != nil {
			return nil, commons.ToGRPCStatus(commons.Wrap("grpc", commons.CodeInternal, "encode value", err))
		}
		resp.Points = append(resp.Points, &pb.PointData{
			Timestamp: row.Timestamp,
			Value:     pbVal,
		})
	}
	if s.hooks.OnRPCItems != nil {
		s.hooks.OnRPCItems("QueryRange", 1, len(resp.Points))
	}
	return resp, nil
}

func (s *Server) ExecuteSQL(ctx context.Context, req *pb.ExecuteSQLRequest) (*pb.ExecuteSQLResponse, error) {
	begin := time.Now()
	sess, err := s.requireSession(ctx, req.GetSessionId())
	if err != nil {
		return nil, err
	}
	sqlText := strings.TrimSpace(req.GetSql())
	if sqlText == "" {
		s.observeSQLFailure(plan.Kind(-1), "parse", begin)
		return nil, commons.ToGRPCStatus(commons.ErrSQLTextEmpty)
	}

	p, err := queryengine.Plan(sqlText)
	if err != nil {
		s.observeSQLFailure(plan.Kind(-1), "parse", begin)
		return nil, commons.ToGRPCStatus(err)
	}
	info := sess.SessionInfo()
	priv := auth.PrivilegeRead
	if p.NeedsWrite() {
		priv = auth.PrivilegeWrite
	}
	if err := s.auth.CheckPrivilege(info.User.UserID, info.User.Username, p.DevicePath(), priv); err != nil {
		s.observeSQLFailure(p.Kind, "auth", begin)
		return nil, commons.ToGRPCStatus(err)
	}

	res, err := s.sql.ExecutePlan(ctx, p)
	if err != nil {
		return nil, commons.ToGRPCStatus(commons.Wrap("grpc", commons.CodeInternal, "execute sql", err))
	}
	resp, err := sink.SQLToExecuteResponse(res)
	if err != nil {
		return nil, err
	}
	if s.hooks.OnRPCItems != nil {
		s.hooks.OnRPCItems("ExecuteSQL", 1, sqlResponseItems(resp))
	}
	return resp, nil
}

func (s *Server) requireSession(ctx context.Context, sessionID int64) (session.IClientSession, error) {
	if sessionID <= 0 {
		return nil, commons.ToGRPCStatus(commons.ErrGRPCSessionIDInvalid)
	}
	sess, err := s.mgr.Require(sessionID)
	if err != nil {
		return nil, commons.ToGRPCStatus(err)
	}
	_ = ctx
	return sess, nil
}

func peerHostPort(ctx context.Context) (host string, port int) {
	host = "unknown"
	p, ok := peer.FromContext(ctx)
	if !ok || p.Addr == nil {
		return host, 0
	}
	h, ps, err := net.SplitHostPort(p.Addr.String())
	if err != nil {
		return p.Addr.String(), 0
	}
	host = h
	if strings.HasPrefix(host, "[") && strings.HasSuffix(host, "]") {
		host = host[1 : len(host)-1]
	}
	port, _ = strconv.Atoi(ps)
	return host, port
}

func (s *Server) observeSQLFailure(kind plan.Kind, errorClass string, begin time.Time) {
	if s.hooks.OnSQL != nil {
		s.hooks.OnSQL(kind, false, errorClass, time.Since(begin))
	}
}

func sqlResponseItems(resp *pb.ExecuteSQLResponse) int {
	if resp == nil {
		return 0
	}
	if len(resp.Rows) > 0 {
		return len(resp.Rows)
	}
	if len(resp.CatalogRows) > 0 {
		return len(resp.CatalogRows)
	}
	if resp.GetAffectedRows() > 0 {
		return int(resp.GetAffectedRows())
	}
	return 0
}
