package controllers

import (
	"github.com/bullteam/zeus/pkg/components"
	"github.com/bullteam/zeus/pkg/service"
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
	var results []map[string]interface{}
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
	}
	pc.Resp(0, "success", map[string]interface{}{
		"result": results,
		"info": map[string]interface{}{
			"id":       pc.Uid,
			"username": pc.Uname,
		},
	})
}

func (c *PermController) CheckPerm() {
	ps := service.PermService{}
	uid, _ := strconv.Atoi(c.Uid)
	perms := c.GetString("perm")
	//pos   := strings.LastIndex(perms,"/")
	//new   := []rune(perms)
	//new[pos] = ':'
	//perms = string(new)
	domain := c.GetString("domain")
	if uid < 0 || !ps.CheckPermByUid(uid, perms, domain) {
		c.Fail(components.ErrPermission, "fail")
		return
	}
	c.Resp(0, "success")
}
