package main

import (
	"demo_backend/api"
	"demo_backend/middleware"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func main() {
	s := g.Server()
	// s.EnableHTTPS("./httpsCertification/server.crt", "./httpsCertification/server.key")

	s.Use(middleware.MiddlewareLog)          //业务日志
	s.Use(middleware.MiddlewareErrorHandler) //异常日志
	s.Use(middleware.MiddlewareCORS)         //跨域

	s.Group("/api", func(group *ghttp.RouterGroup) {
		group.Group("/unauth", api.RouterGroup_ApiUnauth) // 不需要登录的接口

		group.Middleware(middleware.MiddlewareAuth)         // 下列为需要登录的接口
		group.Group("/account", api.RouterGroup_ApiAccount) // 账号类，其中有管理员操作

		group.Group("/staticmsg", api.RouterGroup_StaticMsg)   // 账号类，都是管理员操作
		group.Group("/dynamicmsg", api.RouterGroup_DynamicMsg) // 账号类，都是管理员操作
		group.Group("/test", api.RouterGroup_ApiTest)

	})

	s.SetPort(80)
	s.SetIndexFolder(true) //静态文件
	s.SetServerRoot("./static")
	s.Run()
}
