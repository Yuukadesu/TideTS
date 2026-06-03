package session

import "time"

// ConnectionType 连接类型；单机阶段仅 gRPC 客户端连接。
type ConnectionType int

const ConnectionGRPC ConnectionType = iota

func (t ConnectionType) String() string {
	if t == ConnectionGRPC {
		return "GRPC"
	}
	return "UNKNOWN"
}

// ClientVersion 客户端协议版本。
type ClientVersion string

const ClientVersionV10 ClientVersion = "V_1_0"

// UserEntity 权限校验用的用户实体（对齐 IoTDB commons UserEntity 子集）。
type UserEntity struct {
	UserID        int64
	Username      string
	ClientAddress string
}

// SessionInfo 执行层携带的会话上下文（对齐 IoTDB SessionInfo 子集）。
type SessionInfo struct {
	SessionID int64
	User      UserEntity
	Zone      string
	Version   ClientVersion
	FetchSize int64
}

// LoginParams OpenSession / login 入参。
type LoginParams struct {
	Username             string
	Password             string
	ZoneID               string
	ClientVersion        ClientVersion
	FetchSize            int64
	EnableRPCCompression bool
}

// LoginResult login 结果，供 gRPC 层映射状态码。
type LoginResult struct {
	SessionID int64
	UserID    int64
	OK        bool
	Message   string
}

// OpenParams 创建并登录一条 gRPC 会话。
type OpenParams struct {
	Login      LoginParams
	ClientAddr string
	ClientPort int
}

func nowMillis() int64 {
	return time.Now().UnixMilli()
}
