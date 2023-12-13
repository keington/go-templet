package grpc

import (
	"context"
	"fmt"
	"time"

	"github.com/keington/go-templet/api/manifest/protobuf"
	"github.com/keington/go-templet/pkg/zlog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

/**
 * @author: x.gallagher.anderson@gmail.com
 * @time: 2023/11/22 0:00
 * @file: client.go
 * @description:
 */

func NewGrpcClient(endpoint []string, username, password, serviceName, token string) error {

	// 注册服务到etcd
	// todo: 连接etcd做tls加密
	srv, err := NewRegisterService(endpoint, username, password, serviceName)
	if err != nil {
		return fmt.Errorf("failed to register srv: %v", err)
	}

	resolver, err := srv.NewResolver()
	if err != nil {
		return fmt.Errorf("failed to create resolver: %v", err)
	}

	// 连接服务端
	// todo: 连接服务端做tls加密
	serve := fmt.Sprintf("etcd:///%s", serviceName)
	conn, err := grpc.Dial(serve, grpc.WithResolvers(resolver),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		// 负载均衡策略为轮询
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`))
	if err != nil {
		return fmt.Errorf("failed to dial: %v", err)
	}

	defer func() {
		if err := conn.Close(); err != nil {
			zlog.Errorf("Error while closing connection: %v", err)
		}
	}()

	c := protobuf.NewSystemClient(conn)
	time.Sleep(1 * time.Minute)
	num := 100
	for i := 0; i < num; i++ {
		// 创建2秒超时ctx
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		cancel() // 在每次迭代完成后手动调用取消函数释放资源
		// 创建metadata
		md := metadata.Pairs("authorization", token)
		ctx = metadata.NewOutgoingContext(ctx, md)

		// zlog.Debugf("Initiate RPC request: %d", i)
		res, err := c.HeartBeat(ctx, &protobuf.HeartBeatReq{Message: "ping"})

		if err != nil {
			return fmt.Errorf("RPC request failed: %v", err)
		}

		zlog.Debugf("request result: %s", res.Message)
	}
	return err
}
