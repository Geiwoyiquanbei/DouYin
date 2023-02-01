package mysql

import (
	"DouYin/models"
)

func Register(username, password string) bool {
	user := models.Users{}
	DB.Where("username = ?", username).Find(&user)
	if len(user.Username) != 0 {
		return false
	} else {
		user.Username = username
		user.Password = password
		DB.Create(&user)
		return true
	}
}
