package controllers

import (
	"github.com/bullteam/zeus/components"
	"github.com/bullteam/zeus/models"
	"github.com/bullteam/zeus/service"
	"github.com/bullteam/zeus/utils"
	"github.com/astaxie/beego"
	"github.com/dchest/captcha"
	"strconv"
	"strings"
	"github.com/bullteam/zeus/dto"
	"crypto/md5"
	"fmt"
)

type UserController struct {
	TokenCheckController
} 
type AccountController struct {
	BaseController
}
type LoginoutController struct {
	TokenCheckController
}
type FindpasswdController struct {
	BaseController
}
type ChangepasswdController struct {
	TokenCheckController
}
type ChangeuserroleController struct {
	BaseController
}

func (c *AccountController) Login() {
	form := models.LoginForm{}
	if err := c.ParseForm(&form); err != nil || form.Username == "" || form.Password == "" {
		beego.Debug("ParseLoginForm:", err)
		c.Fail(components.ErrInputData)
		return
	}
    displayCapcha := models.DisplayCapcha(form.Username)//是否校验验证码
    //beego.Debug(displayCapcha)
	if displayCapcha && (form.CaptchaId == "" || form.CaptchaVal == "") {
		beego.Debug("captcha:",form.CaptchaId+"-"+form.CaptchaVal)
		c.Fail(components.ErrCaptchaEmpty)
		return
	}
	if displayCapcha && (!captcha.VerifyString(form.CaptchaId, form.CaptchaVal)) {
		beego.Debug("captcha:",form.CaptchaId+"-"+form.CaptchaVal)
		c.Fail(components.ErrCaptcha)
		return
	}

	//beego.Debug("ParseLoginForm:", &form)
	user := models.User{}
	if _, err := user.FindByID(form.Username); err != nil {
		beego.Error("FindUserById:", err)
		c.Fail(components.ErrNoUser)
		return
	}
	if ok, err := user.CheckPass(form.Password); err != nil || !ok {
		models.SetCapcha(form.Username)//错误三次设置显示验证码
		beego.Error("CheckUserPass:", err)
		c.Fail(components.ErrPass)
		return
	} else if !ok {
		c.Fail(components.ErrSystem)
		return
	}
	//generate jwt with rsa private key
	jwtoken,err := utils.GenerateJwtWithUserInfo(strconv.Itoa(user.Id),user.Username)
	if err != nil{
		controllErr := components.ErrGenJwt
		controllErr.Moreinfo = err.Error()
		c.Fail(controllErr)
		return
	}else {
		md5Ctx := md5.New()
		md5Ctx.Write([]byte(jwtoken))
		cipherStr := md5Ctx.Sum(nil)
		refreshToken,_ := utils.GenerateRefreshJwtWithToken(fmt.Sprintf("auth%xsafe",cipherStr))
		c.Resp(0,"success",map[string]interface{}{
			"access_token":jwtoken,
			"refresh_token":refreshToken,
            "userid": user.Id,
			"username" : user.Username,
		})
	}
}

func (c *UserController ) Add()  {
	username := c.GetString("username")
	password := c.GetString("password")
	mobile := c.GetString("mobile")
	faceicon := c.GetString("faceicon")
	realname := c.GetString("realname")
	email := c.GetString("email")
	roles := c.GetString("roles")
	title := c.GetString("title")
	if username == "" || password == "" {
		c.Fail(components.ErrInputData)
		return
	}
	sex,err := c.GetInt("sex")
	if err != nil{
		c.Fail(components.ErrInputData)
		return
	}
	status,err := c.GetInt("status")
	if err != nil{
		status = 1
	}
	dept_id,err := c.GetInt("dept_id")
	if err != nil{
		c.Fail(components.ErrInputData)
		return
	}
	userCheck := models.User{}
	if _, err := userCheck.FindByID(username); err == nil {
		c.Fail(components.ErrDupRecord,"用户已存在")
		return
	}
	id,err := models.NewUser(username,password,mobile,sex,realname,email,status,faceicon,dept_id,title)
	//用户，角色关联
	if roles != "" {
		sroles := strings.Split(roles, ",")
		us := service.UserService{}
		us.AddRoles(id, sroles)
	}
	c.Resp(0,"success",map[string]interface{}{
		"id":id,
	})
}

func (c *UserController) Logout() {
	c.Fail(components.Err404)
}

func (c *FindpasswdController) Post() {
	c.Fail(components.Err404)
}

func (c *UserController) Edit() {
	username := c.GetString("username")
	password := c.GetString("password")
	mobile := c.GetString("mobile")
	faceicon := c.GetString("faceicon")
	realname := c.GetString("realname")
	email := c.GetString("email")
	roles := c.GetString("roles")
	title := c.GetString("title")
	if username == "" {
		c.Fail(components.ErrInputData)
		return
	}
	id,err := c.GetInt("id")
	if err != nil{
		c.Fail(components.ErrInputData)
		return
	}
	sex,err := c.GetInt("sex")
	if err != nil{
		c.Fail(components.ErrInputData)
		return
	}
	status,err := c.GetInt("status")
	if err != nil{
		c.Fail(components.ErrInputData)
		return
	}
	dept_id,err := c.GetInt("dept_id")
	if err != nil{
		c.Fail(components.ErrInputData)
		return
	}
	if _, err := models.GetUserByUid(int64(id)); err != nil {
		c.Fail(components.ErrNoUser)
		return
	}
	err = models.UpdateUser(id, username, password, mobile, sex, realname, email, status, faceicon, dept_id,title)
	//用户，角色关联
	//if roles != "" {
		sroles := strings.Split(roles, ",")
		us := service.UserService{}
		us.AddRoles(int64(id), sroles)
		if err != nil {
			c.Fail(components.ErrNoUserChange)
			return
		}
	//}
	c.Resp(0,"success",map[string]interface{}{
	})
}

