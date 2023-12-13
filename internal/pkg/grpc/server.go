package grpc

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/keington/go-templet/api/manifest/protobuf"
	serve "github.com/keington/go-templet/internal/service/grpc"
	grpcDep "github.com/keington/go-templet/pkg/grpc"
	"github.com/keington/go-templet/pkg/zlog"
	"google.golang.org/grpc"
)

/**
 * @author: x.gallagher.anderson@gmail.com
 * @time: 2023/11/20 23:21
 * @file: server.go
 * @description: grpc server
 */

type Service interface {
}

type Server struct {
	g           *grpc.Server
	lister      string
	endpoint    []string
	username    string
	password    string
	serviceName string
	version     string
	gwAddr      string
	gwPort      int
}

// NewGrpcServer 创建一个新的rpc服务端
func (s *Server) NewGrpcServer() error {

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%v", s.gwAddr, s.gwPort))
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}
	zlog.Infof("core gateway listening at %s", lis.Addr())

	//s := grpc.NewServer()
	protobuf.RegisterSystemServer(s.g, &serve.Server{})

	srv, err := grpcDep.NewRegisterService(s.endpoint, s.username, s.password, s.serviceName)
	if err != nil {
		return fmt.Errorf("failed to create naming serve: %v", err)
	}

	err = srv.AddEndpoint(grpcDep.ServiceMata{
		Address: s.gwAddr,
		Name:    s.serviceName,
		Port:    s.gwPort,
		Version: s.version,
	})
	if err != nil {
		return fmt.Errorf("failed to add endpoint: %v", err)
	}

	go func() {
		if err := s.g.Serve(lis); err != nil {
			zlog.Fatalf("failed to serve: %v", err)
			err := srv.DelEndpoint(s.serviceName)
			if err != nil {
				return
			}
		}
	}()

	quit := make(chan os.Signal)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	zlog.Infof("Shutdown Server ...")
	s.g.GracefulStop()
	srv.DelAllEndpoint()
	zlog.Infof("Server exiting")
	return nil
}
