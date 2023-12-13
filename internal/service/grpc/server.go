package grpc

import (
	"context"

	"github.com/keington/go-templet/api/manifest/protobuf"
	"github.com/keington/go-templet/pkg/zlog"
)

/**
 * @author: x.gallagher.anderson@gmail.com
 * @time: 2023/10/25 1:16
 * @file: server.go
 * @description: gRPC消息处理
 */

type Server struct {
	protobuf.UnimplementedSystemServer
}

func (s *Server) HeartBeat(ctx context.Context, req *protobuf.HeartBeatReq) (*protobuf.HeartBeatRep, error) {

	zlog.Debugf("receive heartbeat is: %s", req.GetMessage())

	return &protobuf.HeartBeatRep{
		Message: "pong",
	}, nil
}
