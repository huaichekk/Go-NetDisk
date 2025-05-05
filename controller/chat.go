package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-netdisk/utils"
)

func Chat(ctx *gin.Context) {
	msg := ctx.PostForm("message")
	fmt.Println("接收到的信息：", msg)
	message, err := utils.ChatMessage(msg, utils.GetUserID(ctx))
	if err != nil {
		utils.FailServerErr(ctx, err)
		return
	}
	fmt.Println(message)
	utils.OkWithData(ctx, message, "success")
}
