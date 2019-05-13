package controllers

import (
	"crypto/md5"
	"fmt"
	"github.com/astaxie/beego"
	"zeus/pkg/components"
	"zeus/pkg/dto"
	"zeus/pkg/service"
	"zeus/pkg/utils"
	"github.com/dchest/captcha"
	"strconv"
	"strings"
)

type UserController struct {
	TokenCheckController
}

type AccountController struct {
	BaseController
}

func (c *AccountController) Login() {
	c.setLangVer() //设置语言
	loginDto := &dto.LoginDto{}
	err := c.ParseAndValidateFirstErr(loginDto)
	if err != nil {
		c.Fail(components.ErrInvalidParams, err.Error())
		return
	}
	userService := service.UserService{}
	displayCapcha := userService.DisplayCapcha(loginDto.Username) //是否校验验证码
	if displayCapcha && (loginDto.CaptchaId == "" || loginDto.CaptchaVal == "") {
		beego.Debug("captcha:", loginDto.CaptchaId+"-"+loginDto.CaptchaVal)
		c.Fail(components.ErrCaptchaEmpty)
		return
	}
	if displayCapcha && (!captcha.VerifyString(loginDto.CaptchaId, loginDto.CaptchaVal)) {
		beego.Debug("captcha:", loginDto.CaptchaId+"-"+loginDto.CaptchaVal)
		c.Fail(components.ErrCaptcha)
		return
	}

	user, err := userService.FindByUserName(loginDto.Username)
	if err != nil {
		beego.Error("FindByUserName:", err)
		c.Fail(components.ErrNoUser)
		return
	}
	if ok, err := userService.CheckPass(loginDto.Password, user); err != nil || !ok {
		userService.SetCapcha(loginDto.Username) //错误三次设置显示验证码
		beego.Error("CheckUserPass:", err)
		c.Fail(components.ErrPass)
		return
	} else if !ok {
		c.Fail(components.ErrSystem)
		return
	}
	//generate jwt with rsa private key
	jwtoken, err := utils.GenerateJwtWithUserInfo(strconv.Itoa(user.Id), user.Username)
	if err != nil {
		controllErr := components.ErrGenJwt
		controllErr.Moreinfo = err.Error()
		c.Fail(controllErr)
		return
	} else {
		md5Ctx := md5.New()
		md5Ctx.Write([]byte(jwtoken))
		cipherStr := md5Ctx.Sum(nil)
		refreshToken, _ := utils.GenerateRefreshJwtWithToken(fmt.Sprintf("auth%xsafe", cipherStr))
		c.Resp(0, "success", map[string]interface{}{
			"access_token":  jwtoken,
			"refresh_token": refreshToken,
			"userid":        user.Id,
			"username":      user.Username,
		})
	}
}

//钉钉登陆
func (c *AccountController) DingtalkLogin() {
	c.setLangVer() //设置语言
	dingtalkDto := &dto.LoginDingtalkDto{}
	err := c.ParseAndValidateFirstErr(dingtalkDto)
	if err != nil {
		c.Fail(components.ErrInvalidParams, err.Error())
		return
	}
	userService := service.UserService{}
	user,err := userService.LoginByDingtalk(dingtalkDto.Code)
	if err != nil{
		c.Fail(components.ErrNoUser, err.Error())
		return
	}
	//generate jwt with rsa private key
	jwtoken, err := utils.GenerateJwtWithUserInfo(strconv.Itoa(user.Id), user.Name)
	if err != nil {
		controllErr := components.ErrGenJwt
		controllErr.Moreinfo = err.Error()
		c.Fail(controllErr)
		return
	} else {
		md5Ctx := md5.New()
		md5Ctx.Write([]byte(jwtoken))
		cipherStr := md5Ctx.Sum(nil)
		refreshToken, _ := utils.GenerateRefreshJwtWithToken(fmt.Sprintf("auth%xsafe", cipherStr))
		c.Resp(0, "success", map[string]interface{}{
			"access_token":  jwtoken,
			"refresh_token": refreshToken,
			"userid":        user.Id,
			"username":      user.Name,
		})
	}
}

func (c *UserController) Add() {
	userAddDto := &dto.UserAddDto{}
	err := c.ParseAndValidateFirstErr(userAddDto)
	if err != nil {
		c.Fail(components.ErrInvalidParams, err.Error())
		return
	}
	userService := service.UserService{}
	if _, err := userService.FindByUserName(userAddDto.Username); err == nil {
		c.Fail(components.ErrDupRecord, "用户已存在")
		return
	}
	id, err := userService.NewUser(userAddDto)
	//用户，角色关联
	roles := c.GetString("roles")
	if roles != "" {
		sroles := strings.Split(roles, ",")
		userService.AddRoles(id, sroles)
	}
	c.Resp(0, "success", map[string]interface{}{
		"id": id,
	})
}

func (c *UserController) Logout() {
	c.Fail(components.Err404)
}

func (c *AccountController) Post() {
	c.Fail(components.Err404)
}

func (c *UserController) Edit() {
	userEditDto := &dto.UserEditDto{}
	err := c.ParseAndValidateFirstErr(userEditDto)
	if err != nil {
		c.Fail(components.ErrInvalidParams, err.Error())
		return
	}
	userService := service.UserService{}
	if _, err := userService.GetUserByUid(int64(userEditDto.Id)); err != nil {
		c.Fail(components.ErrNoUser)
		return
	}
	err = userService.UpdateUser(userEditDto)
	if err != nil {
		c.Fail(components.ErrEditFail, err.Error())
		return
	}
	//用户，角色关联
	roles := c.GetString("roles")
	//if roles != "" {
	sroles := strings.Split(roles, ",")
	userService.AddRoles(int64(userEditDto.Id), sroles)
	if err != nil {
		c.Fail(components.ErrNoUserChange)
		return
	}
	//}
	c.Resp(0, "success", map[string]interface{}{})
}

