package router

import (
	"DouYin/controller/base"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func BaseRouters(r *server.Hertz) {
	r.GET("/douyin/feed/", base.FeedController)
	r.POST("/douyin/user/register/", base.RegisterController)
}
