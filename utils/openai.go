package utils

import (
	"encoding/json"
	"fmt"
	"github.com/dingdinglz/openai"
	"log"
	"os"
	"strconv"
	"sync"
)

const (
	BaseUrl = "https://api.siliconflow.cn/v1"
	ApiKey  = "sk-hmmthqtpzwgfakcxguwzhpnrizicsjpnwdbkpxtbijzekwzo"
	Model   = "Qwen/Qwen2.5-7B-Instruct"
)

var (
	once   sync.Once
	client *openai.Client
)

func GetAI() *openai.Client {
	if client == nil {
		once.Do(func() {
			client = openai.NewClient(&openai.ClientConfig{
				BaseUrl: BaseUrl,
				ApiKey:  ApiKey,
			})
		})
	}
	return client
}

func ThinkFile(filePath string) (string, error) {
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	msg, err := GetAI().EasyChat(Model,
		"你是一个分析概括文件的机器人，你需要分析我之后即将发给你的文件", "分析概括文件:"+string(fileData))
	if err != nil {
		return "", err
	}
	return msg, nil
}

func ChatMessage(message string, userID int) (string, error) {
	var res string
	GetAI().ChatWithTools(Model, []openai.ToolMessage{
		{Role: "user", Content: message + "(辅助信息：该用户的ID为" + strconv.Itoa(userID) + ",不要在消息回复提及)"},
		{
			Role: "system",
			Content: `你是孔伟杰的项目--Go-NetDisk网盘系统的AI管家，你可以帮助用户管理文件。
当用户要求删除文件时，请按以下步骤操作：
1. 调用 getUserAllFileDetail 获取用户所有文件
2. 根据文件名找到对应的文件ID
3. 使用 removeFile 并传入该文件ID进行删除
请确保始终遵循这个流程！当前用户ID为` + strconv.Itoa(userID),
		},
	}, []openai.ChatToolFunction{
		{
			Type: "function",
			Function: openai.ChatToolFuctionDetail{
				Name: "getUserAllFileDetail",
				Description: "通过传入用户ID获取该用户的所有文件细节，包括文件ID(ID),创建时间(CreatedAt)，" +
					"更新时间(UpdatedAt)，删除时间(DeletedAt)，名称(Name)，大小(Size)，hash值(Hash)等信息",
				Parameters: openai.ChatToolParameters{
					Type: "object",
					Properties: map[string]openai.ChatToolFuctionPropertie{
						"userID": {
							Type:        "string",
							Description: "用户ID",
						},
					},
				},
			},
		},
		{
			Type: "function",
			Function: openai.ChatToolFuctionDetail{
				Name:        "removeFile",
				Description: "通过文件ID删除文件（注意：调用前必须先使用getUserAllFileDetail根据文件名查找对应ID）",
				Parameters: openai.ChatToolParameters{
					Type: "object",
					Properties: map[string]openai.ChatToolFuctionPropertie{
						"userID": {
							Type:        "string",
							Description: "用户ID",
						},
						"fileID": {
							Type:        "string",
							Description: "文件ID",
						},
					},
				},
			},
		},
	}, map[string]func(map[string]interface{}) string{
		"getUserAllFileDetail": func(m map[string]interface{}) string {
			log.Println("[function calling]", "：getUserAllFileDetail")
			userID := m["userID"].(string)
			i, err := strconv.Atoi(userID)
			if err != nil {
				return err.Error()
			}
			file, err := SelectALLFile(i)
			jsonData, err := json.Marshal(file)
			if err != nil {
				return "json编码错误" + err.Error()
			}
			// 转换为字符串
			return string(jsonData)
		},
		"removeFile": func(m map[string]interface{}) string {
			log.Println("[function calling]", "：removeFile")
			userID := m["userID"].(string)
			fileID := m["fileID"].(string)
			user_id, err := strconv.Atoi(userID)
			if err != nil {
				return err.Error()
			}
			file_id, err := strconv.Atoi(fileID)
			if err != nil {
				return err.Error()
			}
			if err = RemoveFile(user_id, file_id); err != nil {
				return "删除失败：" + err.Error()
			} else {
				return "删除成功"
			}
		},
	}, func(s string) {
		fmt.Println("回复内容:", s)
		res = s
	})
	//msg, err := GetAI().EasyChat(Model, "你是一个Go-NetDisk网盘系统的AI管家，你可以帮助用户管理文件", message)
	//if err != nil {
	//	return "", err
	//}
	return res, nil
}
