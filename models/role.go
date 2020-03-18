package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Role struct {
	Id int
	RoleName string
	Status int
	CreateUser string
	CreateTime time.Time
	UpdateTime time.Time
}

type RoleUser struct {
	Id int
	RoleId int
	UserId int
}

type RoleAuth struct {
	Id int
	RoleId int
	AuthId int
}

func init()  {
	orm.RegisterModelWithPrefix("t_", new(Role), new(RoleUser), new(RoleAuth))
}