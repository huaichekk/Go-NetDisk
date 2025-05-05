package controller

import (
	"github.com/gin-gonic/gin"
	"go-netdisk/utils"
)

func Files(ctx *gin.Context) {
	userID := utils.GetUserID(ctx)
	if files, err := utils.SelectALLFile(userID); err != nil {
		utils.FailServerErr(ctx, err)
		return
	} else {
		utils.OkWithData(ctx, files, "Success")
	}
}
