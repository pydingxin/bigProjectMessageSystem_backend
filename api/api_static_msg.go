package api

import (
	"demo_backend/model"
	"demo_backend/tool"

	"github.com/gogf/gf/v2/net/ghttp"

	"github.com/gogf/gf/v2/errors/gerror"

	"github.com/gogf/gf/v2/frame/g"
)

func handler_api_staticmsg_create(r *ghttp.Request) {
	// 管理员创建项目信息，返回项目信息
	//step1 获取&校验
	var in *model.ProjectStaticMsg
	if err := r.Parse(&in); err != nil {
		r.Response.WriteJsonExit(g.Map{"status": false, "msg": gerror.Wrap(err, "handler_api_staticmsg_create")})
	}

	//step2 判断是否为管理员
	accountid := r.Session.MustGet("accountId").Uint() //获取session里账号id
	if !model.IsAccountIdAdministrator(accountid) {
		r.Response.WriteJsonExit(g.Map{"status": false, "msg": "handler_api_staticmsg_create 您不是管理员，无法管理项目"})
	}

	// step3 操作数据

	result := tool.GetGormConnection().Create(in)
	if result.Error != nil {
		r.Response.WriteJsonExit(g.Map{"status": false, "msg": gerror.Wrap(result.Error, "handler_api_staticmsg_create")})
	} else if result.RowsAffected == 1 {
		r.Response.WriteJsonExit(g.Map{"status": true, "msg": "创建项目成功", "data": in})
	} else {
		r.Response.WriteJsonExit(g.Map{"status": false, "msg": "handler_api_staticmsg_create 创建项目失败"})
	}
}

func handler_api_staticmsg_edit(r *ghttp.Request) {
	// 管理员编辑项目信息，返回项目信息
	//step1 获取&校验
	var in *model.ProjectStaticMsg
	if err := r.Parse(&in); err != nil {
		r.Response.WriteJsonExit(g.Map{"status": false, "msg": gerror.Wrap(err, "handler_api_staticmsg_edit")})
	}
	if in.ID == 0 {
		// id为零会创建新数据
		r.Response.WriteJsonExit(g.Map{"status": false, "msg": "handler_api_staticmsg_edit：缺失项目key"})
	}

	//step2 判断是否为管理员
	accountid := r.Session.MustGet("accountId").Uint() //获取session里账号id
	if !model.IsAccountIdAdministrator(accountid) {
		r.Response.WriteJsonExit(g.Map{"status": false, "msg": "handler_api_staticmsg_edit 您不是管理员，无法管理项目"})
	}

	// step3 操作数据
	result := tool.GetGormConnection().Save(in)
	if result.Error != nil {
		r.Response.WriteJsonExit(g.Map{"status": false, "msg": gerror.Wrap(result.Error, "handler_api_staticmsg_edit")})
	} else if result.RowsAffected == 1 {
		r.Response.WriteJsonExit(g.Map{"status": true, "msg": "编辑项目成功", "data": in})
	} else {
		r.Response.WriteJsonExit(g.Map{"status": false, "msg": "handler_api_staticmsg_edit 编辑项目失败"})
	}
}

func handler_api_staticmsg_delete(r *ghttp.Request) {
	// 管理员删除项目信息，输入key，输出true/false 不输出data
	projectKey := r.Get("key").Uint()

	//判断是否为管理员
	accountid := r.Session.MustGet("accountId").Uint() //获取session里账号id
	if !model.IsAccountIdAdministrator(accountid) {
		r.Response.WriteJsonExit(g.Map{"status": false, "msg": "handler_api_staticmsg_delete 您不是管理员，无法管理项目"})
	}

	// 操作数据
	result := tool.GetGormConnection().Delete(&model.ProjectStaticMsg{}, projectKey)
	if result.Error != nil {
		r.Response.WriteJsonExit(g.Map{"status": false, "msg": gerror.Wrap(result.Error, "handler_api_staticmsg_delete")})
	} else if result.RowsAffected == 1 {
		r.Response.WriteJsonExit(g.Map{"status": true, "msg": "删除项目信息成功"})
	} else {
		r.Response.WriteJsonExit(g.Map{"status": false, "msg": "handler_api_staticmsg_delete:未删除任何项目信息"})
	}
}

func handler_api_staticmsg_getall(r *ghttp.Request) {
	// 获取所有项目信息

	//判断是否为管理员
	accountid := r.Session.MustGet("accountId").Uint() //获取session里账号id
	if !model.IsAccountIdAdministrator(accountid) {
		r.Response.WriteJsonExit(g.Map{"status": false, "msg": "handler_api_staticmsg_delete 您不是管理员，无法管理项目"})
	}

	// 操作数据
	var all_static_msgs []model.ProjectStaticMsg
	result := tool.GetGormConnection().Find(&all_static_msgs)
	if result.Error != nil {
		r.Response.WriteJsonExit(g.Map{"status": false, "msg": gerror.Wrap(result.Error, "handler_api_staticmsg_getall")})
	} else {
		r.Response.WriteJsonExit(g.Map{"status": true, "data": all_static_msgs})
	}
}

func RouterGroup_StaticMsg(group *ghttp.RouterGroup) {

	group.POST("/create", handler_api_staticmsg_create) ///api/staticmsg/create
	group.POST("/edit", handler_api_staticmsg_edit)     ///api/staticmsg/edit
	group.POST("/delete", handler_api_staticmsg_delete) ///api/staticmsg/delete
	group.POST("/getall", handler_api_staticmsg_getall) ///api/staticmsg/delete

}
