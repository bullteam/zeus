package routers

import (
	"zeus/pkg/controllers"
	"github.com/astaxie/beego"
	"github.com/dchest/captcha"
	//"github.com/astaxie/beego/plugins/cors"
)

func init() {
	//beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
	//	AllowAllOrigins:  true,
	//	AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	//	AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
	//	ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
	//	AllowCredentials: true,
	//}))

	/*******用户管理*******/
	accountController := &controllers.AccountController{}
	userController := &controllers.UserController{}
	beego.Router("/user/login",accountController,"post:Login")           //用户登录
	beego.Router("/user/refresh-token",userController,"post:RefreshToken") //刷新令牌
	beego.Router("/user/loginout", userController, "get:Logout")            //用户退出登录
	beego.Router("/user/findpasswd", &controllers.FindpasswdController{})   //找回密码
	//beego.Router("/user/changepwd", userController, "post:ChangePwd")            //用户更改密码
	beego.Router("/user/change-password", userController, "post:ChangePwd")
	beego.Router("/user/change-user-password", userController, "post:ChangeUserPwd")
	beego.Router("/user/add", userController, "post:Add")        //用户注册
	//beego.Router("/user", userController, "get:List")            //用户列表
	beego.Router("/user", userController, "get:List")            //用户列表
	beego.Router("/user/edit", userController, "post:Edit")      //用户编辑
	beego.Router("/user/show", userController, "get:Show")       //用户信息
	beego.Router("/user/del", userController, "post:Del")        //删除用户信息
	beego.Router("/user/updatestatus", userController, "post:Updatestatus") //删除用户信息
	beego.Router("/user/menu",userController,"get:GetMenu") // 获取用户相关菜单
	beego.Router("/user/domain",userController,"get:GetDomain") // 获取用户相关域
	beego.Router("/user/move-department",userController,"post:MoveToNewDepartment")

	/******角色管理*****/
	roleController := &controllers.RoleController{}
	beego.Router("/role/list", roleController, "get:List")  //角色列表
	beego.Router("/role/show", roleController, "get:Show")  //角色详情
	beego.Router("/role/add", roleController, "post:Add")   //添加角色
	beego.Router("/role/edit", roleController, "post:Edit") //更新角色信息
	beego.Router("/role/del", roleController, "post:Del")   //删除角色信息
	beego.Router("/role/assign", roleController, "post:Assign")   //删除角色信息


	/******部门管理*****/
	deptController := &controllers.DeptController{}
	beego.Router("/dept/list", deptController, "get:List")  //部门列表
	beego.Router("/dept/show", deptController, "get:Show")  //部门详情
	beego.Router("/dept/add", deptController, "post:Post")   //添加部门
	beego.Router("/dept/edit", deptController, "post:Edit") //更新部门信息
	beego.Router("/dept/del", deptController, "post:Del")   //删除部门信息
	beego.Router("/dept/check-no-member", deptController, "post:CheckNoMember")   //删除菜单

	/******菜单管理****/
	menuController := &controllers.MenuController{}
	beego.Router("/menu", menuController, "get:List")       //菜单列表
	beego.Router("/menu/add", menuController, "post:Add")   //添加菜单
	beego.Router("/menu/edit", menuController, "post:Edit") //更新菜单
	beego.Router("/menu/del", menuController, "post:Del")   //删除菜单


	/******域管理****/
	DomainController := &controllers.DomainController{}
	beego.Router("/domain/list", DomainController, "get:List")  //域列表
	beego.Router("/domain/show", DomainController, "get:Show")  //域的信息
	beego.Router("/domain/add", DomainController, "post:Post")  //添加域
	beego.Router("/domain/edit", DomainController, "post:Edit") //修改域
	beego.Router("/domain/del", DomainController, "post:Del")   //删除域

	/******权限管理*****/
	PermController := &controllers.PermController{}
	beego.Router("/user/perm/list", PermController, "get:GetPermsByLoginUser")
	beego.Router("/user/perm/check",PermController,"post:CheckPerm")

	// 验证码服务
	beego.Router("/captcha/request", &controllers.CaptchaController{})
	beego.Handler("/captcha/*.png", captcha.Server(240, 80)) //验证图片的宽和高(px)
}