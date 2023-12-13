package main

import (
	"github.com/gin-gonic/gin"
	"github.com/keington/go-templet/internal/controller"
	"github.com/keington/go-templet/internal/initialize"
	"github.com/keington/go-templet/pkg/tools"
)

/**
 * @author: x.gallagher.anderson@gmail.com
 * @time: 2023/10/24 23:45
 * @file: main.go
 * @description:
 */

func init() {

	printEnv()

}

func main() {
	initialize.Initialize()

	r := gin.Default()
	newController := controller.NewController(nil)
	r = newController.SetupRoutes(r)

	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}

	//r := gin.Default() //启动gin路由，携带基础中间件启动 logger and recovery (crash-free) 中间件
	//v1 := r.Group("/api/v1")
	//v1.POST("/login", func(g *gin.Context) {
	//	req := view.LoginReq{}
	//
	//	err := g.ShouldBindJSON(&req)
	//	if err != nil {
	//		return
	//	}
	//	//pw := "21232f297a57a5a743894a0e4a801fc3"
	//	pw, err := entity.GetUserByPasswd(req.Username)
	//	if err != nil {
	//		return
	//	}
	//	if req.Password != pw.Password {
	//		return
	//	}
	//	httpx.WithRepMsg(g, httpx.Success.Code, httpx.Success.Msg)
	//	return
	//})
	//
	//r.Run(":8080")
}

func printEnv() {
	tools.Init()
}
