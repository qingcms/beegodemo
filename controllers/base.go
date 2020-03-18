package controllers

import (
	"beegodemo/models"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"strings"
)

const  AdminUserStateKey  = "admin_userstate_key"
type BaseController struct {
	beego.Controller
}



func (this *BaseController) getUserState() models.User {
	//获取当前session信息
	  id :=this.GetSession(AdminUserStateKey+"_id")
	  name :=this.GetSession(AdminUserStateKey+"_name")
	  username :=this.GetSession(AdminUserStateKey+"_username")
	  this.CheckUserAuth()
	  //当前session
    if id==nil || name==nil || username==nil {
    	this.Ctx.Redirect(302,"/login")
	}

	userinfo := models.User{UserName: username.(string), Id: id.(int), Name: name.(string)}
	return userinfo
}

func (this *BaseController) CheckUserAuth(){

	//设置公共不拦截页面
	publicUrlData :=[]string{
		"/index",
		"/login",
	}

	//读取当前页面url
	selfURI :=this.Ctx.Request.RequestURI
    //获取用户权限列表
	userauthlist :=this.GetAuthList()

	//循环当前用户权限，判断用户是否具有权限
	var isAuth bool
	isAuth = false

	for _,rowPublicAuth :=range publicUrlData{
		if strings.Contains("*"+selfURI,"*"+fmt.Sprint(rowPublicAuth)) {
			isAuth = true
			break
		}
	}

   if isAuth==false {
		for _ , rowUserAuth:=range userauthlist{
			//判断为空本次不执行，循环下一次
			if rowUserAuth[5] ==nil{
				continue
			}
			//判断系统字符串是否包含当前url
			if strings.Contains("*"+selfURI,"*"+fmt.Sprint(rowUserAuth[5])) {
				isAuth = true
				break
			}
		}
   }
	//没有权限
	if isAuth==false{
		this.Ctx.Redirect(504,"/login")
	}
	//递归处理用户权限列表


}

/*
 *获取用户权限列表
 */
func (this *BaseController) GetAuthList() [] orm.ParamsList {

	userid :=this.GetSession(AdminUserStateKey+"_id")

	//根据userid查询角色Id
	o :=orm.NewOrm()

	var rolelist [] orm.ParamsList
    var roleauthlist [] orm.ParamsList
	var userauthlist [] orm.ParamsList
	//查询数据库
	rsuser :=o.QueryTable(new(models.RoleUser))
	num,err :=rsuser.Filter("userId",userid).ValuesList(&rolelist,"RoleId")

	if err!=nil {
		this.Ctx.Redirect(504,"/login")
	}
	//循环角色
	//获取角色Id
	var roleids [] string
	if num>0{
		for _,rowRoleUser:=range rolelist{
			roleids=append(roleids,fmt.Sprint(rowRoleUser[0]))
		}
	}
	//根据角色获取权限ID
	rsroleauth :=o.QueryTable(new(models.RoleAuth))
	numroleauth,err :=rsroleauth.Filter("RoleId__in",roleids).Distinct().OrderBy("AuthId").ValuesList(&roleauthlist,"AuthId")

	if err!=nil {
		this.Ctx.Redirect(504,"/login")
	}
	var authids [] string
	if numroleauth>0{
		for _,rowRoleAuth:=range roleauthlist{
			authids=append(authids,fmt.Sprint(rowRoleAuth[0]))
		}
	}
	//fmt.Println(authids)
    //获取有效的用户权限列表
	rsuserauth :=o.QueryTable(new(models.Auth))

	numuserauth,err :=rsuserauth.Filter("Id__in",authids).OrderBy("Id").Limit(50).ValuesList(&userauthlist,"Id","AuthCode","AuthName","Status","ParentId","AuthUrl","AuthPath","ParentPath")
	//fmt.Println(err)
	if err!=nil {
		this.Ctx.Redirect(504,"/login")
	}
	//获取权限顶级目录列表
/*	var usertopauthlist [] orm.ParamsList
	numtopauth,err :=rsuserauth.Filter("Id__in",authids).Filter("ParentId",1).OrderBy("Id").Limit(10).ValuesList(&usertopauthlist,"Id","AuthCode","AuthName","Status","ParentId","AuthUrl","AuthPath","ParentPath")
*/	var usertopauthlist [] orm.Params
	numtopauth,err :=rsuserauth.Filter("Id__in",authids).Filter("ParentId",1).OrderBy("Id").Limit(10).Values(&usertopauthlist,"Id","AuthCode","AuthName","Status","ParentId","AuthUrl","AuthPath","ParentPath")

	if numtopauth>0{
		this.Data["TopMenu"]=usertopauthlist
	}
	//读取顶级菜单code
	pauthcode :=this.GetString(":authcode")
	//判断为空时 设置默认读取系统管理
    if pauthcode == "" {
		pauthcode="sys"
	}
	//获取左边菜单权限
	var leftUserAuthList [] orm.Params
	numLeftauth,err :=rsuserauth.Filter("Id__in",authids).Filter("ParentPath__istartswith","root."+pauthcode).OrderBy("Id").Limit(10).Values(&leftUserAuthList,"Id","AuthCode","AuthName","Status","ParentId","AuthUrl","AuthPath","ParentPath")

	fmt.Println(leftUserAuthList)

	var treeAuthList [] orm.Params
    var parentPath string
	   parentPath ="root."+pauthcode
	   //左边菜单 第一层
	for _,rowUserAuth :=range leftUserAuthList{

		parentPath2 :=rowUserAuth["ParentPath"]
		if parentPath2==parentPath{
			rowUserAuth["SubList"]=nil
			treeAuthList=append(treeAuthList,rowUserAuth)
		}
	}
    //左边菜单第二层
	var treeSubAuthList [] orm.Params
	for lineone,rowOneUserAuth :=range treeAuthList{
		authcode :=fmt.Sprint(rowOneUserAuth["AuthCode"])
		parentPath ="root."+pauthcode+"."+authcode
		treeSubAuthList =nil
		for _,rowTwoUserAuth :=range leftUserAuthList {
			parentPath2 :=rowTwoUserAuth["ParentPath"]
			if parentPath2 == parentPath {
				treeSubAuthList = append(treeSubAuthList, rowTwoUserAuth)
			}
		}
		fmt.Println(treeSubAuthList)
		treeAuthList[lineone]["SubList"]=treeSubAuthList
	}
  fmt.Println(treeAuthList)

	if numLeftauth>0{
		this.Data["LeftMenu"]=treeAuthList
	}

    if numuserauth >0 {
		return userauthlist
	}
	return nil
}

/**
 * 处理权限树
 */
/*func (this *BaseController) GetTreeAuthList(pauthcode string,userAuthList *orm.Params) map[string] orm.Params{
    var treeAuthList orm.Params

	for _,rowUserAuth :=range userAuthList{
          fmt.Println(rowUserAuth)
	   /* if fmt.Sprint(rowUserAuth["ParentPath"]) == fmt.Sprint("root.",pauthcode) {
			treeAuthList=append(treeAuthList,rowUserAuth)
		}*/
/*	}
return treeAuthList
}*/




