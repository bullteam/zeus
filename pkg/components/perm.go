package components

import (
	"github.com/astaxie/beego/logs"
	"github.com/casbin/casbin"
	"path/filepath"
	"sync"
)

var (
	permSync = &sync.Once{}
	permOnce *perm
)

// 监听权限是否有变化，有变化则重新加载到内存
func updateCallback(msg string) {
	logs.Info(msg)
	_ = permOnce.enforcer.LoadPolicy()
}

func NewPerm() *perm {
	permSync.Do(func() {
		rbacmodelconf, err := filepath.Abs(Args.ConfigFile + "/rbac_model.conf")
		if err != nil {
			return
		}
		a := NewAdapter()
		permOnce = &perm{
			casbin.NewEnforcer(rbacmodelconf, a),
		}

		// 设置观察者，实现分布式
		/*redisHost := beego.AppConfig.String("redis_conn")
		redisPort := beego.AppConfig.String("redis_port")
		redisPws := beego.AppConfig.String("redis_pwd")
		host := fmt.Sprintf("%s:%s", redisHost, redisPort)
		w, _ := rediswatcher.NewWatcher(host, rediswatcher.Password(redisPws))
		permOnce.enforcer.SetWatcher(w)
		_ = w.SetUpdateCallback(updateCallback)*/
		// permOnce.enforcer.EnableAutoSave(true)
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

func (p *perm) DelPerm(params ...interface{}) bool {
	return p.enforcer.RemovePolicy(params...)
}

func (p *perm) Check(params ...interface{}) bool {
	return p.enforcer.Enforce(params...)
}

func (p *perm) DeleteRoleByDomain(role string, domain string) {
	p.enforcer.RemoveFilteredNamedPolicy("p", 0, role, "", "", domain)
}

func (p *perm) DeleteRole(role string) {
	p.enforcer.RemoveFilteredNamedPolicy("p", 0, role)
}

// 通过角色和域获取权限列表 并载入到内存
func (p *perm) GetAllPermByRole(role string, domain string) [][]string {
	p.enforcer.LoadPolicy()
	roles := p.enforcer.GetFilteredNamedPolicy("p", 0, role, "", "", domain)
	return roles
}

// 通过角色和域获取权限列表 不载入内存
func (p *perm) GetAllPermByRoleName(role string, domain string) [][]string {
	roles := p.enforcer.GetFilteredNamedPolicy("p", 0, role, "", "", domain)
	return roles
}

//dangerous! do not call until you really need it
func (p *perm) CommitChange() {
	p.enforcer.SavePolicy()
}
