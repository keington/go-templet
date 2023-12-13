package interceptor

import (
	"runtime"

	"github.com/gin-gonic/gin"
	"github.com/keington/go-templet/internal/pkg/httpx"
	"github.com/keington/go-templet/pkg/zlog"
)

/**
 * @author: x.gallagher.anderson@gmail.com
 * @time: 2023/11/23 21:43
 * @file: exception_interceptor.go
 * @description: 异常捕获拦截 gin
 */

func ExceptionInterceptor(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			switch t := r.(type) {
			case *httpx.Response:
				zlog.Errorf("panic: %v", t)
				httpx.WithRepErrMsg(c, httpx.InternalError.Code, httpx.InternalError.Msg, c.Request.URL.Path)
			default:
				zlog.Errorf("panic: %v", t)
				// print stack trace for debugging
				buf := make([]byte, 1<<16)
				stackSize := runtime.Stack(buf, true)
				zlog.Errorf("panic: %s", buf[:stackSize])
				httpx.WithRepErrMsg(c, httpx.InternalError.Code, httpx.InternalError.Msg, c.Request.URL.Path)
			}
			c.Abort()
		}
	}()
	c.Next()
}
