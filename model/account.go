package model

import (
	"demo_backend/tool"
)

/*
	用户账号表 Model, org单位名 name账号 pass密码
*/
type Account struct {
	ID   uint   `gorm:"primaryKey" json:"key"`
	Org  string `gorm:"size:40;not null" json:"org"`
	Name string `gorm:"size:40;index:name;not null;unique" json:"name"`
	Pass string `gorm:"size:40;not null" json:"pass"`
}

///handler_api_account_change_password
type Input_ChangePassword struct {
	Passold string `v:"passold@required|length:4,40#请输入旧密码|密码长度为{min}到{max}位"`
	Passnew string `v:"passnew@required|length:4,40#请输入新密码|密码长度为{min}到{max}位"`
}

//handler_api_account_create
type Input_CreateAccount struct {
	Org   string `v:"org@required|length:4,40#请输入单位名|单位名长度为{min}到{max}位"`
	Name  string `v:"name@required|length:4,40#请输入账号|账号长度为{min}到{max}位"`
	Pass  string `v:"pass@required|length:4,40#请输入密码|密码长度为{min}到{max}位"`
	Pass2 string `v:"pass2@required|same:Pass#请确认密码|两次输入密码不同"`
}

// handler_api_account_edit
type Input_EditAccount struct {
	ID   uint
	Org  string `v:"org@required|length:4,40#请输入单位名|单位名长度为{min}到{max}位"`
	Name string `v:"name@required|length:4,40#请输入账号|账号长度为{min}到{max}位"`
	Pass string `v:"pass@required|length:4,40#请输入密码|密码长度为{min}到{max}位"`
}

// handler_api_account_delete
type Input_DeleteAccount struct {
	ID uint
}

func IsAccountIdAdministrator(hisid uint) bool {
	// 判断当前账号是否为管理员用户，目前是根据账号名是否为"admin"
	var user Account
	result := tool.GetGormConnection().Where("id = ?", hisid).First(&user)
	if result.Error != nil {
		panic(result.Error)
	}
	// 找到并且其名字为admin
	if result.RowsAffected == 1 && user.Name == "admin" {
		return true
	} else {
		return false
	}
}

func init() {
	tool.GetGormConnection().AutoMigrate(&Account{})
}
