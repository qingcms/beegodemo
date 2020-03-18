package controllers

type IndexController struct {
	BaseController
}



func (this *IndexController) Get () {

    userinfo :=this.getUserState()
    this.Data["userinfo"]=userinfo
	this.TplName="index.tpl"

}

