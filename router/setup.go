package router

import "github.com/cloudwego/hertz/pkg/app/server"

func SetUp(r *server.Hertz) {
	r.Use()        //鉴权等中间件
	BaseRouters(r) //基础功能路由
}
