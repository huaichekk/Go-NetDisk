package controller

import (
	"github.com/gin-gonic/gin"
	"go-netdisk/utils"
)

func AuthPing(ctx *gin.Context) {
	utils.OkWithData(ctx, nil, "success")
}
