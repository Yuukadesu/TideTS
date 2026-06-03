// Package session 提供 TideTS DataNode 的客户端会话 API（对齐 IoTDB Session 用法）。
//
// 典型流程：
//
//	s, err := session.New(session.WithHost("127.0.0.1"), session.WithPort(5556))
//	if err != nil { ... }
//	if err := s.Open(ctx); err != nil { ... }
//	defer s.Close()
//
//	_ = s.InsertPoint(ctx, "root.sg1.d1", "temperature", 100, session.Double(25.5))
//	res, err := s.ExecuteSQL(ctx, `SELECT temperature FROM root.sg1.d1 WHERE time >= 100 AND time <= 200`)
//
// 支持 RPC：InsertPoint、InsertBatch、QueryRange、ExecuteSQL（最简 INSERT/SELECT SQL）。
package session
