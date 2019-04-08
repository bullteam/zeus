package tests

import (
	"testing"
	"github.com/bullteam/zeus/pkg/service"
	"github.com/bullteam/zeus/pkg/dto"
    "github.com/astaxie/beego/validation"
)
var userService = service.UserService{}
var deptService = service.DepartmentService{}
func TestUserService_GetMenusByDomain(t *testing.T){
	menus := userService.GetMenusByDomain("6","admin-finance")
	if len(menus) > 0{
		t.Log(menus)
		t.Log("Menu receiving successfully")
	}
}
func TestUserService_ResetPassword(t *testing.T){
	dto := &dto.PwdResetDto{
		"123456",
		"123456",
	}
	v := &validation.Validation{}
	dto.Valid(v)

	if len(v.Errors) > 0{
		es := ""
		for _,e := range v.Errors{
			es += " // "+e.Error()
		}
		t.Error(es)
		return
	}
	if err := userService.ResetPassword(6,dto);err != nil{
		t.Error(err.Error())
	}else{
		t.Log("update success")
	}
}

func TestUserService_SwitchDepartment(t *testing.T){
	if _,err := userService.SwitchDepartment([]string{"6","9"},1);err != nil{
		t.Error(err.Error())
	}
	t.Log("Users were moved to antoher department")
}