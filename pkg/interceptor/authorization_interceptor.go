package interceptor

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/keington/go-templet/internal/pkg/httpx"
)

/**
 * @author: x.gallagher.anderson@gmail.com
 * @time: 2023/11/25 23:22
 * @file: authorization_interceptor.go
 * @description: 鉴权拦截器 gin
 */

// AuthorizationInterceptor 鉴权拦截器
// This function is used as the middleware of gin.
func AuthorizationInterceptor() gin.HandlerFunc {
	return func(c *gin.Context) {
		aToken := c.Request.Header.Get("Authorization")
		if aToken == "" {
			httpx.WithRepErrMsg(c, httpx.AuthorizationEmpty.Code, httpx.AuthorizationEmpty.Msg, c.Request.URL.Path)
			c.Abort()
			return
		}

		// 按空格分割
		parts := strings.SplitN(aToken, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			httpx.WithRepErrMsg(c, httpx.AuthorizationIncorrect.Code, httpx.AuthorizationIncorrect.Msg, c.Request.URL.Path)
			c.Abort()
			return
		}

		mc, err := httpx.VerifyToken(parts[1])
		if err != nil {
			httpx.WithRepErrMsg(c, httpx.TokenInvalid.Code, httpx.Unauthorized.Msg, c.Request.URL.Path)
			c.Abort()
			return
		}

		aToken, rToken, err := httpx.RefreshToken(parts[1], parts[2])
		if err != nil {
			httpx.WithRepErr(c, http.StatusUnauthorized, httpx.Unauthorized.Msg, httpx.Unauthorized.Msg, c.Request.URL.Path)
			c.Abort()
			return
		} else {
			httpx.WithRepJSON(c, gin.H{
				"accessToken":  aToken,
				"refreshToken": rToken,
			})
		}

		c.Set("claims", mc)
		c.Next()
	}
}
