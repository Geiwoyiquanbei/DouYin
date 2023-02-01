package models

type Users struct {
	Id            string
	Username      string
	Password      string
	Face_image    string
	Nickname      string
	Fans_counts   int
	Follow_counts int
}

func (Users) TableName() string {
	return "users"
}
