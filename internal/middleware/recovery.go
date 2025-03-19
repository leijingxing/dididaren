package middleware

import (
	"dididaren/pkg/logger"
	"dididaren/pkg/response"
	"runtime/debug"

	"github.com/gin-gonic/gin"
)

// Recovery 恢复中间件
func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// 记录堆栈信息
				logger.Error("panic recovered",
					"error", err,
					"stack", string(debug.Stack()),
					"path", c.Request.URL.Path,
					"method", c.Request.Method,
					"client_ip", c.ClientIP(),
				)

				// 返回500错误
				response.InternalServerError(c, "服务器内部错误")
			}
		}()

		c.Next()
	}
}
