package components

import (
	"github.com/casbin/casbin"
	//"github.com/casbin/beego-orm-adapter"
	"github.com/astaxie/beego"
	"github.com/funlake/beego-orm-adapter"
	_ "github.com/go-sql-driver/mysql"
	"sync"
	//"strings"
	"path/filepath"
	"runtime"
	"strings"
)

var (
	permSync = &sync.Once{}
	permOnce *perm
)

func NewPerm() *perm {
	permSync.Do(func() {
		//for unit testing
		confFilePath := beego.AppPath
		//if beego.AppConfig.String("mysqlurls") == "" {
		if strings.Index(confFilePath, "/tmp") == 0 {
			//confFilePath = "../.."
			_, file, _, _ := runtime.Caller(0)
			//beego.Warning(file)
			confFilePath = filepath.Dir(filepath.Join(file, ".."+string(filepath.Separator)))
			//err := beego.LoadAppConfig("ini", confFilePath+"/conf/app.conf")
			//if err != nil{
			//	beego.Error(err.Error())
			//}
		}
		mysqluser := beego.AppConfig.String("mysqluser")
		mysqlpass := beego.AppConfig.String("mysqlpass")
		mysqlurls := beego.AppConfig.String("mysqlurls")
		mysqlport := beego.AppConfig.String("mysqlport")
		mysqldb := beego.AppConfig.String("mysqldbcasbin")
		a := beegoormadapter.NewAdapter("mysql", mysqluser+":"+mysqlpass+"@tcp("+mysqlurls+":"+mysqlport+")/"+mysqldb+"?charset=utf8mb4", true)
		permOnce = &perm{
			casbin.NewEnforcer(confFilePath+"/conf/rbac_model.conf", a),
		}
		//permOnce.enforcer.EnableAutoSave(true)
	})

	return permOnce
}

type perm struct {
	enforcer *casbin.Enforcer
}

//first : user
//second : group
func (p *perm) AddGroup(params ...interface{}) bool {
	return p.enforcer.AddGroupingPolicy(params...)
}

//sub,obj,act,domain
func (p *perm) AddPerm(params ...interface{}) bool {
	return p.enforcer.AddPolicy(params...)
}
func (p *perm) Check(params ...interface{}) bool {
	return p.enforcer.Enforce(params...)
}

func (p *perm) DeleteRoleByDomain(role string, domain string) {
	p.enforcer.RemoveFilteredNamedPolicy("p", 0, role, "", "", domain)
	//o := orm.NewOrm()
	//e,_ := o.Raw("delete from casbin_rule where p_type=? and v0=? and v3=?").Prepare()
	//e.Exec("p",role,domain)
}
func (p *perm) DeleteRole(role string) {
	p.enforcer.RemoveFilteredNamedPolicy("p", 0, role)
}
func (p *perm) GetAllPermByRole(role string, domain string) [][]string {
	p.enforcer.LoadPolicy()
	roles := p.enforcer.GetFilteredNamedPolicy("p", 0, role, "", "", domain)
	return roles
}

//dangerous! do not call until you really need it
func (p *perm) CommitChange() {
	p.enforcer.SavePolicy()
}
