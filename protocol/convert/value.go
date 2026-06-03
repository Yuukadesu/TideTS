// Package convert 在 gRPC TSValue 与存储层 model.Value 之间转换（客户端与服务端共用）。
package convert

import (
	"fmt"

	"github.com/hanami/tidets/commons/errors"
	"github.com/hanami/tidets/core/storageengine/model"
	pb "github.com/hanami/tidets/protocol/grpc-datanode/pb"
)

// FromPB 将 RPC 值转为存储层 Value。
func FromPB(v *pb.TSValue) (model.Value, error) {
	if v == nil {
		return model.Value{}, commons.ErrValueRequired
	}
	switch v.GetDataType() {
	case pb.TSDataType_TS_DATA_TYPE_BOOLEAN:
		b, ok := v.GetValue().(*pb.TSValue_BoolValue)
		if !ok {
			return model.Value{}, commons.ErrValueBoolRequired
		}
		return model.NewBoolean(b.BoolValue), nil
	case pb.TSDataType_TS_DATA_TYPE_INT32:
		x, ok := v.GetValue().(*pb.TSValue_Int32Value)
		if !ok {
			return model.Value{}, commons.ErrValueInt32Required
		}
		return model.NewInt32(x.Int32Value), nil
	case pb.TSDataType_TS_DATA_TYPE_INT64:
		x, ok := v.GetValue().(*pb.TSValue_Int64Value)
		if !ok {
			return model.Value{}, commons.ErrValueInt64Required
		}
		return model.NewInt64(x.Int64Value), nil
	case pb.TSDataType_TS_DATA_TYPE_FLOAT:
		x, ok := v.GetValue().(*pb.TSValue_FloatValue)
		if !ok {
			return model.Value{}, commons.ErrValueFloatRequired
		}
		return model.NewFloat(x.FloatValue), nil
	case pb.TSDataType_TS_DATA_TYPE_DOUBLE:
		x, ok := v.GetValue().(*pb.TSValue_DoubleValue)
		if !ok {
			return model.Value{}, commons.ErrValueDoubleRequired
		}
		return model.NewDouble(x.DoubleValue), nil
	case pb.TSDataType_TS_DATA_TYPE_TEXT:
		x, ok := v.GetValue().(*pb.TSValue_TextValue)
		if !ok {
			return model.Value{}, commons.ErrValueTextRequired
		}
		return model.NewText(x.TextValue), nil
	default:
		return model.Value{}, commons.ErrValueUnsupportedDataType(v.GetDataType())
	}
}

// ToPB 将存储层 Value 转为 RPC 值。
func ToPB(v model.Value) (*pb.TSValue, error) {
	if err := v.Validate(); err != nil {
		return nil, fmt.Errorf("convert: %w", err)
	}
	out := &pb.TSValue{DataType: dataTypeToPB(v.Type)}
	switch v.Type {
	case model.DataTypeBoolean:
		out.Value = &pb.TSValue_BoolValue{BoolValue: v.Boolean}
	case model.DataTypeInt32:
		out.Value = &pb.TSValue_Int32Value{Int32Value: v.Int32}
	case model.DataTypeInt64:
		out.Value = &pb.TSValue_Int64Value{Int64Value: v.Int64}
	case model.DataTypeFloat:
		out.Value = &pb.TSValue_FloatValue{FloatValue: v.Float}
	case model.DataTypeDouble:
		out.Value = &pb.TSValue_DoubleValue{DoubleValue: v.Double}
	case model.DataTypeText:
		out.Value = &pb.TSValue_TextValue{TextValue: v.Text}
	}
	return out, nil
}

func dataTypeToPB(dt model.DataType) pb.TSDataType {
	switch dt {
	case model.DataTypeBoolean:
		return pb.TSDataType_TS_DATA_TYPE_BOOLEAN
	case model.DataTypeInt32:
		return pb.TSDataType_TS_DATA_TYPE_INT32
	case model.DataTypeInt64:
		return pb.TSDataType_TS_DATA_TYPE_INT64
	case model.DataTypeFloat:
		return pb.TSDataType_TS_DATA_TYPE_FLOAT
	case model.DataTypeDouble:
		return pb.TSDataType_TS_DATA_TYPE_DOUBLE
	case model.DataTypeText:
		return pb.TSDataType_TS_DATA_TYPE_TEXT
	default:
		return pb.TSDataType_TS_DATA_TYPE_UNSPECIFIED
	}
}
