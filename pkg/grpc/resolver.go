package grpc

import (
	"github.com/keington/go-templet/pkg/zlog"
	"go.etcd.io/etcd/client/v3/naming/resolver"
	gresolver "google.golang.org/grpc/resolver"
)

/**
 * @author: x.gallagher.anderson@gmail.com
 * @time: 2023/11/22 0:02
 * @file: resolver.go
 * @description: 客户端服务发现
 */

// NewResolver 创建一个Resolver用于客户端服务发现
func (naming *NamingService) NewResolver() (gresolver.Builder, error) {
	etcdResolver, err := resolver.NewBuilder(naming.Client)
	if err != nil {
		zlog.Fatalf("Etcd resolver error %s", err)
		return nil, err
	}
	return etcdResolver, nil
}
