package session

import "context"

// Version 协议版本（对齐 IoTDB Session.Builder.version）。
type Version string

const (
	V_1_0 Version = "V_1_0"

	DefaultHost      = "127.0.0.1"
	DefaultPort      = 5556
	DefaultUsername  = "root"
	DefaultPassword  = "root"
	DefaultFetchSize = 10000
)

// Session 对外会话接口，用法类似 IoTDB Java Session：
//
//	s, _ := session.New()
//	_ = s.Open(ctx)
//	_ = s.InsertPoint(ctx, "root.sg1.d1", "temperature", time.Now().UnixMilli(), Double(25.5))
//	res, _ := s.ExecuteSQL(ctx, "SELECT temperature FROM root.sg1.d1 WHERE time >= 1 AND time <= 100")
//	_ = s.Close()
type Session interface {
	Open(ctx context.Context) error
	Close() error
	IsOpen() bool

	InsertPoint(ctx context.Context, devicePath, measurement string, timestamp int64, value Value) error
	InsertBatch(ctx context.Context, devicePath string, points []BatchPoint) error
	// QueryRange 按时间范围查询；limit<=0 时使用会话 FetchSize。
	QueryRange(ctx context.Context, devicePath, measurement string, startTime, endTime int64) ([]Point, error)
	QueryRangeWithLimit(ctx context.Context, devicePath, measurement string, startTime, endTime int64, limit int) ([]Point, error)
	ExecuteSQL(ctx context.Context, sql string) (*SQLResult, error)

	SetFetchSize(n int) Session

	// SessionID 服务端分配的会话 ID；未 Open 时为 0。
	SessionID() int64

	Host() string
	Port() int
	Username() string
	Password() string
	FetchSize() int64
	Version() Version
}
