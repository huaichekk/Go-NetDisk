package controller

import (
	"github.com/gin-gonic/gin"
	"go-netdisk/utils"
)

func Ping(ctx *gin.Context) {
	utils.OkWithData(ctx, nil, "pong")
}
