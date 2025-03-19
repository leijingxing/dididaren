package middleware

import (
	"dididaren/pkg/response"
	"fmt"
	"runtime/debug"

	"github.com/gin-gonic/gin"
)

// Recovery 恢复中间件
func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// 打印错误堆栈信息
				fmt.Printf("[Recovery] panic recovered:\n%s\n%s\n",
					err,
					debug.Stack(),
				)

				// 返回500错误
				response.Error(c, fmt.Errorf("Internal Server Error"))
				c.Abort()
			}
		}()
		c.Next()
	}
}
