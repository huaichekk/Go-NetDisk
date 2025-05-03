package controller

import (
	"github.com/gin-gonic/gin"
	"go-netdisk/common"
	"go-netdisk/models"
	"go-netdisk/utils"
	"net/http"
)

func Register(ctx *gin.Context) {
	var user models.User
	user.Username = ctx.PostForm("username")
	user.Password = ctx.PostForm("password")
	if err := common.DB().Create(&user).Error; err != nil {
		utils.FailWithMessage(ctx, http.StatusBadRequest, err.Error())
		return
	} else {
		utils.OkWithData(ctx, nil, "Success")
		return
	}
}
