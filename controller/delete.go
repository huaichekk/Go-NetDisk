package controller

import (
	"github.com/gin-gonic/gin"
	"go-netdisk/common"
	"go-netdisk/models"
	"go-netdisk/utils"
	"net/http"
)

func Delete(ctx *gin.Context) {
	userID := utils.GetUserID(ctx)
	fileID := ctx.Query("id")
	if err := common.DB().Where("id = ?", fileID).Delete(&models.File{}).Error; err != nil {
		utils.FailWithMessage(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	if err := common.DB().Where("user_id=?", userID).
		Where("file_id=?", fileID).Delete(&models.UserFile{}).Error; err != nil {
		utils.FailWithMessage(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	utils.OkWithData(ctx, nil, "success")
}
