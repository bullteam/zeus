package service

import (
	"database/sql"
	"github.com/astaxie/beego/orm"
	"zeus/pkg/components"
	"zeus/pkg/dao"
	"zeus/pkg/dto"
	"zeus/pkg/models"
	"strconv"
	"strings"
)

type RoleService struct {
	dao       *dao.RoleDao
	domainDao *dao.DomainDao
	urDao     *dao.UserRoleDao
	rdpDao    *dao.RoleDataPermDao
}

func (r *RoleService) GetList(start int, limit int, q []string) ([]orm.Params, int64) {
	return r.dao.List(start, limit, q)
}

func (r *RoleService) GetRoleById(roleId int) (v *models.Role, perms [][]string, err error) {
	perm := components.NewPerm()
	roleRs, err := r.dao.GetRoleById(roleId)
	if err != nil {
		return roleRs, perms, err
	}
	return roleRs, perm.GetAllPermByRole(roleRs.RoleName, roleRs.Domain.Code), err
}

func (r *RoleService) GetRoleByDomainId(rid int, domainId int) (*models.Role, error) {
	return r.dao.GetRoleByDid(rid, domainId)
}

func (r *RoleService) CreateRole(dto *dto.RoleDto) (sql.Result, error) {
	return r.dao.Create(models.RoleEntity{
		Name:       dto.Name,
		DomainId:   dto.DomainId,
		RoleName:   dto.RoleName,
		Remark:     dto.Remark,
		MenuIds:    dto.MenusIds,
		MenuIdsEle: dto.MenusIdsEle,
	})
}

func (r *RoleService) UpdateRole(dto *dto.RoleDto) error {
	return r.dao.Update(models.RoleEntity{
		Id:         dto.Id,
		Name:       dto.Name,
		DomainId:   dto.DomainId,
		RoleName:   dto.RoleName,
		Remark:     dto.Remark,
		MenuIds:    dto.MenusIds,
		MenuIdsEle: dto.MenusIdsEle,
	})
}

func (r *RoleService) DeleteRole(id int) error {
	roleData, err := r.dao.GetRoleById(id)
	if err != nil {
		return err
	}
	//1.删除数据库本表记录
	err = r.dao.Delete(models.RoleEntity{
		Id: id,
	})
	if err != nil {
		return err
	}
	//2.删除user_role关联记录
	_, err = r.urDao.DeleteByRid(int64(roleData.Id))
	if err != nil {
		return err
	}
	//3.删除casbin权限记录
	perm := components.NewPerm()
	perm.DeleteRoleByDomain(roleData.RoleName, roleData.Domain.Code)

	//4.删除数据权限
	_ = r.rdpDao.DeleteByRoleId(id)
	return nil
}

// 分配功能权限
func (r *RoleService) AssignPerm(domainId int, roleId int, menuIds string) error {
	domain, err := r.domainDao.GetDomain(domainId)
	if err != nil {
		return err
	}
	role, err := r.dao.GetRoleById(roleId)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	perm := components.NewPerm()
	//1.删除旧有权限
	perm.DeleteRoleByDomain(role.RoleName, domain.Code)
	for _, v := range strings.Split(menuIds, ",") {
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

// 分配数据权限
func (r *RoleService) AssignDataPerm(roleId int, dataPermIds string) error {
	var (
		dtos           []dto.AssignDataPermDto
		dtoOne         dto.AssignDataPermDto
		oldDataPermIds []int
		dataIds        []string
	)

	// 查询旧的数据权限列表
	oldRoleDataPerms, _ := r.rdpDao.GetByRoleId(roleId)

	// 删除该角色所有旧的数据权限再插入新的
	if len(oldRoleDataPerms) > 0 {
		for _, v := range oldRoleDataPerms {
			tmpId, _ := strconv.Atoi(v["id"].(string))
			oldDataPermIds = append(oldDataPermIds, tmpId)
		}
		_ = r.rdpDao.DeleteMulti(roleId, oldDataPermIds)
	}
	// 插入新的数据权限
	dataIds = strings.Split(dataPermIds, ",")
	if len(dataIds) > 0 {
		for _, v := range dataIds {
			tmpId, _ := strconv.Atoi(v)
			dtoOne.RoleId = roleId
			dtoOne.DataPermId = tmpId
			dtos = append(dtos, dtoOne)
		}
		_ = r.rdpDao.InsertMulti(dtos)
	}

	return nil
}

func (r *RoleService) GetRolesByUid(uid int) []orm.Params {
	return r.dao.GetRolesByUid(uid)
}

// 通过角色id获取数据权限列表
func (r *RoleService) GetRoleDataPermsByRoleId(roleId int) ([]orm.Params, error) {
	return r.rdpDao.GetByRoleId(roleId)
}
