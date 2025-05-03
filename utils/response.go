package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func OkWithData(ctx *gin.Context, data interface{}, msg string) {
	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"data":    data,
		"message": msg,
	})
}

func FailWithMessage(ctx *gin.Context, code int, msg string) {
	ctx.JSON(code, gin.H{
		"code":    code,
		"data":    nil,
		"message": msg,
	})
}

func FailServerErr(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusInternalServerError, gin.H{
		"code":    500,
		"data":    nil,
		"message": err.Error(),
	})
}
