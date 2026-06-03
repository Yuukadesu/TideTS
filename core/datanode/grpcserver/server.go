// Package grpcserver 实现 DataNode 面向客户端的 gRPC 会话服务。
package grpcserver

import (
	"context"
	"log"
	"net"
	"strconv"
	"strings"

	"github.com/hanami/tidets/commons/errors"
	"github.com/hanami/tidets/core/datanode/auth"
	"github.com/hanami/tidets/core/datanode/session"
	"github.com/hanami/tidets/core/datanode/sink"
	"github.com/hanami/tidets/core/queryengine"
	querystore "github.com/hanami/tidets/core/queryengine/storage"
	"github.com/hanami/tidets/core/storageengine"
	"github.com/hanami/tidets/protocol/convert"
	pb "github.com/hanami/tidets/protocol/grpc-datanode/pb"

	"google.golang.org/grpc/peer"
)

// Server 实现 DataNodeSessionService。
type Server struct {
	pb.UnimplementedDataNodeSessionServiceServer

	engine *storageengine.Engine
	sql    *queryengine.Service
	mgr    *session.Manager
	auth   auth.Checker
}

// New 创建 gRPC 服务实现；engine 由调用方负责 Open/Close。
func New(engine *storageengine.Engine) *Server {
	store := &querystore.EngineAdapter{Engine: engine}
	return &Server{
		engine: engine,
		sql:    queryengine.NewService(store),
		mgr:    session.NewManager(auth.DefaultAuthenticator()),
		auth:   auth.Checker{},
	}
}

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
	if err := s.engine.Insert(storageengine.SeriesKey{
		DevicePath:  req.GetDevicePath(),
		Measurement: req.GetMeasurement(),
	}, storageengine.Point{
		Timestamp: req.GetTimestamp(),
		Value:     val,
	}); err != nil {
		return nil, commons.ToGRPCStatus(commons.Wrap("grpc", commons.CodeInternal, "insert", err))
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
		records = append(records, storageengine.Record{
			Key: storageengine.SeriesKey{
				DevicePath:  req.GetDevicePath(),
				Measurement: pt.GetMeasurement(),
			},
			Point: storageengine.Point{
				Timestamp: pt.GetTimestamp(),
				Value:     val,
			},
		})
	}

	if err := s.engine.InsertBatch(records); err != nil {
		return nil, commons.ToGRPCStatus(commons.Wrap("grpc", commons.CodeInternal, "insert batch", err))
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

	limit := int(req.GetLimit())
	if limit <= 0 {
		limit = int(info.FetchSize)
	}
	if limit <= 0 {
		limit = 10000
	}

	points, err := s.engine.Query(storageengine.SeriesKey{
		DevicePath:  req.GetDevicePath(),
		Measurement: req.GetMeasurement(),
	}, req.GetStartTime(), req.GetEndTime(), limit)
	if err != nil {
		return nil, commons.ToGRPCStatus(commons.Wrap("grpc", commons.CodeInternal, "query", err))
	}

	resp := &pb.QueryRangeResponse{Points: make([]*pb.PointData, 0, len(points))}
	for _, p := range points {
		pbVal, err := convert.ToPB(p.Value)
		if err != nil {
			return nil, commons.ToGRPCStatus(commons.Wrap("grpc", commons.CodeInternal, "encode value", err))
		}
		resp.Points = append(resp.Points, &pb.PointData{
			Timestamp: p.Timestamp,
			Value:     pbVal,
		})
	}
	return resp, nil
}

func (s *Server) ExecuteSQL(ctx context.Context, req *pb.ExecuteSQLRequest) (*pb.ExecuteSQLResponse, error) {
	sess, err := s.requireSession(ctx, req.GetSessionId())
	if err != nil {
		return nil, err
	}
	sqlText := strings.TrimSpace(req.GetSql())
	if sqlText == "" {
		return nil, commons.ToGRPCStatus(commons.ErrSQLTextEmpty)
	}

	p, err := queryengine.Plan(sqlText)
	if err != nil {
		return nil, commons.ToGRPCStatus(err)
	}
	info := sess.SessionInfo()
	priv := auth.PrivilegeRead
	if p.NeedsWrite() {
		priv = auth.PrivilegeWrite
	}
	if err := s.auth.CheckPrivilege(info.User.UserID, info.User.Username, p.DevicePath(), priv); err != nil {
		return nil, commons.ToGRPCStatus(err)
	}

	res, err := s.sql.ExecutePlan(ctx, p)
	if err != nil {
		return nil, commons.ToGRPCStatus(commons.Wrap("grpc", commons.CodeInternal, "execute sql", err))
	}
	return sink.SQLToExecuteResponse(res)
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
