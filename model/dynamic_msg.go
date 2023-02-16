package model

import (
	"demo_backend/tool"
	"time"

	"github.com/gogf/gf/v2/net/ghttp"
)

/*
	项目动态信息表
	动态信息表是静态信息的辅表，不需要自动id
	每个字段的更新时间保存到动态信息编辑表里
*/

// 立项 用地 规划 环评 能评 许可证 形象进度 年度投资的历史
type DynamicHistory struct {
	Projectid uint      `gorm:"index:projectfield" json:"xmkey"`
	Field     string    `gorm:"index:projectfield;size:20" json:"field"` //字段名，lixaing/yongdi/...
	Content   string    `gorm:"size:500" json:"content"`                 //内容,都保存为字符串
	CreatedAt time.Time `json:"time"`
	Accountid uint      `json:"userkey"`
}

type ProjectDynamicMsg struct {
	Projectid  uint
	Lixiang    string `gorm:"size:200" json:"lixiang"`
	Yongdi     string `gorm:"size:200" json:"yongdi"`
	Guihua     string `gorm:"size:200" json:"guihua"`
	Huanping   string `gorm:"size:200" json:"huanping"`
	Nengping   string `gorm:"size:200" json:"nengping"`
	Xukezheng  string `gorm:"size:200" json:"xukezheng"`
	Xingxiang  string `gorm:"size:500" json:"xingxiang"`
	Yearcosted uint   `json:"yearcosted"`
}

func init() {
	conn := tool.GetGormConnection()
	conn.AutoMigrate(&DynamicHistory{})
	conn.AutoMigrate(&ProjectDynamicMsg{})

}

func Delete_dynamicMsg_by_projectId(id uint) {
	tool.GetGormConnection().Where("Projectid = ?", id).Delete(&ProjectDynamicMsg{})
}
func Create_dynamicMsg_by_projectId(id uint) {
	tool.GetGormConnection().Create(&ProjectDynamicMsg{Projectid: id})
}
func Delete_dynamicHistory_by_projectId(id uint) {
	tool.GetGormConnection().Where("Projectid = ?", id).Delete(&DynamicHistory{})
}

// 当提报一条动态信息到DynamicHistory时，也要保存在ProjectDynamicMsg
func Save_dynamicHistory_to_ProjectDynamicMsg(r *ghttp.Request) {
	projectid := r.Get("key").Uint()
	field := r.Get("field").String()
	content := r.Get("content")

	db := tool.GetGormConnection().Model(&ProjectDynamicMsg{}).Where("projectid = ?", projectid)
	if field == "yearcosted" {
		db.Update(field, content.Uint())
	} else {
		db.Update(field, content.String())
	}

}
