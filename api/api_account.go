package api

import (
	"demo_backend/model"
	"demo_backend/tool"

	"demo_backend/middleware"

	"github.com/gogf/gf/v2/errors/gerror"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func handler_api_account_logout(r *ghttp.Request) {
	// 删除session信息即可 不返回data
	r.Session.RemoveAll()
	r.Response.WriteJsonExit(g.Map{"status": true})
}

func handler_api_account_change_password(r *ghttp.Request) {
	// 用户修改自己的密码（但不能修改账号和单位名） 不返回data
	// step1 获取&校验
	var in *model.Input_ChangePassword
	if err := r.Parse(&in); err != nil {
		r.Response.WriteJsonExit(g.Map{"status": false, "msg": gerror.Wrap(err, `handler_api_account_change_password`)})
	}

	// 根据session里的账号id更新密码
	accountid := r.Session.MustGet("accountId").Uint()

	// step3 操作数据
	var account model.Account
	result2 := tool.GetGormConnection().Where("id = ?", accountid).Find(&account)
	if result2.Error != nil {
		r.Response.WriteJsonExit(g.Map{"status": false, "msg": gerror.Wrap(result2.Error, "handler_api_account_change_password")})
	} else {
		if account.Pass != in.Passold {
			r.Response.WriteJsonExit(g.Map{"status": false, "msg": "旧密码输入错误"})
		}
	}

	result := tool.GetGormConnection().Model(&model.Account{}).Where("id = ?", accountid).Update("pass", in.Passnew)
	if result.Error != nil {
		r.Response.WriteJsonExit(g.Map{"status": false, "msg": gerror.Wrap(result.Error, "handler_api_account_change_password")})
	} else if result.RowsAffected == 1 {
		r.Response.WriteJsonExit(g.Map{"status": true, "msg": "修改密码成功"})
	} else {
		r.Response.WriteJsonExit(g.Map{"status": false, "msg": "handler_api_account_change_password 修改密码失败"})
	}

}

func handler_api_account_create(r *ghttp.Request) {
	/* 管理员创建账号，返回Account，成功则返回：
	{
		"data": {
			"key": 7,
			"org": "平邑县发改局",
			"name": "pyxfgj",
			"pass": "123456"
		},
		"msg": "创建账号成功",
		"status": true
		}
	*/
	//step1 获取&校验
	var in *model.Input_CreateAccount
	if err := r.Parse(&in); err != nil {
		r.Response.WriteJsonExit(g.Map{"status": false, "msg": gerror.Wrap(err, "handler_api_account_create")})
	}

	// step3 操作数据
	account := model.Account{Name: in.Name, Org: in.Org, Pass: in.Pass}
	result := tool.GetGormConnection().Create(&account)
	if result.Error != nil {
		r.Response.WriteJsonExit(g.Map{"status": false, "msg": gerror.Wrap(result.Error, "handler_api_account_create")})
	} else if result.RowsAffected == 1 {
		r.Response.WriteJsonExit(g.Map{"status": true, "msg": "创建账号成功", "data": account})
	} else {
		r.Response.WriteJsonExit(g.Map{"status": false, "msg": "handler_api_account_create 创建账号失败"})
	}
}

func handler_api_account_edit(r *ghttp.Request) {
	/*
		管理员编辑账号信息，根据id修改。成功则返回：
		{
		"data:": {
			"key": 7,
			"org": "平邑县发改局",
			"name": "pyxfgj",
			"pass": "alibaba"
		},
		"msg": "编辑账号成功",
		"status": true
		}
	*/
	//step1 获取&校验
	var in *model.Input_EditAccount
	if err := r.Parse(&in); err != nil {
		r.Response.WriteJsonExit(g.Map{"status": false, "msg": gerror.Wrap(err, "handler_api_account_edit")})
	}

	accountid := r.Session.MustGet("accountId").Uint() //获取session里账号id
	if in.ID == accountid {
		// 不能编辑管理员账户
		r.Response.WriteJsonExit(g.Map{
			"status": false,
			"msg":    "handler_api_account_edit:不可编辑管理员账户，您可以修改自己的密码"})
	}

	// step3 操作数据
	account := model.Account{ID: in.ID, Name: in.Name, Org: in.Org, Pass: in.Pass}
	result := tool.GetGormConnection().Save(&account)
	if result.Error != nil {
		r.Response.WriteJsonExit(g.Map{"status": false, "msg": gerror.Wrap(result.Error, "handler_api_account_edit")})
	} else if result.RowsAffected == 1 {
		r.Response.WriteJsonExit(g.Map{"status": true, "msg": "handler_api_account_edit:编辑账号成功", "data:": account})
	} else {
		r.Response.WriteJsonExit(g.Map{"status": false, "msg": "handler_api_account_edit:编辑账号失败"})
	}
}

func handler_api_account_delete(r *ghttp.Request) {
	// 管理员删除账号信息，根据id修改，不返回data
	//step1 获取&校验
	var in *model.Input_DeleteAccount
	if err := r.Parse(&in); err != nil {
		r.Response.WriteJsonExit(g.Map{"status": false, "msg": gerror.Wrap(err, "handler_api_account_delete")})
	}

	accountid := r.Session.MustGet("accountId").Uint() //获取session里账号id
	if in.ID == accountid {
		// 不能删除管理员账户
		r.Response.WriteJsonExit(g.Map{
			"status": false,
			"msg":    "handler_api_account_edit:不可删除管理员账户"})
	}

	// step3 操作数据
	result := tool.GetGormConnection().Delete(&model.Account{ID: in.ID})
	if result.Error != nil {
		r.Response.WriteJsonExit(g.Map{"status": false, "msg": gerror.Wrap(result.Error, "handler_api_account_delete")})
	} else if result.RowsAffected == 1 {
		r.Response.WriteJsonExit(g.Map{"status": true, "msg": "删除账号成功"})
	} else {
		r.Response.WriteJsonExit(g.Map{"status": false, "msg": "handler_api_account_delete:删除账号失败"})
	}
}

func handler_api_account_getall(r *ghttp.Request) {
	// 管理员获取所有账号信息

	//操作数据
	var allaccounts []model.Account
	result := tool.GetGormConnection().Find(&allaccounts)
	if result.Error != nil {
		r.Response.WriteJsonExit(g.Map{"status": false, "msg": gerror.Wrap(result.Error, "handler_api_account_getall")})
	} else {
		r.Response.WriteJsonExit(g.Map{"status": true, "data": allaccounts})
	}

}

func RouterGroup_ApiAccount(group *ghttp.RouterGroup) {
	// 普通用户行为
	group.POST("/logout", handler_api_account_logout)                   ///api/account/logout
	group.POST("/change_password", handler_api_account_change_password) ///api/account/change_password

	// 管理员行为
	group.Group("/", func(group *ghttp.RouterGroup) {
		group.Middleware(middleware.MiddlewareIsAdmin)
		group.POST("/create", handler_api_account_create) ///api/account/create
		group.POST("/edit", handler_api_account_edit)     ///api/account/edit
		group.POST("/delete", handler_api_account_delete) ///api/account/delete
		group.POST("/getall", handler_api_account_getall) ///api/account/getall
	})

}
