package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type User struct {
	Id int
	UserName string  `orm:"column(user_name)"`
	Name string `orm:"column(name)"`
	PassWord string `orm:"column(pass_word)"`
	Mobile string `orm:"column(mobile)"`
	Email string `orm:"column(email)"`
	Status int `orm:"column(status)"`
	CreateUser string `orm:"column(create_user)"`
	CreateTime time.Time `orm:"column(create_time)"`
	UpdateTime time.Time `orm:"column(update_time)"`
}


func init()  {
	orm.RegisterModelWithPrefix("t_", new(User))
}