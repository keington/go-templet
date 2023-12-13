package grpc

import (
	"context"
	"time"

	"github.com/keington/go-templet/internal/pkg/httpx"
	"github.com/keington/go-templet/pkg/net"
	"github.com/keington/go-templet/pkg/zlog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

/**
 * @author: x.gallagher.anderson@gmail.com
 * @time: 2023/11/26 20:29
 * @file: server_interceptor.go
 * @description: grpc 服务端拦截器
 */

// UnaryServerInterceptor returns a new unary server interceptors that performs per-request auth.
func UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		s := time.Now()

		metaData, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, status.Errorf(codes.Unauthenticated, httpx.Unauthorized.Msg)
		}

		// 校验token
		token := metaData.Get("authorization")
		if len(token) == 0 {
			return nil, status.Errorf(codes.Unauthenticated, httpx.TokenEmpty.Msg)
		}
		_, err = httpx.VerifyToken(token[0])
		if err != nil {
			return nil, status.Errorf(codes.Unauthenticated, httpx.TokenInvalid.Msg)
		}

		ip, err := net.GetClientIP(ctx)
		if err != nil {
			return nil, err
		}

		// invoking RPC method
		m, err := handler(ctx, req)
		if err != nil {
			return nil, err
		}

		e := time.Now()
		zlog.Infof("Method: %s, IP: %s, startTime: %v, endTime: %v",
			info.FullMethod, ip, s.Format(time.RFC3339), e.Format(time.RFC3339))

		return m, err
	}
}
