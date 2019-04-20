package service

import (
	"github.com/astaxie/beego/orm"
	"github.com/bullteam/zeus/pkg/dto"
	"github.com/bullteam/zeus/pkg/models"
	"strconv"
)

type UserService struct {
}

func (us UserService) GetList(start, limit int, q []string) ([]*models.User, int64) {
	user := models.User{}
	return user.List(start, limit, q)
}
func (us UserService) GetRolesByUid(uid string) []orm.Params {
	role := models.Role{}
	return role.GetRolesAndDomainByUid(uid)
}
func (us UserService) AddRoles(uid int64, roles []string) bool {
	userRole := models.UserRole{}
	//1.remove previously connections
	userRole.DeleteByUid(uid)
	for _, role := range roles {
		rid, err := strconv.Atoi(role)
		if err == nil {
			//create new connections
			//ignore duplicated record error
			userRole.Create(int(uid), rid)
		} else {
			return false
		}
	}
	return true
}

//获取用户相关项目域
func (us UserService) GetRelatedDomains(uid string) []map[string]interface{} {
	//menuService := MenuService{}
	//domainModel := models.Domain{}
	roleModel := models.Role{}
	domains := roleModel.GetRolesAndDomainByUid(uid)
	var d []map[string]interface{}
	s := map[interface{}]bool{}
	for _, domain := range domains {
		if _, ok := s[domain["domain_code"]]; !ok && domain["domain_code"] != "root" {
			d = append(d, map[string]interface{}{
				"name": domain["domain_name"],
				"code": domain["domain_code"],
				"url":  domain["domain_url"],
			})
			s[domain["domain_code"]] = true
		}
	}
	return d
}
func (us UserService) GetMenusByDomain(uid string, domain string) []orm.Params {
	//roleService := RoleService{}
	menuService := MenuService{}
	domainModel := models.Domain{}
	roleModel := models.Role{}
	roles := roleModel.GetRolesAndDomainByUid(uid)
	//beego.Info(roles)
	var menus string
	d, err := domainModel.GetDomainByCode(domain)
	if err != nil {
		return nil
	}
	for _, role := range roles {
		if role["domain_id"].(string) == strconv.Itoa(d.Id) {
			if menus != "" {
				menus += ","
			}
			menus += role["menu_ids"].(string)
		}
	}
	if menus != "" {
		return menuService.GetMenusByIds(menus)
	}
	return nil
}

//修改自己的密码
func (us UserService) ResetPassword(uid int, dto *dto.PwdResetDto) error {
	return models.UpdatePassword(uid, dto.NewRePwd)
}

//修改其它用户的密码
func (us UserService) ResetUserPassword(uid int, dto *dto.PwdUserResetDto) error {
	return models.UpdatePassword(uid, dto.NewRePwd)
}
func (us UserService) SwitchDepartment(uids []string, did int) (int64, error) {
	euid := []interface{}{}
	for _, s := range uids {
		euid = append(euid, s)
	}
	return models.UpdateDepartment(euid, did)
}
