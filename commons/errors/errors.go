package commons

// ---------- storage ----------

var (
	ErrStorageInvalidRange              = New("storage", CodeInvalidArgument, "start must be <= end")
	ErrStorageDataDirRequired           = New("storage", CodeInvalidArgument, "dataDir is required")
	ErrStorageEngineClosing             = New("storage", CodeUnavailable, "engine is closing")
	ErrStorageDeviceMeasurementRequired = New("storage", CodeInvalidArgument, "device path and measurement are required")
	ErrStorageTimestampInvalid          = New("storage", CodeInvalidArgument, "timestamp must be positive")
	ErrStorageDataTypeMismatch          = New("storage", CodeInvalidArgument, "data type mismatch for series")
	ErrStorageSchemaRequired            = New("storage", CodeInvalidArgument, "timeseries must be registered in schema before writing")
	ErrStorageMixedDataTypesInChunk     = New("storage", CodeInvalidArgument, "mixed data types in one series chunk")
)

// ErrStorageUnsupportedDataType 不支持的 DataType 枚举值。
func ErrStorageUnsupportedDataType(dt uint8) *Error {
	return Errorf("storage", CodeInvalidArgument, "unsupported data type %d", dt)
}

// ErrStorageUnknownDataType 读取到未知 DataType。
func ErrStorageUnknownDataType(dt uint8) *Error {
	return Errorf("storage", CodeInvalidArgument, "unknown data type %d", dt)
}

// ---------- segment ----------

var (
	ErrSegmentCorrupt        = New("segment", CodeCorrupt, "corrupt file")
	ErrSegmentNothingToFlush = New("segment", CodeInvalidArgument, "nothing to flush")
)

// ErrSegmentUnsupportedVersion .seg 版本与当前代码不一致。
func ErrSegmentUnsupportedVersion(got, want uint32) *Error {
	return Errorf("segment", CodeInvalidArgument,
		"unsupported file version %d (want %d), remove data-dir/segments", got, want)
}

// ---------- wal ----------

var (
	ErrWALCorruptRecord = New("wal", CodeCorrupt, "corrupt record")
	ErrWALApplyRequired = New("wal", CodeInvalidArgument, "apply is required")
)

// ---------- flush ----------

var (
	ErrFlushManagerClosed = New("flush", CodeUnavailable, "manager closed")
	ErrFlushQueueFull     = New("flush", CodeUnavailable, "queue full")
)

// ---------- codec ----------

var ErrCodecStringTooLong = New("codec", CodeInvalidArgument, "string too long")

// ---------- datanode session ----------

var ErrSessionNotFound = New("session", CodeNotFound, "not found or closed")

// ---------- auth ----------

var (
	ErrAuthInvalidPath      = New("auth", CodeInvalidArgument, "path is required")
	ErrAuthPermissionDenied = New("auth", CodePermissionDenied, "permission denied")
)

// ---------- convert / RPC 值 ----------

var (
	ErrValueRequired       = New("convert", CodeInvalidArgument, "value is required")
	ErrValueBoolRequired   = New("convert", CodeInvalidArgument, "bool_value is required")
	ErrValueInt32Required  = New("convert", CodeInvalidArgument, "int32_value is required")
	ErrValueInt64Required  = New("convert", CodeInvalidArgument, "int64_value is required")
	ErrValueFloatRequired  = New("convert", CodeInvalidArgument, "float_value is required")
	ErrValueDoubleRequired = New("convert", CodeInvalidArgument, "double_value is required")
	ErrValueTextRequired   = New("convert", CodeInvalidArgument, "text_value is required")
)

// ErrValueUnsupportedDataType proto data_type 不在支持列表。
func ErrValueUnsupportedDataType(v any) *Error {
	return Errorf("convert", CodeInvalidArgument, "unsupported data_type %v", v)
}

// ---------- client session ----------

var (
	ErrClientSessionHostRequired     = New("session", CodeInvalidArgument, "host is required")
	ErrClientSessionPortInvalid      = New("session", CodeInvalidArgument, "port must be positive")
	ErrClientSessionUsernameRequired = New("session", CodeInvalidArgument, "username is required")
	ErrClientSessionVersionRequired  = New("session", CodeInvalidArgument, "version is required")
	ErrClientSessionAlreadyOpened    = New("session", CodeInvalidArgument, "already opened")
	ErrClientSessionNotOpened        = New("session", CodeInvalidArgument, "not opened")
	ErrClientSessionConnectionNil    = New("session", CodeInvalidArgument, "connection is nil")
	ErrClientSessionConnEstablished  = New("session", CodeInvalidArgument, "connection already established")
	ErrClientSessionGRPCNotReady     = New("session", CodeUnavailable, "grpc client is not ready")
)

