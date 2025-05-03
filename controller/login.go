package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go-netdisk/common"
	"go-netdisk/middleware"
	"go-netdisk/models"
	"go-netdisk/utils"
	"gorm.io/gorm"
	"net/http"
)

func Login(ctx *gin.Context) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	var user models.User
	if err := common.DB().Where("username = ?", username).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 处理记录不存在的情况
			utils.FailWithMessage(ctx, http.StatusBadRequest, "该用户不存在，请注册")
			return
		} else {
			utils.FailWithMessage(ctx, http.StatusBadRequest, err.Error())
		}
	} else {
		// 处理其他错误
		if user.Password != password { //密码错误
			utils.FailWithMessage(ctx, http.StatusBadRequest, "密码错误，请重新输入")
			return
		} else { //正确
			data := map[string]string{}
			token, err := middleware.GetToken(user)
			if err != nil {
				utils.FailWithMessage(ctx, http.StatusBadRequest, err.Error())
				return
			}
			data["X-Token"] = token
			utils.OkWithData(ctx, data, "success")
			return
		}
	}
}
