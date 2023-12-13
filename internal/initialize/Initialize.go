package initialize

import (
	"github.com/keington/go-templet/internal/pkg/httpx"
	"github.com/keington/go-templet/pkg/cache"
	"github.com/keington/go-templet/pkg/cfg"
	"github.com/keington/go-templet/pkg/database"
	"github.com/keington/go-templet/pkg/zlog"
)

/**
 * @author: x.gallagher.anderson@gmail.com
 * @time: 2023/11/13 23:10
 * @file: Initialize.go
 * @description: 项目全局初始化
 */

func Initialize() {

	var (
		logConfig zlog.LogConfig
		value     = zlog.Xml("./conf.d/log.xml", &logConfig).(*zlog.LogConfig)
	)

	// initialize log
	// 必须第一个初始化
	zlog.NewLogger(value)

	// initialize config
	config, err := cfg.InitCfg("./conf.d", "config", "toml", cfg.Config{})
	if err != nil {
		panic(err)
	}

	if _, err := database.NewDatabase(config.Database); err != nil {
		panic(err)
	}

	if _, err := cache.NewRedis(config.Redis); err != nil {
		panic(err)
	}

	aToken, rToken, err := httpx.GenToken("test", "test")
	if err != nil {
		return
	}
	zlog.Infof("aToken: %s", aToken)
	zlog.Infof("rToken: %s", rToken)

	// register grpc service
	//grpc.NewGrpcServer(config.GetString("gateway.address")+":"+config.GetString("gateway.port"),
	//	config.GetStringSlice("etcd.endpoint"), config.GetString("etcd.username"), config.GetString("etcd.password"),
	//	config.GetString("gateway.service_name"), consts.Version())
	//zlog.Info("This is client")
	//grpc2.NewGrpcClient(config.GetStringSlice("etcd.endpoint"), config.GetString("etcd.username"), config.GetString("etcd.password"), "lark.core.svc")
	//
	//gin.Default()
}
