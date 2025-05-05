package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-netdisk/utils"
	"strconv"
)

func Delete(ctx *gin.Context) {
	userID := utils.GetUserID(ctx)
	fileID := ctx.Query("file_id")
	i, _ := strconv.Atoi(fileID)
	fmt.Println("delete:", userID, i)
	if err := utils.RemoveFile(userID, i); err != nil {
		utils.FailServerErr(ctx, err)
		return
	}
	utils.OkWithData(ctx, nil, "success")
}
