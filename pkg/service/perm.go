package service

import (
	"strconv"
	"strings"
	"zeus/pkg/components"
	"zeus/pkg/dao"
)

type PermService struct {
	roleDao *dao.RoleDao
}

func (ps *PermService) GetPermsByRoleAndDomain(role string, domain string) [][]string {
	perm := components.NewPerm()
	return perm.GetAllPermByRole(role, domain)
}

func (ps *PermService) CheckPermByUid(uid int, permission string, domain string) bool {
	perm := components.NewPerm()
	roles := ps.roleDao.GetRolesByUid(uid)
	for _, r := range roles {
		rid, _ := strconv.Atoi(r["id"].(string))
		roleData, err := ps.roleDao.GetRoleById(rid)
		if err != nil || roleData == nil {
			return false
		}
		if perm.Check(roleData.RoleName, permission, "*", domain) {
			return true
		}
	}
	return false
}

func (ps *PermService) TransformPerm(route string) string {
	pos := strings.LastIndex(route, "/")
	newSli := []rune(route)
	newSli[pos] = ':'
	return string(newSli)
}
