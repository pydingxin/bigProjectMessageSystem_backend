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
	ID        uint           `gorm:"primaryKey"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	//序号
	Index uint

	//项目名
	Name string

	// 标签
	//建设性质
	Xingzhi string
	//建设级别
	Jibie string
	//建设领域
	Lingyu string

	// 干系人
	//责任领导
	Leader string
	//责任单位id列表 就是Account的id
	Dutyorg []uint `gorm:"serializer:json"`
	//联系方式
	Contact string

	// 整体计划
	//主要建设内容和规模
	Neroguimo string
	//建设单位
	Builder string
	//建设地点
	Place string
	//开工时间
	Kaigong string
	//竣工时间
	Jungong string

	// 资金回顾
	//资金来源
	Costfrom string
	//总计划投资，万元整数
	Allcost uint
	//往年已投资，万元整数
	Hadcost uint

	// 今年计划
	//今年计划投资，万元整数
	Yearcost uint
	//今年建设计划
	Yearplan string
	//时间节点
	Yearnode string
}

func init() {
	tool.GetGormConnection().AutoMigrate(&ProjectStaticMsg{})
}
