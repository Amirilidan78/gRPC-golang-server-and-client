package grpcServers

import (
	"context"
	"go-grpc/pkg/logs"
)

type LogServer struct {
	logs.UnimplementedLogServiceServer
}

func (l *LogServer) WriteLog(ctx context.Context, req *logs.LogRequest) (*logs.LogResponse, error) {

	input := req.GetLogEntry()

	res := &logs.LogResponse{
		Result: input.Name + "-" + input.Data,
	}

	return res, nil
}
