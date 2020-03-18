package main

import (
	"beegodemo/models"
	_ "beegodemo/routers"
	"encoding/gob"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init(){
	//配置数据库
	orm.RegisterDriver("mysql",orm.DRMySQL)

	orm.RegisterDataBase("default","mysql","root:123456@/beegodemo?charset=utf8")

   //注册session 结构体
   gob.Register(models.User{})
}


func main() {

	o := orm.NewOrm()
	o.Using("default")

	beego.Run()
}