// ---------- SQL ----------

var (
	ErrSQLTextEmpty                 = New("sql", CodeInvalidArgument, "sql text is empty")
	ErrSQLSyntax                    = New("sql", CodeInvalidArgument, "sql syntax error")
	ErrSQLParseFailed               = New("sql", CodeInvalidArgument, "failed to build ast from parse tree")
	ErrSQLUnsupportedStmt           = New("sql", CodeInvalidArgument, "unsupported statement")
	ErrSQLPathRequired              = New("sql", CodeInvalidArgument, "device path is required")
	ErrSQLMeasurementRequired       = New("sql", CodeInvalidArgument, "measurement is required")
	ErrSQLDeviceMeasurementRequired = New("sql", CodeInvalidArgument, "device path and measurement are required")
	ErrSQLTimestampRequired         = New("sql", CodeInvalidArgument, "timestamp is required")
	ErrSQLTimestampInvalid          = New("sql", CodeInvalidArgument, "timestamp must be positive")
	ErrSQLValueRequired             = New("sql", CodeInvalidArgument, "value is required")
	ErrSQLStringLiteral             = New("sql", CodeInvalidArgument, "invalid string literal")
	ErrSQLInvalidCmpOp              = New("sql", CodeInvalidArgument, "invalid time comparison operator")
	ErrSQLInvalidTimeRange          = New("sql", CodeInvalidArgument, "invalid time range from WHERE clause")
	ErrSQLLimitInvalid              = New("sql", CodeInvalidArgument, "limit must be positive")
	ErrSQLDataTypeInvalid           = New("sql", CodeInvalidArgument, "invalid data type name")
	ErrSQLWhereRequired             = New("sql", CodeInvalidArgument, "WHERE clause is required")
)

// ---------- schema ----------

var (
	ErrSchemaPathRequired     = New("schema", CodeInvalidArgument, "device path and measurement are required")
	ErrSchemaDataTypeRequired = New("schema", CodeInvalidArgument, "data type is required")
	ErrSchemaTimeseriesExists = New("schema", CodeInvalidArgument, "timeseries already exists")
	ErrSchemaDataTypeMismatch = New("schema", CodeInvalidArgument, "data type mismatch for timeseries")
	ErrSchemaMlogCorrupt      = New("schema", CodeCorrupt, "corrupt metadata log")
	ErrSchemaSnapshotCorrupt  = New("schema", CodeCorrupt, "corrupt metadata snapshot")
)

// ErrSchemaSnapshotUnsupported snapshot 版本与当前代码不一致。
func ErrSchemaSnapshotUnsupported(got, want uint32) *Error {
	return Errorf("schema", CodeInvalidArgument,
		"unsupported snapshot version %d (want %d), remove data-dir/system/schema", got, want)
}

// ---------- metadata ----------

var (
	ErrMetadataPathRequired = New("metadata", CodeInvalidArgument, "path is required")
	ErrMetadataInvalidPath  = New("metadata", CodeInvalidArgument, "path must start with root")
)

// ---------- gRPC API 参数（DataNode） ----------

var (
	ErrGRPCUsernameRequired          = New("grpc", CodeInvalidArgument, "username is required")
	ErrGRPCSessionIDInvalid          = New("grpc", CodeInvalidArgument, "session_id must be positive")
	ErrGRPCDevicePathRequired        = New("grpc", CodeInvalidArgument, "device_path is required")
	ErrGRPCDeviceMeasurementRequired = New("grpc", CodeInvalidArgument, "device_path and measurement are required")
	ErrGRPCTimestampInvalid          = New("grpc", CodeInvalidArgument, "timestamp must be positive")
	ErrGRPCPointsEmpty               = New("grpc", CodeInvalidArgument, "points must not be empty")
	ErrGRPCMeasurementRequired       = New("grpc", CodeInvalidArgument, "measurement is required")
	ErrGRPCTimeRangeInvalid          = New("grpc", CodeInvalidArgument, "start_time must be <= end_time")
)
