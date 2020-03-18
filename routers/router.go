package routers

import (
	"beegodemo/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/logout", &controllers.LoginController{},"get:Logout")
	beego.Router("/index", &controllers.IndexController{})
	beego.Router("/index/:authcode", &controllers.IndexController{})
/*	beego.Router("/role/add", &controllers.RegisterController{},"get:AddRole")
	beego.Router("/auth/add", &controllers.RegisterController{},"get:AddAuth")
	beego.Router("/user/update", &controllers.RegisterController{},"get:Update")*/
}
