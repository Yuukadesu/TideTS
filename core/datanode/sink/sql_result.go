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
		ColumnNames:  res.ColumnNames,
	}
	if res.Kind == result.KindSelect {
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
	if len(res.CatalogRows) > 0 {
		out.CatalogRows = make([]*pb.SQLCatalogRow, 0, len(res.CatalogRows))
		for _, row := range res.CatalogRows {
			out.CatalogRows = append(out.CatalogRows, &pb.SQLCatalogRow{
				Columns: row.Columns,
			})
		}
	}
	return out, nil
}