func (c *UserController)  Updatestatus(){
	id,err := c.GetInt("id")
	if err != nil{
		c.Fail(components.ErrInputData)
		return
	}
	status,err := c.GetInt("status")
	if err != nil{
		c.Fail(components.ErrInputData)
		return
	}
	err = models.UpdateStatus(id, status)
	if err != nil {
		c.Fail(components.ErrNoUserChange)
		return
	}
	c.Resp(0,"success",map[string]interface{}{
	})
}

//func (c *UserController) List()  {
//	page, page_err := c.GetInt("p")
//	if page_err != nil {
//		page = 1
//	}
//	offset := 20
//    user,cnt := models.User_list(page,offset)
//	c.Resp(0,"success",map[string]interface{}{
//		"user" : user,
//		"total" : cnt,
//		"page" : page,
//	})
//}
//用户列表，支持查询
func (c *UserController) List() {
	us := service.UserService{}
	start, _ := c.GetInt("start", 0)
	limit, _ := c.GetInt("limit", LIST_ROWS_PERPAGE)
	q := c.GetString("q")
	data, total := us.GetList(start, limit,strings.Split(q,","))
	c.Resp(0, "success", map[string]interface{}{
		"result": data,
		"total":  total,
	})
}
func (c *UserController) Show()  {
	id,err := c.GetInt("id")
	if err != nil {
		c.Fail(components.ErrIdData)
		return
	}
	user,err := models.GetUserByUid(int64(id))
	if err != nil {
		c.Fail(components.ErrInputData)
		return
	}
	role := models.Role{}
	roles := role.GetRolesByUid(id)
	c.Resp(0,"success",map[string]interface{}{
		"userinfo":user,
		"role":roles,
	})
}

/**删除用户**/
func (c *UserController) Del() {
	id,err := c.GetInt("id")
	if err != nil {
		c.Fail(components.ErrInputData)
		return
	}
	err = models.DeleteUser(id)
	if err != nil {
		c.Fail(components.ErrInputData)
		return
	}
	c.Resp(0, "success", map[string]interface{}{})
}
func (c *UserController) GetMenu(){
	userService := service.UserService{}
	code := c.GetString("code")
	menus := userService.GetMenusByDomain(c.Uid,code)
	c.Resp(0,code,menus)
}

func (c *UserController) GetDomain(){
	userService := service.UserService{}
	domains := userService.GetRelatedDomains(c.Uid)
	c.Resp(0,"success",domains)
}

func (c *UserController) ChangePwd(){
	userService := service.UserService{}
	pwdDto := &dto.PwdResetDto{}
	c.ParseAndValidate(pwdDto)
	uid,_ := strconv.Atoi(c.Uid)
	if err := userService.ResetPassword(uid,pwdDto);err != nil{
		c.Fail(components.ErrEditFail,err.Error())
	}
	c.Resp(0,"success")
}
func (c *UserController) ChangeUserPwd(){
	userService := service.UserService{}
	pwdDto := &dto.PwdUserResetDto{}
	c.ParseAndValidate(pwdDto)

	if err := userService.ResetUserPassword(pwdDto.Uid,pwdDto);err != nil{
		c.Fail(components.ErrEditFail,err.Error())
	}
	c.Resp(0,"success")
}

func (c *UserController) MoveToNewDepartment(){
	userService := service.UserService{}
	deptDto := &dto.MoveDepartmentDto{}
	c.ParseAndValidate(deptDto)
	if _,err := userService.SwitchDepartment(strings.Split(deptDto.Uids,","),deptDto.Did);err != nil{
		c.Fail(components.ErrEditFail,err.Error())
	}
	c.Resp(0,"success")
}

func (c *UserController) RefreshToken(){
	rt := c.GetString("refresh_token")
	if rt == ""{
		c.Fail(components.ErrWrongRefreshToken)
		return
	}
	jwtHandler := components.NewJwtHandler()
	if claims,err := jwtHandler.ValidateRefreshToken(rt);err == nil{
		vmd5 := md5.New()
		vmd5.Write([]byte(c.RawToken))
		cipherStr := vmd5.Sum(nil)
		if claims.Token != fmt.Sprintf("auth%xsafe",cipherStr){
			c.Fail(components.ErrWrongRefreshToken)
			return
		}
		jwtoken,err := utils.GenerateJwtWithUserInfo(c.Uid,c.Uname)
		if err != nil{
			controllErr := components.ErrGenJwt
			controllErr.Moreinfo = err.Error()
			c.Fail(controllErr)
			return
		}else {
			md5Ctx := md5.New()
			md5Ctx.Write([]byte(jwtoken))
			cipherStr := md5Ctx.Sum(nil)
			refreshToken,_ := utils.GenerateRefreshJwtWithToken(fmt.Sprintf("auth%xsafe",cipherStr))
			c.Resp(0,"success",map[string]interface{}{
				"access_token":jwtoken,
				"refresh_token":refreshToken,
				"userid": c.Uid,
				"username" : c.Uname,
			})
		}
	}
	c.Fail(components.ErrWrongRefreshToken)
	return
}