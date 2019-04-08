package service

import (
	"github.com/bullteam/zeus/pkg/components"
	"github.com/bullteam/zeus/pkg/models"
	"strconv"
	"strings"
)

type PermService struct{

}
func (ps PermService) GetPermsByRoleAndDomain(role string , domain string) [][]string {
	perm := components.NewPerm()
	return perm.GetAllPermByRole(role,domain)
}

func (ps PermService) CheckPermByUid(uid int,permission string,domain string) bool{
	role := models.Role{}
	perm := components.NewPerm()
	roles := role.GetRolesByUid(uid)
	for _,r := range roles {
		rid,_ := strconv.Atoi(r["id"].(string))
		roleData,err := role.GetRoleById(rid)
		if err != nil || roleData == nil{
			return false
		}
		if perm.Check(roleData.RoleName,permission,"*",domain){
			return true
		}
	}
	return false
}

func (ps PermService) TransformPerm(route string) string{
	pos   := strings.LastIndex(route,"/")
	new   := []rune(route)
	new[pos] = ':'
	return string(new)
}