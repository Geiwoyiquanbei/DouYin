package main

import (
	"DouYin/dao/mysql"
	"DouYin/router"
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"gopkg.in/ini.v1"
)

func main() {
	conf, err := ini.Load("./conf/conf.ini") //加载配置文件
	if err != nil {
		fmt.Println("ini load failed")
		return
	}
	err = mysql.MysqlInit(conf)
	if err != nil {
		fmt.Println("Mysql init failed")
		return
	}
	r := server.Default()
	r.GET("/hello", func(ctx context.Context, c *app.RequestContext) {
		c.String(consts.StatusOK, "Hello hertz!")
	})
	router.SetUp(r)
	r.Spin()
}