func (c *UserController) UpdateStatus() {
	id, err := c.GetInt("id")
	if err != nil {
		c.Fail(components.ErrInputData)
		return
	}
	status, err := c.GetInt("status")
	if err != nil {
		c.Fail(components.ErrInputData)
		return
	}
	userService := service.UserService{}
	err = userService.UpdateStatus(id, status)
	if err != nil {
		c.Fail(components.ErrNoUserChange)
		return
	}
	c.Resp(0, "success", map[string]interface{}{})
}

//用户列表，支持查询
func (c *UserController) List() {
	us := service.UserService{}
	start, _ := c.GetInt("start", 0)
	limit, _ := c.GetInt("limit", listRowsPerPage)
	q := c.GetString("q")
	data, total := us.GetList(start, limit, strings.Split(q, ","))
	c.Resp(0, "success", map[string]interface{}{
		"result": data,
		"total":  total,
	})
}

func (c *UserController) Show() {
	id, err := c.GetInt("id")
	if err != nil {
		c.Fail(components.ErrIdData)
		return
	}
	user_id, err := strconv.Atoi(c.Uid)
	if err != nil {
		c.Fail(components.ErrInvalidUser, err.Error())
		return
	}
	userService := service.UserService{}
	user, err := userService.GetUserByUid(int64(id))
	if err != nil {
		c.Fail(components.ErrInputData)
		return
	}
	roleService := service.RoleService{}
	roles := roleService.GetRolesByUid(id)
	OauthUserInfoService := service.UserService{}
	OauthUserInfo,_ := OauthUserInfoService.GetBindOauthUserInfo(user_id)
	c.Resp(0, "success", map[string]interface{}{
		"userinfo": user,
		"role":     roles,
		"oauth_user_info" : OauthUserInfo,
	})
}

/**删除用户**/
func (c *UserController) Del() {
	id, err := c.GetInt("id")
	if err != nil {
		c.Fail(components.ErrInputData)
		return
	}
	userService := service.UserService{}
	err = userService.DeleteUser(id)
	if err != nil {
		c.Fail(components.ErrInputData)
		return
	}
	c.Resp(0, "success", map[string]interface{}{})
}

func (c *UserController) GetMenu() {
	userService := service.UserService{}
	code := c.GetString("code")
	menus := userService.GetMenusByDomain(c.Uid, code)
	c.Resp(0, code, menus)
}

func (c *UserController) GetDomain() {
	userService := service.UserService{}
	domains := userService.GetRelatedDomains(c.Uid)
	c.Resp(0, "success", domains)
}

func (c *UserController) ChangePwd() {
	userService := service.UserService{}
	pwdDto := &dto.PwdResetDto{}
	err := c.ParseAndValidateFirstErr(pwdDto)
	if err != nil {
		c.Fail(components.ErrInvalidParams, err.Error())
		return
	}
	uid, _ := strconv.Atoi(c.Uid)
	if err = userService.ResetPassword(uid, pwdDto); err != nil {
		c.Fail(components.ErrEditFail, err.Error())
		return
	}
	c.Resp(0, "success")
}

func (c *UserController) ChangeUserPwd() {
	userService := service.UserService{}
	pwdDto := &dto.PwdUserResetDto{}
	err := c.ParseAndValidateFirstErr(pwdDto)
	if err != nil {
		c.Fail(components.ErrInvalidParams, err.Error())
		return
	}

	if err = userService.ResetUserPassword(pwdDto.Uid, pwdDto); err != nil {
		c.Fail(components.ErrEditFail, err.Error())
		return
	}
	c.Resp(0, "success")
}

func (c *UserController) MoveToNewDepartment() {
	userService := service.UserService{}
	deptDto := &dto.MoveDepartmentDto{}
	err := c.ParseAndValidateFirstErr(deptDto)
	if err != nil {
		c.Fail(components.ErrInvalidParams, err.Error())
		return
	}
	if _, err := userService.SwitchDepartment(strings.Split(deptDto.Uids, ","), deptDto.Did); err != nil {
		c.Fail(components.ErrEditFail, err.Error())
		return
	}
	c.Resp(0, "success")
}

func (c *UserController) RefreshToken() {
	rt := c.GetString("refresh_token")
	if rt == "" {
		c.Fail(components.ErrWrongRefreshToken)
		return
	}
	jwtHandler := components.NewJwtHandler()
	if claims, err := jwtHandler.ValidateRefreshToken(rt); err == nil {
		vmd5 := md5.New()
		vmd5.Write([]byte(c.RawToken))
		cipherStr := vmd5.Sum(nil)
		if claims.Token != fmt.Sprintf("auth%xsafe", cipherStr) {
			c.Fail(components.ErrWrongRefreshToken)
			return
		}
		jwtoken, err := utils.GenerateJwtWithUserInfo(c.Uid, c.Uname)
		if err != nil {
			controllErr := components.ErrGenJwt
			controllErr.Moreinfo = err.Error()
			c.Fail(controllErr)
			return
		} else {
			md5Ctx := md5.New()
			md5Ctx.Write([]byte(jwtoken))
			cipherStr := md5Ctx.Sum(nil)
			refreshToken, _ := utils.GenerateRefreshJwtWithToken(fmt.Sprintf("auth%xsafe", cipherStr))
			c.Resp(0, "success", map[string]interface{}{
				"access_token":  jwtoken,
				"refresh_token": refreshToken,
				"userid":        c.Uid,
				"username":      c.Uname,
			})
		}
	}
	c.Fail(components.ErrWrongRefreshToken)
	return
}
