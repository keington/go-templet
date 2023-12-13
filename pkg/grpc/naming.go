package grpc

import (
	"context"
	"fmt"
	"time"

	"github.com/keington/go-templet/pkg/cfg"
	"github.com/keington/go-templet/pkg/zlog"
	v3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/naming/endpoints"
	"go.etcd.io/etcd/pkg/transport"
)

/**
 * @author: x.gallagher.anderson@gmail.com
 * @time: 2023/11/20 22:05
 * @file: naming.go
 * @description: 服务管理
 */

const (
	serviceNamePrefix = "lark"
	leaseSecond       = 600
)

var serviceMap map[string]*ServiceMataInfo

// ServiceMata 服务元数据
type ServiceMata struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Port    int    `json:"port"`
	Version string `json:"version"`
}

// ServiceMataInfo 服务元数据信息
type ServiceMataInfo struct {
	Name    string
	LeaseId v3.LeaseID
	Mata    ServiceMata
}

// NamingService 命名服务结构体
type NamingService struct {
	Endpoints []string
	Name      string
	Target    string
	Client    *v3.Client
	manager   endpoints.Manager
}

// NewRegisterService 新注册服务
func NewRegisterService(endpoint []string, username, password string, serviceName string) (*NamingService, error) {

	target := fmt.Sprintf("%s/%s", serviceNamePrefix, serviceName)

	var (
		client *v3.Client
		err    error
	)

	if cfg.GetBool("etcd.tls.enabled") {
		tlsInfo := transport.TLSInfo{
			TrustedCAFile: cfg.GetString("etcd.tls.ca"),
			CertFile:      cfg.GetString("etcd.tls.cert"),
			KeyFile:       cfg.GetString("etcd.tls.key"),
		}
		tlsConfig, err := tlsInfo.ClientConfig()
		if err != nil {
			zlog.Fatalf("TLS config failed: %v", err)
			return nil, err
		}

		client, err = v3.New(v3.Config{
			Endpoints:   endpoint,
			DialTimeout: 240 * time.Second,
			Username:    username,
			Password:    password,
			TLS:         tlsConfig,
		})
		if err != nil {
			zlog.Fatalf("Connect etcd error: %s", err)
			return nil, err
		}
	} else {
		client, err = v3.New(v3.Config{
			Endpoints:   endpoint,
			DialTimeout: 240 * time.Second,
			Username:    username,
			Password:    password,
		})
		if err != nil {
			zlog.Fatalf("Connect etcd error: %s", err)
			return nil, err
		}
	}

	// etcd的endpoints管理
	manager, err := endpoints.NewManager(client, serviceName)
	if err != nil {
		return nil, err
	}
	ns := NamingService{
		Endpoints: endpoint,
		Name:      serviceName,
		Target:    target,
		manager:   manager,
		Client:    client,
	}
	serviceMap = make(map[string]*ServiceMataInfo)
	return &ns, nil
}

// GetFullServeName 获取服务名称
func (naming *NamingService) GetFullServeName(name string, leaseID v3.LeaseID) string {
	return fmt.Sprintf("%s/%d", name, leaseID)
}

// GetPathServeName 获取服务路径名称
func (naming *NamingService) GetPathServeName(name string) string {
	return fmt.Sprintf("%s", name)
}

// AddEndpoint 添加/注册新的服务
func (naming *NamingService) AddEndpoint(e ServiceMata) error {
	//b, _ := sonic.Marshal(e)
	ep := endpoints.Endpoint{
		Addr: fmt.Sprintf("%s:%d", e.Address, e.Port),
		// todo: 这里有个问题，必须传入字符串
		//Metadata: string(e),
		Metadata: e,
	}
	// 在etcd创建一个续期的lease对象
	lease, err := naming.Client.Grant(context.TODO(), leaseSecond)
	if err != nil {
		return err
	}
	key := naming.GetFullServeName(e.Name, lease.ID)
	// 向etcd注册一个Endpoint并绑定续期
	err = naming.manager.AddEndpoint(context.TODO(), key, ep, v3.WithLease(lease.ID))
	if err != nil {
		return err
	}
	zlog.Infof("Add endpoint success: %s", key)
	// 开启自动续期KeepAlive
	ch, err := naming.Client.KeepAlive(context.TODO(), lease.ID)
	if err != nil {
		return err
	}
	// 这个方法会异步打印出每次续期调用的日志
	go func() {
		for {
			ka := <-ch
			zlog.Debugf("Rev reply from service: %s, ttl: %d", key, ka.TTL)
		}
	}()
	serviceMap[e.Name] = &ServiceMataInfo{Name: e.Name, LeaseId: lease.ID, Mata: e}
	return nil
}

// DelEndpoint 移除一个服务
func (naming *NamingService) DelEndpoint(name string) error {
	eu := serviceMap[name]
	if eu == nil {
		return nil
	}
	err := naming.manager.DeleteEndpoint(context.TODO(), naming.GetFullServeName(name, eu.LeaseId))
	if err != nil {
		zlog.Fatalf("Delete endpoint error %s", err)
		return err
	}
	_, err = naming.Client.Revoke(context.TODO(), eu.LeaseId)
	if err != nil {
		zlog.Fatalf("Revoke lease error %s", err)
		return err
	}
	delete(serviceMap, name)
	zlog.Infof("Delete endpoint [%s] success", name)
	return nil
}

// DelAllEndpoint 移除所有服务
func (naming *NamingService) DelAllEndpoint() {
	for k := range serviceMap {
		err := naming.DelEndpoint(k)
		if err != nil {
			zlog.Fatal("Ignore Failure Continue...")
		}
	}
}

// NewLocalDefNamingService 从本地etcd创建
func NewLocalDefNamingService(serviceName string) (*NamingService, error) {

	return NewRegisterService([]string{"localhost:2379"}, "root", "123456", serviceName)
}
