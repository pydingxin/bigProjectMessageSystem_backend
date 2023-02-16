package model

import (
	"demo_backend/tool"
	"time"
)

/*
	动态信息编辑记录表
*/
type DymamicItemHistoryMsg struct {
	Projectid string
	Fieldname string
	Content   string
	UpdatedAt time.Time
}

func init() {
	tool.GetGormConnection().AutoMigrate(&DymamicItemHistoryMsg{})
}
