package controller

import (
	"github.com/gin-gonic/gin"
	"go-netdisk/common"
	"go-netdisk/models"
	"go-netdisk/utils"
)

func Download(ctx *gin.Context) {
	fileid := ctx.Query("file_id")
	var file models.File
	if err := common.DB().Select([]string{"hash"}).
		Where("id = ?", fileid).First(&file).Error; err != nil {
		utils.FailServerErr(ctx, err)
		return
	}
	filepath := STORAGE_PATH + file.Hash
	ctx.File(filepath)
}
