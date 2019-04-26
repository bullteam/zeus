package dto

import "github.com/astaxie/beego/validation"

var DepartmentSearch = map[string]interface{}{
	"n": "name",
}

type DepartmentAddDto struct {
	Name     string `form:"name" valid:"Required"`
	ParentId int    `form:"parent_id"`
	OrderNum int    `form:"order_num"`
}

type DepartmentEditDto struct {
	Id       int    `form:"id" valid:"Required;Min(1)"`
	Name     string `form:"name" valid:"Required"`
	ParentId int    `form:"parent_id"`
	OrderNum int    `form:"order_num" valid:"Min(1)"`
}

// 如果你的 struct 实现了接口 validation.ValidFormer
// 当 StructTag 中的测试都成功时，将会执行 Valid 函数进行自定义验证
func (u *DepartmentAddDto) Valid(v *validation.Validation) {
	//if strings.Index(u.Name, "admin") != -1 {
	//	// 通过 SetError 设置 Name 的错误信息，HasErrors 将会返回 true
	//	v.SetError("Name", "名称里不能含有 admin")
	//}
}

func (u *DepartmentEditDto) Valid(v *validation.Validation) {
	//if strings.Index(u.Name, "admin") != -1 {
	//	// 通过 SetError 设置 Name 的错误信息，HasErrors 将会返回 true
	//	v.SetError("Name", "名称里不能含有 admin")
	//}
}
