package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-netdisk/common"
	"go-netdisk/models"
	"go-netdisk/utils"
)

// http://127.0.0.1/files/think?file_id=1
func ThinkFile(ctx *gin.Context) {
	fileid := ctx.Query("file_id")
	var file models.File
	if err := common.DB().Where("id=?", fileid).First(&file).Error; err != nil {
		utils.FailServerErr(ctx, err)
		return
	}
	msg, err := utils.ThinkFile(STORAGE_PATH + file.Hash)
	if err != nil {
		fmt.Println(err)
		utils.FailServerErr(ctx, err)
		return
	}
	utils.OkWithData(ctx, msg, "success")
}
