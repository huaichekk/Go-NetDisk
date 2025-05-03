package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUserID(ctx *gin.Context) int {
	userID, exists := ctx.Get("userID")
	if !exists {
		FailWithMessage(ctx, http.StatusBadRequest, "用户不存在")
		ctx.Abort()
	}
	return int(userID.(uint))
}
