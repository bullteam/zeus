package controllers

import (
	"zeus/pkg/components"
	"zeus/pkg/service"
	"strconv"
	"strings"
)

type PermController struct {
	TokenCheckController
}

func (pc *PermController) GetPermsByLoginUser() {
	userService := service.UserService{}
	roles := userService.GetRolesByUid(pc.Uid)
	permService := service.PermService{}
	roleService := service.RoleService{}
	var (
		results   []map[string]interface{} // 功能权限列表
		dataPerms []map[string]interface{} // 数据权限列表
	)
	var path = map[string]bool{}
	domain := pc.GetString("domain")
	for _, role := range roles {
		if domain != "" && domain != role["domain_code"] {
			continue
		}
		for _, perm := range permService.GetPermsByRoleAndDomain(role["role_name"].(string), role["domain_code"].(string)) {
			prefix := strings.Split(perm[1], ":")
			seg := strings.Split(prefix[0], "/")
			if len(seg) == 3 {
				if ok := path[seg[1]]; !ok {
					path[seg[1]] = true
					results = append(results, map[string]interface{}{
						"sub":    perm[0],
						"obj":    "/" + seg[1],
						"act":    "nil",
						"domain": perm[3],
					})
				}
				if ok := path[seg[2]]; !ok {
					path[seg[2]] = true
					results = append(results, map[string]interface{}{
						"sub":    perm[0],
						"obj":    prefix[0],
						"act":    "nil",
						"domain": perm[3],
					})
				}
			}
			results = append(results, map[string]interface{}{
				"sub":    perm[0],
				"obj":    perm[1],
				"act":    perm[2],
				"domain": perm[3],
			})
		}

		// 获取数据权限列表
		roleId, _ := strconv.Atoi(role["id"].(string))
		roleDataPerms, _ := roleService.GetRoleDataPermsByRoleId(roleId)
		if len(roleDataPerms) > 0 {
			for _, roleDataPerm := range roleDataPerms {
				dataPerms = append(dataPerms, map[string]interface{}{
					"id":    roleDataPerm["id"],
					"name":  roleDataPerm["name"],
					"perms": roleDataPerm["perms"],
				})
			}
			// 去重
			dataPerms = removeRepeatedDataPerm(dataPerms)
		}
	}
	pc.Resp(0, "success", map[string]interface{}{
		"result":     results,
		"data_perms": dataPerms,
		"info": map[string]interface{}{
			"id":       pc.Uid,
			"username": pc.Uname,
		},
	})
}

// 数据权限去重
func removeRepeatedDataPerm(arr []map[string]interface{}) (newArr []map[string]interface{}) {
	newArr = make([]map[string]interface{}, 0)
	for i := 0; i < len(arr); i++ {
		isRepeat := false
		for j := i + 1; j < len(arr); j++ {
			if arr[i]["id"] == arr[j]["id"] {
				isRepeat = true
				break
			}
		}
		if !isRepeat {
			newArr = append(newArr, arr[i])
		}
	}
	return
}

func (c *PermController) CheckPerm() {
	ps := service.PermService{}
	uid, _ := strconv.Atoi(c.Uid)
	perms := c.GetString("perm")
	domain := c.GetString("domain")
	if uid < 0 || !ps.CheckPermByUid(uid, perms, domain) {
		c.Fail(components.ErrPermission, "fail")
		return
	}
	c.Resp(0, "success")
}
