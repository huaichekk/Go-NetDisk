package utils

import (
	"fmt"
	"go-netdisk/common"
	"go-netdisk/models"
)

func SelectALLFile(userID int) ([]models.File, error) {
	userFiles := []models.UserFile{}
	if err := common.DB().Debug().
		Where("user_id = ?", userID).
		Where("father_file_id=?", -1).
		Find(&userFiles).Error; err != nil {
		return nil, err
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
		return nil, err
	}
	return files, nil
}

func RemoveFile(userID, fileID int) error {
	if err := common.DB().Where("id = ?", fileID).Delete(&models.File{}).Error; err != nil {
		return err
	}
	if err := common.DB().Where("user_id=?", userID).
		Where("file_id=?", fileID).Delete(&models.UserFile{}).Error; err != nil {
		return err
	}
	return nil
}
