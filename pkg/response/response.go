package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code:    0,
		Message: "success",
		Data:    data,
	})
}

func Error(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError, Response{
		Code:    1,
		Message: err.Error(),
	})
}

func BadRequest(c *gin.Context, message string) {
	c.JSON(http.StatusBadRequest, Response{
		Code:    1,
		Message: message,
	})
}

func Unauthorized(c *gin.Context) {
	c.JSON(http.StatusUnauthorized, Response{
		Code:    1,
		Message: "未授权",
	})
}

func Forbidden(c *gin.Context) {
	c.JSON(http.StatusForbidden, Response{
		Code:    1,
		Message: "禁止访问",
	})
}
