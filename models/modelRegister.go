package models

import "go-netdisk/common"

// ModelRegister 模型注册器
type ModelRegister struct {
	Models []interface{}
}

// Register 注册模型
func (r *ModelRegister) Register(models interface{}) {
	if err := common.DB().AutoMigrate(models); err != nil {
		panic(err)
	}
	r.Models = append(r.Models, models)
}

// GetModels 获取所有注册的模型
func (r *ModelRegister) GetModels() []interface{} {
	return r.Models
}

// 全局模型注册器
var globalRegister = &ModelRegister{}

// RegisterModel 供其他模型调用的注册函数
func RegisterModel(models interface{}) {
	globalRegister.Register(models)
}

// GetRegisteredModels 获取所有已注册模型
func GetRegisteredModels() []interface{} {
	return globalRegister.GetModels()
}
