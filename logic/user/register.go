package user

import (
	"DouYin/dao/mysql"
	"DouYin/models/api"
)

func Register(username, password string) (api *api.RegisterModel) {
	if mysql.Register(username, password) == true {
		api.Token = "" //鉴权中间件
		api.StatusCode = 0
		api.StatusMsg = "successful"
		api.UserID = 0 //雪花id
		return api
	} else {
		api.StatusCode = -1
		api.StatusMsg = "failed"
		return api
	}
}
