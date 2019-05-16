package dao

import (
	"github.com/astaxie/beego/orm"
	"zeus/pkg/dto"
	"zeus/pkg/models"
)

type RoleDataPermDao struct {
}

// 分配数据权限
func (dao *RoleDataPermDao) InsertMulti(dtos []dto.AssignDataPermDto) error {
	o := GetOrmer()
	var roleDataPerms []models.RoleDataPerm

	for _, v := range dtos {
		roleDataPerm := models.RoleDataPerm{}
		roleDataPerm.RoleId = v.RoleId
		roleDataPerm.DataPermId = v.DataPermId
		roleDataPerms = append(roleDataPerms, roleDataPerm)
	}
	_, err := o.InsertMulti(len(roleDataPerms), roleDataPerms)

	return err
}

// 批量删除
func (dao *RoleDataPermDao) DeleteMulti(roleId int, dataPermIds []int) error {
	o := GetOrmer()
	for _, v := range dataPermIds {
		_, _ = o.QueryTable(new(models.RoleDataPerm)).Filter("role_id", roleId).Filter("data_perm_id", v).Delete()
	}

	return nil
}

// 通过数据权限id删除
func (dao *RoleDataPermDao) DeleteByDataPermId(dataPermId int) (int64, error) {
	o := GetOrmer()
	num, err := o.QueryTable(new(models.RoleDataPerm)).Filter("data_perm_id", dataPermId).Delete()

	return num, err
}

// 通过角色id删除所有数据权限
func (dao *RoleDataPermDao) DeleteByRoleId(roleId int) error {
	o := GetOrmer()
	_, err := o.QueryTable(new(models.RoleDataPerm)).Filter("role_id", roleId).Delete()

	return err
}

// 根据角色id查找数据权限
func (dao *RoleDataPermDao) GetByRoleId(roleId int) (rdps []orm.Params, err error) {
	o := GetOrmer()
	sql := "select rdp.role_id,dp.name,dp.perms,dp.id from role_data_perm as rdp left join data_perm as dp on rdp.data_perm_id=dp.id where rdp.role_id=?"
	_, err = o.Raw(sql, roleId).Values(&rdps)

	return rdps, err
}
