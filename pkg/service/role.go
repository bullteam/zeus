package service

import (
	"database/sql"
	"github.com/astaxie/beego/orm"
	"github.com/bullteam/zeus/pkg/components"
	"github.com/bullteam/zeus/pkg/dto"
	"github.com/bullteam/zeus/pkg/models"
	"strconv"
	"strings"
)

type RoleService struct {
}

//GetList
func (r RoleService) GetList(start int, limit int, q []string) ([]models.Role, int64) {
	role := models.Role{}
	return role.List(start, limit, q)
}

//Show
func (r RoleService) GetRoleById(role_id int) (v *models.Role, perms [][]string, err error) {
	role := models.Role{}
	perm := components.NewPerm()
	roleRs, err := role.GetRoleById(role_id)
	return roleRs, perm.GetAllPermByRole(roleRs.RoleName, roleRs.Domain.Code), err
}

func (r RoleService) GetRoleByDomainId(rid int, domain_id int) (*models.Role, error) {
	role := models.Role{}
	return role.GetRoleByDid(rid, domain_id)
}

func (r RoleService) CreateRole(dto *dto.RoleDto) (sql.Result, error) {
	role := models.Role{}
	return role.Create(models.RoleEntity{
		Name:       dto.Name,
		DomainId:   dto.DomainId,
		RoleName:   dto.RoleName,
		Remark:     dto.Remark,
		MenuIds:    dto.MenusIds,
		MenuIdsEle: dto.MenusIdsEle,
	})
}

func (r RoleService) UpdateRole(dto *dto.RoleDto) error {
	role := models.Role{}
	return role.Update(models.RoleEntity{
		Id:         dto.Id,
		Name:       dto.Name,
		DomainId:   dto.DomainId,
		RoleName:   dto.RoleName,
		Remark:     dto.Remark,
		MenuIds:    dto.MenusIds,
		MenuIdsEle: dto.MenusIdsEle,
	})
}

func (r RoleService) DeleteRole(id int) error {
	role := models.Role{}
	userRole := models.UserRole{}
	roleData, err := role.GetRoleById(id)
	if err != nil {
		return err
	}
	//1.删除数据库本表记录
	err = role.Delete(models.RoleEntity{
		Id: id,
	})
	if err != nil {
		return err
	}
	//2.删除user_role关联记录
	_, err = userRole.DeleteByRid(int64(roleData.Id))
	if err != nil {
		return err
	}
	//3.删除casbin权限记录
	perm := components.NewPerm()
	perm.DeleteRoleByDomain(roleData.RoleName, roleData.Domain.Code)
	return nil
}

func (r RoleService) AssignPerm(domain_id int, role_id int, menu_ids string) error {
	roleModel := models.Role{}
	domain, err := models.GetDomain(domain_id)
	if err != nil {
		return err
	}
	role, err := roleModel.GetRoleById(role_id)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	perm := components.NewPerm()
	//1.删除旧有权限
	perm.DeleteRoleByDomain(role.RoleName, domain.Code)
	for _, v := range strings.Split(menu_ids, ",") {
		mid, _ := strconv.Atoi(v)
		menu := &models.Menu{Id: mid}
		if err := o.Read(menu); err != nil {
			return err
		}
		if menu.Perms != "" {
			//2.重新汇入权限
			//beego.Info(role.RoleName, menu.Perms, "*", domain.Code)
			perm.AddPerm(role.RoleName, menu.Perms, "*", domain.Code)
		}
	}
	//perm.CommitChange()
	return nil
}
