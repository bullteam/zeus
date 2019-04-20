package tests

import (
	"github.com/astaxie/beego/validation"
	"github.com/bullteam/zeus/pkg/components"
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

//func TestRawCasbin(t *testing.T){
//	//orm.RegisterDriver("mysql", orm.DRMySQL)
//	a := beegoormadapter.NewAdapter("mysql", "root:100200Lzy-mysql@tcp(120.24.83.114:3306)/auth",true)
//	e := casbin.NewEnforcer("../conf/rbac_model.conf",a)
//	if e.Enforce("alice","data2","read","project1"){
//		t.Log("success")
//	}else{
//		t.Log("No permission")
//	}
//}
func TestMyCasbin(t *testing.T) {
	perm := components.NewPerm()
	if perm.Check("alice", "v1:ssss", "read", "project1") {
		t.Log("有权限")
	} else {
		t.Log("无权限")
	}
}

//func TestAddPermission(t *testing.T) {
//	perm := NewPerm()
//	if perm.AddPerm("liang.fu","menu1","view","project1"){
//		t.Log("权限创建成功")
//	}else{
//		t.Log("权限记录已存在")
//	}
//}
//
//func TestAddGroup(t *testing.T) {
//	perm := NewPerm()
//	if perm.AddGroup("liang.fu","group1"){
//		t.Log("用户成功关联用户组")
//	}else{
//		t.Log("关联记录已存在")
//	}
//}
//
//func TestAddGroupPermission(t *testing.T) {
//	perm := NewPerm()
//	if perm.AddPerm("group1","database","insert","project2"){
//		t.Log("组权限创建成功")
//	}else{
//		t.Log("组权限记录已存在")
//	}
//}

//func TestInitCasbinFromCategory(t *testing.T){
//	menus := models.Menu_list(1)
//	perm := NewPerm()
//	for _,v := range menus{
//		perm.AddPerm("superadmin",v["perms"],"*","root")
//	}
//	perm.AddPerm("role_test","*","*","parking")
//}

func TestGetAllByRole(t *testing.T) {
	perm := components.NewPerm()
	for _, v := range perm.GetAllPermByRole("superadmin", "root") {
		t.Log(v)
	}
}

func TestDeletePerm(t *testing.T) {
	perm := components.NewPerm()
	perm.DeleteRoleByDomain("role_test", "parking")
}

func TestValidate(t *testing.T) {
	type user struct {
		Id     int
		Name   string `form:"name" valid:"Required"` // Name 不能为空并且以 Bee 开头
		Age    int    `valid:"Range(1, 140)"`        // 1 <= Age <= 140，超出此范围即为不合法
		Email  string `valid:"Email; MaxSize(100)"`  // Email 字段需要符合邮箱格式，并且最大长度不能大于 100 个字符
		Mobile string `valid:"Mobile"`               // Mobile 必须为正确的手机号
		IP     string `valid:"IP"`                   // IP 必须为一个正确的 IPv4 地址
	}
	valid := validation.Validation{}
	u := &user{Name: "", Age: 2, Email: "dev@beego.me"}
	b, err := valid.Valid(u)
	if err != nil {
		// handle error
	}
	if !b {
		// validation does not pass
		// blabla...
		for _, err := range valid.Errors {
			t.Logf("key:%s,msg:%s", err.Key, err.Message)
		}
	}
}
