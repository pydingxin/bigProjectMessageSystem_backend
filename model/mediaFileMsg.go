package model

import (
	"demo_backend/tool"

	"gorm.io/gorm"
)

/*
	媒体文件信息表
*/
type MediaFileMsg struct {
	gorm.Model      //文件id
	Projectid  uint //静态信息表中的项目id
	Filename   string
}

func init() {
	tool.GetGormConnection().AutoMigrate(&MediaFileMsg{})
}
