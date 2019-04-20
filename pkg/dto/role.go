package dto

import "github.com/astaxie/beego/validation"

var ROLE_SEARCH = map[string]interface{}{
	"d":  "domain_id",
	"rn": "role_name",
	"n":  "name",
}

type RoleDto struct {
	Id          int    `form:"id"`
	Name        string `form:"name" valid:"Required"`
	DomainId    int    `form:"domain_id" valid:"Required"`
	RoleName    string `form:"role_name" valid:"Required"`
	Remark      string `form:"remark"`
	MenusIds    string `form:"menu_ids"`
	MenusIdsEle string `form:"menu_ids_ele"`
}

// 如果你的 struct 实现了接口 validation.ValidFormer
// 当 StructTag 中的测试都成功时，将会执行 Valid 函数进行自定义验证
func (u *RoleDto) Valid(v *validation.Validation) {
	//if strings.Index(u.Name, "admin") != -1 {
	//	// 通过 SetError 设置 Name 的错误信息，HasErrors 将会返回 true
	//	v.SetError("Name", "名称里不能含有 admin")
	//}
}
