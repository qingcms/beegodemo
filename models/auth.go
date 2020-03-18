package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Auth struct {
	Id int
	AuthCode string
	AuthName string
	Status int
	ParentId int
	AuthPath string
	ParentPath string
	CreateUser string
	CreateTime time.Time
	UpdateTime time.Time
}



func init()  {
	orm.RegisterModelWithPrefix("t_", new(Auth))
}