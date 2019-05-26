package service

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/utils"
	"strconv"
	"strings"
	"time"
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

// 检查权限， 暂时不使用casbin检查权限方法，有分布式问题
func (ps *PermService) CheckPermByUidBACKUP(uid int, permission string, domain string) bool {
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

// 使用redis缓存，解决分布式问题
func (ps *PermService) CheckPermByUid(uid int, permission string, domain string) bool {
	userPermList := ps.GetUserDomainPerms(uid, domain)
	writeList := []interface{}{
		"/user/menu",
		"/user/perm/list",
		"/user/perm/check",
	}
	if utils.InSliceIface(permission, writeList) {
		return true
	}
	if userPermList != nil && len(userPermList) > 0 {
		return utils.InSliceIface(permission, userPermList)
	}
	return false
}

// 获取用户拥有的某个域的权限列表(路由)，缓存6小时
func (ps *PermService) GetUserDomainPerms(uid int, domainCode string) (permList []interface{}) {
	cacheKey := fmt.Sprintf("check_perm_%d_%s", uid, domainCode)
	rds := components.Cache
	if rds.IsExist(cacheKey) && rds.Get(cacheKey) != nil {
		rdsData := rds.Get(cacheKey)
		if rdsData != nil {
			if err := json.Unmarshal(rdsData.([]byte), &permList); err != nil {
				return nil
			}
		}
	} else {
		roles := ps.roleDao.GetRoleListAndDomain(uid, domainCode)
		if len(roles) > 0 {
			perm := components.NewPerm()
			for _, role := range roles {
				roleName := role["role_name"].(string)
				casbinRules := perm.GetAllPermByRoleName(roleName, domainCode)
				if len(casbinRules) > 0 {
					for _, casbinRule := range casbinRules {
						permList = append(permList, casbinRule[1])
					}
				}
			}
		}
		// 去重
		if len(permList) > 1 {
			permList = utils.SliceUnique(permList)
		}

		if len(permList) > 0 {
			permListStr, err := json.Marshal(permList)
			if err != nil {
				return nil
			}
			_ = rds.Put(cacheKey, permListStr, 6*3600*time.Second)
		}
	}

	return permList
}

func (ps *PermService) TransformPerm(route string) string {
	pos := strings.LastIndex(route, "/")
	newSli := []rune(route)
	newSli[pos] = ':'
	return string(newSli)
}
