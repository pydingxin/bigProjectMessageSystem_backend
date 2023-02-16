package api

import (
	"demo_backend/model"
	"demo_backend/tool"

	"demo_backend/middleware"

	"github.com/gogf/gf/v2/net/ghttp"

	"github.com/gogf/gf/v2/errors/gerror"

	"github.com/gogf/gf/v2/frame/g"
)

func handler_api_staticmsg_create(r *ghttp.Request) {
	// 管理员创建项目信息，返回项目信息
	//获取&校验
	var in *model.ProjectStaticMsg
	if err := r.Parse(&in); err != nil {
		r.Response.WriteJsonExit(g.Map{"status": false, "msg": gerror.Wrap(err, "handler_api_staticmsg_create")})
	}

	// 操作数据
	result := tool.GetGormConnection().Create(in)
	if result.Error != nil {
		r.Response.WriteJsonExit(g.Map{"status": false, "msg": gerror.Wrap(result.Error, "handler_api_staticmsg_create")})
	} else if result.RowsAffected == 1 {
		model.Create_dynamicMsg_by_projectId(in.ID) //动态信息
		r.Response.WriteJsonExit(g.Map{"status": true, "msg": "创建项目成功", "data": in})
	} else {
		r.Response.WriteJsonExit(g.Map{"status": false, "msg": "handler_api_staticmsg_create 创建项目失败"})
	}
}

func handler_api_staticmsg_edit(r *ghttp.Request) {
	// 管理员编辑项目信息，返回项目信息
	//获取&校验
	var in *model.ProjectStaticMsg
	if err := r.Parse(&in); err != nil {
		r.Response.WriteJsonExit(g.Map{"status": false, "msg": gerror.Wrap(err, "handler_api_staticmsg_edit")})
	}
	if in.ID == 0 {
		// id为零会创建新数据
		r.Response.WriteJsonExit(g.Map{"status": false, "msg": "handler_api_staticmsg_edit：缺失项目key"})
	}

	// 操作数据
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
	if projectKey == 0 {
		r.Response.WriteJsonExit(g.Map{"status": false, "msg": "handler_api_staticmsg_delete:请输入项目key/id"})
	}

	// 操作数据
	result := tool.GetGormConnection().Delete(&model.ProjectStaticMsg{}, projectKey)
	if result.Error != nil {
		r.Response.WriteJsonExit(g.Map{"status": false, "msg": gerror.Wrap(result.Error, "handler_api_staticmsg_delete")})
	} else if result.RowsAffected == 1 {
		model.Delete_dynamicMsg_by_projectId(projectKey) //删除动态信息及其历史
		model.Delete_dynamicHistory_by_projectId(projectKey)
		r.Response.WriteJsonExit(g.Map{"status": true, "msg": "删除项目信息成功"})
	} else {
		r.Response.WriteJsonExit(g.Map{"status": false, "msg": "handler_api_staticmsg_delete:未删除任何项目信息"})
	}
}

func handler_api_staticmsg_getall(r *ghttp.Request) {
	// 获取所有项目信息
	// 操作数据
	var all_static_msgs []model.ProjectStaticMsg
	result := tool.GetGormConnection().Find(&all_static_msgs)
	if result.Error != nil {
		r.Response.WriteJsonExit(g.Map{"status": false, "msg": gerror.Wrap(result.Error, "handler_api_staticmsg_getall")})
	} else {
		r.Response.WriteJsonExit(g.Map{"status": true, "data": all_static_msgs})
	}
}

func handler_api_staticmsg_getbyid(r *ghttp.Request) {
	// 根据项目key/id获取项目静态信息，不需要管理员
	projectKey := r.Get("key").Uint()
	if projectKey == 0 {
		r.Response.WriteJsonExit(g.Map{"status": false, "msg": "handler_api_staticmsg_getbyid: 请输入项目key/id"})
	}

	var promsg model.ProjectStaticMsg
	result := tool.GetGormConnection().Find(&promsg, projectKey)
	if result.Error != nil {
		r.Response.WriteJsonExit(g.Map{"status": false, "msg": gerror.Wrap(result.Error, "handler_api_staticmsg_getbyid")})
	} else if result.RowsAffected == 1 {
		r.Response.WriteJsonExit(g.Map{"status": true, "data": promsg})
	} else {
		r.Response.WriteJsonExit(g.Map{"status": false, "msg": "handler_api_staticmsg_getbyid:未获取任何项目信息"})
	}
}

func RouterGroup_StaticMsg(group *ghttp.RouterGroup) {

	group.POST("/getbykey", handler_api_staticmsg_getbyid) ///api/staticmsg/getbykey
	group.POST("/getall", handler_api_staticmsg_getall)    ///api/staticmsg/getall

	// 管理员行为
	group.Group("/", func(group *ghttp.RouterGroup) {
		group.Middleware(middleware.MiddlewareIsAdmin)      //管理员操作
		group.POST("/create", handler_api_staticmsg_create) ///api/staticmsg/create
		group.POST("/edit", handler_api_staticmsg_edit)     ///api/staticmsg/edit
		group.POST("/delete", handler_api_staticmsg_delete) ///api/staticmsg/delete
	})

}
