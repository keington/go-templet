package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/keington/go-templet/pkg/interceptor"
)

/**
 * @author: x.gallagher.anderson@gmail.com
 * @time: 2023/11/14 21:35
 * @file: controller.go
 * @description: 路由控制器
 */

type Controller struct {
	Ctx *gin.Context
}

func NewController(ctx *gin.Context) *Controller {
	return &Controller{Ctx: ctx}
}

func (c *Controller) SetupRoutes(r *gin.Engine) *gin.Engine {
	r.Use(interceptor.ExceptionInterceptor)

	v1 := r.Group("/api/v1/")
	v1.POST("/login", c.Login)

	return r
}
