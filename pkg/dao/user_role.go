package dao

import (
	"github.com/astaxie/beego/orm"
	"zeus/pkg/models"
)

type UserRoleDao struct {
}

func (dao *UserRoleDao) Create(uid int, roleId int) (int64, error) {
	var user models.User
	user.Id = uid
	o := GetOrmer()
	_ = o.Read(&user, "id")

	var role models.Role
	role.Id = roleId
	_ = o.Read(&role, "id")

	ins := &models.UserRole{User: &user, Role: &role}

	return o.Insert(ins)
}

func (dao *UserRoleDao) DeleteByUid(uid int64) (int64, error) {
	o := orm.NewOrm()

	return o.QueryTable("user_role").Filter("user_id", uid).Delete()
}

func (dao *UserRoleDao) DeleteByRid(rid int64) (int64, error) {
	o := orm.NewOrm()

	return o.QueryTable("user_role").Filter("role_id", rid).Delete()
}
