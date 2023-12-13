package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/keington/go-templet/internal/models/view"
	httpx2 "github.com/keington/go-templet/internal/pkg/httpx"
)

/**
 * @author: x.gallagher.anderson@gmail.com
 * @time: 2023/11/15 22:53
 * @file: user_controller.go
 * @description: 用户控制器
 */

func (c *Controller) Login(g *gin.Context) {
	req := view.LoginReq{}

	err := g.ShouldBindJSON(&req)
	if err != nil {
		return
	}

	/*pw, err := entity.GetUserByPasswd(req.Username)
	if err != nil {
		return
	}*/
	pw := "21232f297a57a5a743894a0e4a801fc3"
	if req.Password != pw {
		httpx2.WithRepMsg(g, httpx2.Success.Code, httpx2.Success.Msg)
		return
	}
}
