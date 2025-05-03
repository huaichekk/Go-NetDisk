package controller

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-netdisk/common"
	"go-netdisk/models"
	"go-netdisk/utils"
	"gorm.io/gorm"
	"net/http"
	"os"
)

const STORAGE_PATH = "./uploads"

type FileData struct {
	FileName string `json:"file_name"`
	FileType string `json:"file_type"`
	FileSize int    `json:"file_size"`
	FileData string `json:"file_data"` // Base64编码的文件内容
}

func UploadFile(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		utils.FailWithMessage(c, http.StatusBadRequest, "用户不存在")
		return
	}
	var fileData FileData
	if err := c.ShouldBindJSON(&fileData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的JSON格式"})
		return
	}
	decoded, err := base64.StdEncoding.DecodeString(fileData.FileData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "文件内容解码失败"})
		return
	}
	hash := md5.New()
	hash.Write(decoded)
	sum := hash.Sum(nil)
	md5Str := hex.EncodeToString(sum)
	//如果hash值存在直接触发秒传
	var existfile models.File
	err = common.DB().Where("hash = ?", md5Str).First(&existfile).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 写入文件
			filePath := fmt.Sprintf("%s/%s", STORAGE_PATH, md5Str)
			if err := os.WriteFile(filePath, decoded, 0644); err != nil {
				utils.FailWithMessage(c, http.StatusInternalServerError,
					fmt.Sprintf("文件保存失败:%s", err.Error()))
				return
			}
			var file models.File
			file.Hash = md5Str
			file.Size = int64(fileData.FileSize)
			file.Name = fileData.FileName
			if err := common.DB().Create(&file).Error; err != nil {
				utils.FailWithMessage(c, http.StatusBadGateway, err.Error())
				return
			}
			fmt.Println("触发秒传")
		} else {
			utils.FailWithMessage(c, http.StatusBadGateway, err.Error())
			return
		}
	}

	var file1 models.File
	if err := common.DB().Where("hash=?", md5Str).First(&file1).Error; err != nil {
		utils.FailWithMessage(c, http.StatusBadRequest, err.Error())
		return
	}
	var userFile models.UserFile
	userFile.UserID = int(userID.(uint))
	userFile.FileID = int(file1.ID)
	userFile.FatherFileID = -1
	userFile.Type = 1
	if err := common.DB().Create(&userFile).Error; err != nil {
		utils.FailWithMessage(c, http.StatusBadRequest, err.Error())
		return
	}
	utils.OkWithData(c, nil, "Success")
}
