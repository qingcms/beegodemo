package main

import (
	_ "beegodemo/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"

)

func init(){
	orm.RegisterDriver("mysql",orm.DRMySQL)

	orm.RegisterDataBase("default","mysql","root:123456@/beegodemo?charset=utf8")
}


func main() {

	o := orm.NewOrm()
	o.Using("default")

	beego.Run()
}

