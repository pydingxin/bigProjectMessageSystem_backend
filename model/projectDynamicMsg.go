package model

import (
	"demo_backend/tool"

	"gorm.io/gorm"
)

/*
	项目动态信息表
	动态信息表是静态信息的辅表，不需要自动id
	每个字段的更新时间保存到动态信息编辑表里
*/
type ProjectDynamicMsg struct {
	gorm.Model
	Projectid  uint
	Lixiang    string //立项
	Yongdi     string //用地
	Guihua     string //规划
	Huanping   string //环评
	Nengping   string //能评
	Xukezheng  string //许可证
	Xingxiang  string //形象进度
	Yearcosted string //年度已投资
}

func init() {
	tool.GetGormConnection().AutoMigrate(&ProjectDynamicMsg{})
}
