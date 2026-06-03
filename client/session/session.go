package session

import (
	"context"
	"sync"
	"sync/atomic"

	"github.com/hanami/tidets/commons/errors"
)

type Config struct {
	Host                 string
	Port                 int
	Username             string
	Password             string
	Version              Version
	FetchSize            int
	EnableRPCCompression bool
}

type Option func(*Config)

func WithHost(host string) Option         { return func(c *Config) { c.Host = host } }
func WithPort(port int) Option            { return func(c *Config) { c.Port = port } }
func WithUsername(username string) Option { return func(c *Config) { c.Username = username } }
func WithPassword(password string) Option { return func(c *Config) { c.Password = password } }
func WithVersion(v Version) Option        { return func(c *Config) { c.Version = v } }
func WithFetchSize(n int) Option          { return func(c *Config) { c.FetchSize = n } }
func WithEnableRPCCompression(enable bool) Option {
	return func(c *Config) { c.EnableRPCCompression = enable }
}

// New 创建 Session，未传的字段使用默认值（127.0.0.1:5556, root/root）。
func New(opts ...Option) (Session, error) {
	cfg := Config{
		Host:      DefaultHost,
		Port:      DefaultPort,
		Username:  DefaultUsername,
		Password:  DefaultPassword,
		Version:   V_1_0,
		FetchSize: DefaultFetchSize,
	}
	for _, opt := range opts {
		if opt != nil {
			opt(&cfg)
		}
	}

	if cfg.Host == "" {
		return nil, commons.ErrClientSessionHostRequired
	}
	if cfg.Port <= 0 {
		return nil, commons.ErrClientSessionPortInvalid
	}
	if cfg.Username == "" {
		return nil, commons.ErrClientSessionUsernameRequired
	}
	if cfg.Version == "" {
		return nil, commons.ErrClientSessionVersionRequired
	}
	if cfg.FetchSize < 0 {
		cfg.FetchSize = 0
	}

	return &clientSession{
		host:                 cfg.Host,
		port:                 cfg.Port,
		username:             cfg.Username,
		password:             cfg.Password,
		version:              cfg.Version,
		fetchSz:              int64(cfg.FetchSize),
		enableRPCCompression: cfg.EnableRPCCompression,
	}, nil
}

type clientSession struct {
	host     string
	port     int
	username string
	password string
	version  Version

	enableRPCCompression bool
	fetchSz              int64
	open                 int32

	mu         sync.Mutex
	connection *sessionConnection
}

func (s *clientSession) Open(ctx context.Context) error {
	if ctx == nil {
		ctx = context.Background()
	}
	if !atomic.CompareAndSwapInt32(&s.open, 0, 1) {
		return commons.ErrClientSessionAlreadyOpened
	}

	conn, err := openSessionConnection(ctx, dialParams{
		host:                 s.host,
		port:                 s.port,
		username:             s.username,
		password:             s.password,
		version:              s.version,
		enableRPCCompression: s.enableRPCCompression,
		fetchSize:            atomic.LoadInt64(&s.fetchSz),
	})
	if err != nil {
		atomic.StoreInt32(&s.open, 0)
		return err
	}

	s.mu.Lock()
	s.connection = conn
	s.mu.Unlock()
	return nil
}

func (s *clientSession) Close() error {
	if !atomic.CompareAndSwapInt32(&s.open, 1, 0) {
		return nil
	}

	s.mu.Lock()
	conn := s.connection
	s.connection = nil
	s.mu.Unlock()

	if conn != nil {
		return conn.close()
	}
	return nil
}

func (s *clientSession) InsertBatch(ctx context.Context, devicePath string, points []BatchPoint) error {
	if atomic.LoadInt32(&s.open) != 1 {
		return commons.ErrClientSessionNotOpened
	}
	s.mu.Lock()
	conn := s.connection
	s.mu.Unlock()
	if conn == nil {
		return commons.ErrClientSessionConnectionNil
	}
	return conn.insertBatch(ctx, devicePath, points)
}

func (s *clientSession) InsertPoint(ctx context.Context, devicePath, measurement string, timestamp int64, value Value) error {
	if atomic.LoadInt32(&s.open) != 1 {
		return commons.ErrClientSessionNotOpened
	}
	s.mu.Lock()
	conn := s.connection
	s.mu.Unlock()
	if conn == nil {
		return commons.ErrClientSessionConnectionNil
	}
	return conn.insertPoint(ctx, devicePath, measurement, timestamp, value)
}

func (s *clientSession) IsOpen() bool {
	return atomic.LoadInt32(&s.open) == 1
}

func (s *clientSession) ExecuteSQL(ctx context.Context, sql string) (*SQLResult, error) {
	if atomic.LoadInt32(&s.open) != 1 {
		return nil, commons.ErrClientSessionNotOpened
	}
	s.mu.Lock()
	conn := s.connection
	s.mu.Unlock()
	if conn == nil {
		return nil, commons.ErrClientSessionConnectionNil
	}
	return conn.executeSQL(ctx, sql)
}

func (s *clientSession) SessionID() int64 {
	if atomic.LoadInt32(&s.open) != 1 {
		return 0
	}
	s.mu.Lock()
	conn := s.connection
	s.mu.Unlock()
	if conn == nil {
		return 0
	}
	return conn.getSessionID()
}

func (s *clientSession) QueryRange(ctx context.Context, devicePath, measurement string, startTime, endTime int64) ([]Point, error) {
	return s.QueryRangeWithLimit(ctx, devicePath, measurement, startTime, endTime, 0)
}

func (s *clientSession) QueryRangeWithLimit(ctx context.Context, devicePath, measurement string, startTime, endTime int64, limit int) ([]Point, error) {
	if atomic.LoadInt32(&s.open) != 1 {
		return nil, commons.ErrClientSessionNotOpened
	}
	s.mu.Lock()
	conn := s.connection
	s.mu.Unlock()
	if conn == nil {
		return nil, commons.ErrClientSessionConnectionNil
	}
	return conn.queryRange(ctx, devicePath, measurement, startTime, endTime, limit)
}

func (s *clientSession) SetFetchSize(n int) Session {
	if n < 0 {
		n = 0
	}
	atomic.StoreInt64(&s.fetchSz, int64(n))
	return s
}

func (s *clientSession) Host() string     { return s.host }
func (s *clientSession) Port() int        { return s.port }
func (s *clientSession) Username() string { return s.username }
func (s *clientSession) Password() string { return s.password }
func (s *clientSession) FetchSize() int64 { return atomic.LoadInt64(&s.fetchSz) }
func (s *clientSession) Version() Version { return s.version }
