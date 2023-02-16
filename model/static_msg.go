package model

import (
	"demo_backend/tool"

	"gorm.io/gorm"
)

/*
	项目静态信息表：
*/
type ProjectStaticMsg struct {
	//项目id
	ID        uint           `gorm:"primaryKey" json:"key"`
	DeletedAt gorm.DeletedAt `gorm:"index"`

	//序号
	Index uint `json:"index"` //主键必须唯一
	//项目名
	Name string `gorm:"size:256;not null" json:"name" v:"name@required#请输入项目名"`

	// 标签
	//建设性质
	Xingzhi string `gorm:"size:50" json:"xingzhi"`
	//建设级别
	Jibie string `gorm:"size:50" json:"jibie"`
	//建设领域
	Lingyu string `gorm:"size:50" json:"lingyu"`

	// 干系人
	//责任领导
	Leader string `gorm:"size:100" json:"leader"`
	//责任单位id列表 就是Account的id
	Dutyorg []uint `gorm:"size:100;serializer:json" json:"dutyorg"`
	//联系方式
	Contact string `gorm:"size:100" json:"contact"`

	// 整体计划
	//主要建设内容和规模
	Neroguimo string `gorm:"size:1000" json:"neroguimo"`
	//建设单位
	Builder string `gorm:"size:100" json:"builder"`
	//建设地点
	Place string `gorm:"size:100" json:"place"`
	//开工时间
	Kaigong string `gorm:"size:50" json:"kaigong"`
	//竣工时间
	Jungong string `gorm:"size:50" json:"jungong"`

	// 资金回顾
	//资金来源
	Costfrom string `gorm:"size:50" json:"costfrom"`
	//总计划投资，万元整数，无法添加gorm check，可能数据库不支持？
	Allcost uint `json:"allcost"`
	//往年已投资，万元整数
	Hadcost uint `json:"hadcost"`

	// 今年计划
	//今年计划投资，万元整数
	Yearcost uint `json:"yearcost"`
	//今年建设计划
	Yearplan string `gorm:"size:500" json:"yearplan"`
	//时间节点
	Yearnode string `gorm:"size:1000" json:"yearnode"`
}

func init() {
	tool.GetGormConnection().AutoMigrate(&ProjectStaticMsg{})
}
