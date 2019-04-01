package dto

import (
	"github.com/astaxie/beego/validation"
)
var USER_SEARCH = map[string]interface{}{
	"d"		: "department_id",
	"n"  	: "username",
}
type PwdResetDto struct {
	//PrevPwd string `form:"password" valid:"Required"`
	NewPwd string `form:"new_password" valid:"Required"`
	NewRePwd string `form:"re_password" valid:"Required"`
}
type PwdUserResetDto struct {
	//PrevPwd string `form:"password" valid:"Required"`
	Uid int `form:"user_id" valid:"Required"`
	NewPwd string `form:"new_password" valid:"Required"`
	NewRePwd string `form:"re_password" valid:"Required"`
}
func (prd *PwdResetDto) Valid(v *validation.Validation){
	//if prd.PrevPwd == prd.NewRePwd {
	//	v.SetError("PrevPwd","新旧密码不可相同")
	//}
	if prd.NewPwd != prd.NewRePwd{
		v.SetError("NewPwd","新密码与确认密码不一致")
	}
}

type MoveDepartmentDto struct {
	Uids string `form:"uids" valid:"Required"`
	Did int `form:"department_id" valid:"Required"`
}
