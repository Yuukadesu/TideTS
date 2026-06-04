// Package convert 在 gRPC TSValue 与 tsmodel.Value 之间转换（客户端与服务端共用）。
package convert

import (
	"fmt"

	"github.com/hanami/tidets/commons/errors"
	"github.com/hanami/tidets/core/tsmodel"
	pb "github.com/hanami/tidets/protocol/grpc-datanode/pb"
)

// FromPB 将 RPC 值转为存储层 Value。
func FromPB(v *pb.TSValue) (tsmodel.Value, error) {
	if v == nil {
		return tsmodel.Value{}, commons.ErrValueRequired
	}
	switch v.GetDataType() {
	case pb.TSDataType_TS_DATA_TYPE_BOOLEAN:
		b, ok := v.GetValue().(*pb.TSValue_BoolValue)
		if !ok {
			return tsmodel.Value{}, commons.ErrValueBoolRequired
		}
		return tsmodel.NewBoolean(b.BoolValue), nil
	case pb.TSDataType_TS_DATA_TYPE_INT32:
		x, ok := v.GetValue().(*pb.TSValue_Int32Value)
		if !ok {
			return tsmodel.Value{}, commons.ErrValueInt32Required
		}
		return tsmodel.NewInt32(x.Int32Value), nil
	case pb.TSDataType_TS_DATA_TYPE_INT64:
		x, ok := v.GetValue().(*pb.TSValue_Int64Value)
		if !ok {
			return tsmodel.Value{}, commons.ErrValueInt64Required
		}
		return tsmodel.NewInt64(x.Int64Value), nil
	case pb.TSDataType_TS_DATA_TYPE_FLOAT:
		x, ok := v.GetValue().(*pb.TSValue_FloatValue)
		if !ok {
			return tsmodel.Value{}, commons.ErrValueFloatRequired
		}
		return tsmodel.NewFloat(x.FloatValue), nil
	case pb.TSDataType_TS_DATA_TYPE_DOUBLE:
		x, ok := v.GetValue().(*pb.TSValue_DoubleValue)
		if !ok {
			return tsmodel.Value{}, commons.ErrValueDoubleRequired
		}
		return tsmodel.NewDouble(x.DoubleValue), nil
	case pb.TSDataType_TS_DATA_TYPE_TEXT:
		x, ok := v.GetValue().(*pb.TSValue_TextValue)
		if !ok {
			return tsmodel.Value{}, commons.ErrValueTextRequired
		}
		return tsmodel.NewText(x.TextValue), nil
	default:
		return tsmodel.Value{}, commons.ErrValueUnsupportedDataType(v.GetDataType())
	}
}

// ToPB 将存储层 Value 转为 RPC 值。
func ToPB(v tsmodel.Value) (*pb.TSValue, error) {
	if err := v.Validate(); err != nil {
		return nil, fmt.Errorf("convert: %w", err)
	}
	out := &pb.TSValue{DataType: dataTypeToPB(v.Type)}
	switch v.Type {
	case tsmodel.DataTypeBoolean:
		out.Value = &pb.TSValue_BoolValue{BoolValue: v.Boolean}
	case tsmodel.DataTypeInt32:
		out.Value = &pb.TSValue_Int32Value{Int32Value: v.Int32}
	case tsmodel.DataTypeInt64:
		out.Value = &pb.TSValue_Int64Value{Int64Value: v.Int64}
	case tsmodel.DataTypeFloat:
		out.Value = &pb.TSValue_FloatValue{FloatValue: v.Float}
	case tsmodel.DataTypeDouble:
		out.Value = &pb.TSValue_DoubleValue{DoubleValue: v.Double}
	case tsmodel.DataTypeText:
		out.Value = &pb.TSValue_TextValue{TextValue: v.Text}
	}
	return out, nil
}

func dataTypeToPB(dt tsmodel.DataType) pb.TSDataType {
	switch dt {
	case tsmodel.DataTypeBoolean:
		return pb.TSDataType_TS_DATA_TYPE_BOOLEAN
	case tsmodel.DataTypeInt32:
		return pb.TSDataType_TS_DATA_TYPE_INT32
	case tsmodel.DataTypeInt64:
		return pb.TSDataType_TS_DATA_TYPE_INT64
	case tsmodel.DataTypeFloat:
		return pb.TSDataType_TS_DATA_TYPE_FLOAT
	case tsmodel.DataTypeDouble:
		return pb.TSDataType_TS_DATA_TYPE_DOUBLE
	case tsmodel.DataTypeText:
		return pb.TSDataType_TS_DATA_TYPE_TEXT
	default:
		return pb.TSDataType_TS_DATA_TYPE_UNSPECIFIED
	}
}
