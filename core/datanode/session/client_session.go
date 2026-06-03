package session

import "fmt"

// ClientSession 面向外部 gRPC 客户端的会话（对齐 IoTDB ClientSession）。
// ClientSession is the only identity for a connection.
type ClientSession struct {
	baseSession
	clientAddr string
	clientPort int
}

// NewClientSession 根据对端地址创建会话，须在 Login 之前注册到 Manager。
func NewClientSession(clientAddr string, clientPort int) *ClientSession {
	if clientAddr == "" {
		clientAddr = "unknown"
	}
	cs := &ClientSession{
		clientAddr: clientAddr,
		clientPort: clientPort,
	}
	cs.TouchActiveTime()
	return cs
}

func (s *ClientSession) ClientAddress() string { return s.clientAddr }
func (s *ClientSession) ClientPort() int       { return s.clientPort }

func (s *ClientSession) ConnectionType() ConnectionType { return ConnectionGRPC }

func (s *ClientSession) ConnectionID() string {
	if s.clientPort > 0 {
		return fmt.Sprintf("%s:%d", s.clientAddr, s.clientPort)
	}
	return s.clientAddr
}

func (s *ClientSession) String() string {
	return formatSessionString(s.ID(), s.Username(), s.ConnectionID())
}

func (s *ClientSession) SessionInfo() SessionInfo { return sessionInfoFrom(s) }
