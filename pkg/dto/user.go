package dto

import "github.com/astaxie/beego/validation"

var UserSearch = map[string]interface{}{
	"d": "department_id",
	"n": "username",
}

type UserAddDto struct {
	Username     string `form:"username" valid:"Required"`
	Mobile       string `form:"mobile"`
	Email        string `form:"email"`
	Password     string `form:"password" valid:"Required"`
	Faceicon     string `form:"faceicon"`
	Realname     string `form:"realname"`
	Title        string `form:"title"`
	Sex          int    `form:"sex"`
	Status       int    `form:"status"`
	DepartmentId int    `form:"dept_id"`
}

type UserEditDto struct {
	Id           int    `form:"id" valid:"Required"`
	Username     string `form:"username" valid:"Required"`
	Mobile       string `form:"mobile"`
	Email        string `form:"email"`
	Password     string `form:"password"`
	Faceicon     string `form:"faceicon"`
	Realname     string `form:"realname"`
	Title        string `form:"title"`
	Sex          int    `form:"sex"`
	Status       int    `form:"status"`
	DepartmentId int    `form:"dept_id"`
}

type LoginDto struct {
	Username   string `form:"username"`
	Password   string `form:"password" valid:"Required"`
	CaptchaId  string `form:"captchaid"`
	CaptchaVal string `form:"captchaval"`
}

type PwdResetDto struct {
	NewPwd   string `form:"new_password" valid:"Required"`
	NewRePwd string `form:"re_password" valid:"Required"`
}

type ChangeUserRoleForm struct {
	Id       string `form:"id"`
	Username string `form:"username"`
}

type PwdUserResetDto struct {
	Uid      int    `form:"user_id" valid:"Required"`
	NewPwd   string `form:"new_password" valid:"Required"`
	NewRePwd string `form:"re_password" valid:"Required"`
}

type MoveDepartmentDto struct {
	Uids string `form:"uids" valid:"Required"`
	Did  int    `form:"department_id" valid:"Required"`
}

type LoginDingtalkDto struct {
	Code string `form:"code"`
}

type BindThirdDto struct {
	From int    `form:"from"`
	Code string `form:"code"`
}

type UnBindThirdDto struct {
	From int `form:"from"`
}

func (prd *PwdResetDto) Valid(v *validation.Validation) {
	if prd.NewPwd != prd.NewRePwd {
		v.SetError("NewPwd", "新密码与确认密码不一致")
	}
}
