package middleware

import (
	"demo_backend/model"
	"demo_backend/tool"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func IsAccountIdAdministrator(hisid uint) bool {
	// 判断当前账号是否为管理员用户，目前是根据账号名是否为"admin"
	var user model.Account
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

func MiddlewareIsAdmin(r *ghttp.Request) {
	//判断是否为管理员
	accountid := r.Session.MustGet("accountId").Uint() //获取session里账号id
	if IsAccountIdAdministrator(accountid) {
		r.Middleware.Next()
	} else {
		r.Response.WriteJsonExit(g.Map{"status": false, "msg": "handler_api_staticmsg_delete 您不是管理员"})
	}

}
