package tsmodel

import (
	"encoding/binary"
	"io"

	"github.com/hanami/tidets/commons/errors"
	"github.com/hanami/tidets/core/tsmodel/codec"
)

// DataType 测点值类型（对齐 IoTDB TSDataType 子集）。
type DataType uint8

const (
	DataTypeUnknown DataType = 0
	DataTypeBoolean DataType = 1
	DataTypeInt32   DataType = 2
	DataTypeInt64   DataType = 3
	DataTypeFloat   DataType = 4
	DataTypeDouble  DataType = 5
	DataTypeText    DataType = 6
)

// Value 单值测点 payload；仅 Type 对应字段有效。
type Value struct {
	Type    DataType
	Boolean bool
	Int32   int32
	Int64   int64
	Float   float32
	Double  float64
	Text    string
}

func NewBoolean(v bool) Value   { return Value{Type: DataTypeBoolean, Boolean: v} }
func NewInt32(v int32) Value    { return Value{Type: DataTypeInt32, Int32: v} }
func NewInt64(v int64) Value    { return Value{Type: DataTypeInt64, Int64: v} }
func NewFloat(v float32) Value  { return Value{Type: DataTypeFloat, Float: v} }
func NewDouble(v float64) Value { return Value{Type: DataTypeDouble, Double: v} }
func NewText(v string) Value    { return Value{Type: DataTypeText, Text: v} }

// Validate 检查类型与字段是否一致。
func (v Value) Validate() error {
	switch v.Type {
	case DataTypeBoolean, DataTypeInt32, DataTypeInt64, DataTypeFloat, DataTypeDouble, DataTypeText:
		return nil
	default:
		return commons.ErrStorageUnsupportedDataType(uint8(v.Type))
	}
}

// Equal 比较两个值（含类型）。
func (v Value) Equal(o Value) bool {
	if v.Type != o.Type {
		return false
	}
	switch v.Type {
	case DataTypeBoolean:
		return v.Boolean == o.Boolean
	case DataTypeInt32:
		return v.Int32 == o.Int32
	case DataTypeInt64:
		return v.Int64 == o.Int64
	case DataTypeFloat:
		return v.Float == o.Float
	case DataTypeDouble:
		return v.Double == o.Double
	case DataTypeText:
		return v.Text == o.Text
	default:
		return false
	}
}

// WritePayload 写入类型字节与 payload（不含 timestamp）。
func (v Value) WritePayload(w io.Writer) error {
	if err := v.Validate(); err != nil {
		return err
	}
	if err := binary.Write(w, binary.LittleEndian, uint8(v.Type)); err != nil {
		return err
	}
	switch v.Type {
	case DataTypeBoolean:
		var b uint8
		if v.Boolean {
			b = 1
		}
		return binary.Write(w, binary.LittleEndian, b)
	case DataTypeInt32:
		return binary.Write(w, binary.LittleEndian, v.Int32)
	case DataTypeInt64:
		return binary.Write(w, binary.LittleEndian, v.Int64)
	case DataTypeFloat:
		return binary.Write(w, binary.LittleEndian, v.Float)
	case DataTypeDouble:
		return binary.Write(w, binary.LittleEndian, v.Double)
	case DataTypeText:
		return codec.WriteString(w, v.Text)
	default:
		return commons.ErrStorageUnsupportedDataType(uint8(v.Type))
	}
}

// ReadValuePayload 读取类型字节与 payload。
func ReadValuePayload(r io.Reader) (Value, error) {
	var dt uint8
	if err := binary.Read(r, binary.LittleEndian, &dt); err != nil {
		return Value{}, err
	}
	return readPayloadBody(r, DataType(dt))
}

