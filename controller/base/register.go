package base

import (
	"DouYin/logic/user"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"net/http"
)

func RegisterController(c context.Context, ctx *app.RequestContext) {
	username := ctx.Query("username")
	password := ctx.Query("password")
	msg := user.Register(username, password)
	ctx.JSON(http.StatusOK, msg)
}
