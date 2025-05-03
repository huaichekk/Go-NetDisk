package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-netdisk/common"
	"go-netdisk/models"
	"go-netdisk/utils"
	"net/http"
)

func Files(ctx *gin.Context) {
	userID := utils.GetUserID(ctx)
	userFiles := []models.UserFile{}
	if err := common.DB().Debug().
		Where("user_id = ?", userID).
		Where("father_file_id=?", -1).
		Find(&userFiles).Error; err != nil {
		utils.FailWithMessage(ctx, http.StatusBadRequest, err.Error())
		return
	}
	fmt.Println("userfile", userFiles)
	// 收集所有 FileID
	fileIDs := make([]int, 0, len(userFiles))
	for _, uf := range userFiles {
		fileIDs = append(fileIDs, uf.FileID)
	}
	fmt.Println("filesID", fileIDs)
	// 查询所有关联的文件
	files := []models.File{}
	if err := common.DB().Where("id IN (?)", fileIDs).Find(&files).Error; err != nil {
		utils.FailWithMessage(ctx, http.StatusBadRequest, err.Error())
		return
	}
	utils.OkWithData(ctx, files, "Success")
}