func readPayloadBody(r io.Reader, dt DataType) (Value, error) {
	switch dt {
	case DataTypeBoolean:
		var b uint8
		if err := binary.Read(r, binary.LittleEndian, &b); err != nil {
			return Value{}, err
		}
		return NewBoolean(b != 0), nil
	case DataTypeInt32:
		var v int32
		if err := binary.Read(r, binary.LittleEndian, &v); err != nil {
			return Value{}, err
		}
		return NewInt32(v), nil
	case DataTypeInt64:
		var v int64
		if err := binary.Read(r, binary.LittleEndian, &v); err != nil {
			return Value{}, err
		}
		return NewInt64(v), nil
	case DataTypeFloat:
		var v float32
		if err := binary.Read(r, binary.LittleEndian, &v); err != nil {
			return Value{}, err
		}
		return NewFloat(v), nil
	case DataTypeDouble:
		var v float64
		if err := binary.Read(r, binary.LittleEndian, &v); err != nil {
			return Value{}, err
		}
		return NewDouble(v), nil
	case DataTypeText:
		s, err := codec.ReadString(r)
		if err != nil {
			return Value{}, err
		}
		return NewText(s), nil
	default:
		return Value{}, commons.ErrStorageUnknownDataType(uint8(dt))
	}
}

// WriteValuesColumn 将同类型点列的值区写入 chunk/WAL 批量列。
func WriteValuesColumn(w io.Writer, dt DataType, pts []Point) error {
	if len(pts) == 0 {
		return nil
	}
	for _, p := range pts {
		if p.Value.Type != dt {
			return commons.ErrStorageMixedDataTypesInChunk
		}
	}
	switch dt {
	case DataTypeBoolean:
		for _, p := range pts {
			var b uint8
			if p.Value.Boolean {
				b = 1
			}
			if err := binary.Write(w, binary.LittleEndian, b); err != nil {
				return err
			}
		}
	case DataTypeInt32:
		for _, p := range pts {
			if err := binary.Write(w, binary.LittleEndian, p.Value.Int32); err != nil {
				return err
			}
		}
	case DataTypeInt64:
		for _, p := range pts {
			if err := binary.Write(w, binary.LittleEndian, p.Value.Int64); err != nil {
				return err
			}
		}
	case DataTypeFloat:
		for _, p := range pts {
			if err := binary.Write(w, binary.LittleEndian, p.Value.Float); err != nil {
				return err
			}
		}
	case DataTypeDouble:
		for _, p := range pts {
			if err := binary.Write(w, binary.LittleEndian, p.Value.Double); err != nil {
				return err
			}
		}
	case DataTypeText:
		for _, p := range pts {
			if err := codec.WriteString(w, p.Value.Text); err != nil {
				return err
			}
		}
	default:
		return commons.ErrStorageUnsupportedDataType(uint8(dt))
	}
	return nil
}

// ReadValuesColumn 读取同类型值列。
func ReadValuesColumn(r io.Reader, dt DataType, n uint32) ([]Value, error) {
	out := make([]Value, n)
	switch dt {
	case DataTypeBoolean:
		for i := uint32(0); i < n; i++ {
			var b uint8
			if err := binary.Read(r, binary.LittleEndian, &b); err != nil {
				return nil, err
			}
			out[i] = NewBoolean(b != 0)
		}
	case DataTypeInt32:
		for i := uint32(0); i < n; i++ {
			var v int32
			if err := binary.Read(r, binary.LittleEndian, &v); err != nil {
				return nil, err
			}
			out[i] = NewInt32(v)
		}
	case DataTypeInt64:
		for i := uint32(0); i < n; i++ {
			var v int64
			if err := binary.Read(r, binary.LittleEndian, &v); err != nil {
				return nil, err
			}
			out[i] = NewInt64(v)
		}
	case DataTypeFloat:
		for i := uint32(0); i < n; i++ {
			var v float32
			if err := binary.Read(r, binary.LittleEndian, &v); err != nil {
				return nil, err
			}
			out[i] = NewFloat(v)
		}
	case DataTypeDouble:
		for i := uint32(0); i < n; i++ {
			var v float64
			if err := binary.Read(r, binary.LittleEndian, &v); err != nil {
				return nil, err
			}
			out[i] = NewDouble(v)
		}
	case DataTypeText:
		for i := uint32(0); i < n; i++ {
			s, err := codec.ReadString(r)
			if err != nil {
				return nil, err
			}
			out[i] = NewText(s)
		}
	default:
		return nil, commons.ErrStorageUnknownDataType(uint8(dt))
	}
	return out, nil
}
