package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ApiResponse struct {
	code       int
	statusText string
}

func NotFound(ctx *gin.Context, message string) {
	code := http.StatusNotFound
	ctx.JSON(code, error(message))
}
func BadRequest(ctx *gin.Context, message string) {
	code := http.StatusBadRequest
	ctx.JSON(code, error(message))
}
func OK(ctx *gin.Context, response interface{}) {
	code := http.StatusOK
	ctx.JSON(code, data(response))
}

func Created(ctx *gin.Context, response interface{}) {
	code := http.StatusCreated
	ctx.JSON(code, data(response))
}
func Conflict(ctx *gin.Context, message string) {
	code := http.StatusConflict
	ctx.JSON(code, error(message))
}
func error(message string) *gin.H {
	return &gin.H{
		"error": message,
	}
}
func data(data interface{}) *gin.H {
	return &gin.H{
		"data": data,
	}
}
