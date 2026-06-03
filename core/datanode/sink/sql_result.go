package sink

import (
	"github.com/hanami/tidets/core/queryengine/result"
	"github.com/hanami/tidets/protocol/convert"
	pb "github.com/hanami/tidets/protocol/grpc-datanode/pb"
)

// SQLToExecuteResponse 将 query.Result 转为 gRPC ExecuteSQLResponse。
func SQLToExecuteResponse(res *result.Result) (*pb.ExecuteSQLResponse, error) {
	if res == nil {
		return &pb.ExecuteSQLResponse{}, nil
	}
	out := &pb.ExecuteSQLResponse{
		AffectedRows: int32(res.AffectedRows),
	}
	if res.Kind != result.KindSelect {
		return out, nil
	}
	out.Rows = make([]*pb.SQLRow, 0, len(res.Rows))
	for _, row := range res.Rows {
		val, err := convert.ToPB(row.Value)
		if err != nil {
			return nil, err
		}
		out.Rows = append(out.Rows, &pb.SQLRow{
			Timestamp: row.Timestamp,
			Value:     val,
		})
	}
	return out, nil
}
