package api

import (
	"demo_backend/model"
	"demo_backend/tool"

	"github.com/gogf/gf/v2/net/ghttp"

	"github.com/gogf/gf/v2/errors/gerror"

	"github.com/gogf/gf/v2/frame/g"
)

func handler_api_dynamicmsg_submit(r *ghttp.Request) {
	/*
		提报项目一条动态信息
		输入信息：key,field, content，content可能是json，可能是数字，都保存为字符串
		返回保存的DynamicHistory
	*/
	item := model.DynamicHistory{
		Projectid: r.Get("key").Uint(),
		Field:     r.Get("field").String(),
		Content:   r.Get("content").String(),
		Accountid: r.Session.MustGet("accountId").Uint(),
	}

	result := tool.GetGormConnection().Create(&item)
	if result.Error != nil {
		r.Response.WriteJsonExit(g.Map{"status": false, "msg": gerror.Wrap(result.Error, "handler_api_dynamicmsg_submit")})
	} else if result.RowsAffected == 1 {
		model.Save_dynamicHistory_to_ProjectDynamicMsg(r) //也保存到ProjectDynamicMsg中
		r.Response.WriteJsonExit(g.Map{"status": true, "msg": "提报成功", "data": item})
	} else {
		r.Response.WriteJsonExit(g.Map{"status": false, "msg": "handler_api_dynamicmsg_submit 提报失败"})
	}
}

func handler_api_dynamicmsg_getbyid(r *ghttp.Request) {
	// 获取项目的完整动态信息

	projectKey := r.Get("key").Uint()
	if projectKey == 0 {
		r.Response.WriteJsonExit(g.Map{"status": false, "msg": "handler_api_dynamicmsg_getbyid: 请输入项目key/id"})
	}

	var pdm model.ProjectDynamicMsg
	result := tool.GetGormConnection().Where("projectid = ?", projectKey).Find(&pdm)

	if result.Error != nil {
		r.Response.WriteJsonExit(g.Map{"status": false, "msg": gerror.Wrap(result.Error, "handler_api_dynamicmsg_getbyid")})
	} else if result.RowsAffected == 1 {
		r.Response.WriteJsonExit(g.Map{"status": true, "data": pdm})
	} else {
		r.Response.WriteJsonExit(g.Map{"status": false, "msg": "handler_api_dynamicmsg_getbyid 无结果"})
	}
}
func handler_api_dynamicmsg_submitHistory(r *ghttp.Request) {
	// 获取项目的某域的编辑历史
	projectKey := r.Get("key").Uint()
	field := r.Get("field").String()
	var history []model.DynamicHistory
	result := tool.GetGormConnection().Where("projectid = ? and field = ?", projectKey, field).Find(&history)
	if result.Error != nil {
		r.Response.WriteJsonExit(g.Map{"status": false, "msg": gerror.Wrap(result.Error, "handler_api_dynamicmsg_submitHistory")})
	} else {
		r.Response.WriteJsonExit(g.Map{"status": true, "data": history})
	}

}

func RouterGroup_DynamicMsg(group *ghttp.RouterGroup) {
	group.POST("/submit", handler_api_dynamicmsg_submit)               ///api/dynamicmsg/submit
	group.POST("/getbyid", handler_api_dynamicmsg_getbyid)             ///api/dynamicmsg/getbyid
	group.POST("/submitHistory", handler_api_dynamicmsg_submitHistory) ///api/dynamicmsg/submitHistory
}
