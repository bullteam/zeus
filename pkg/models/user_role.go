package models

import orm "github.com/astaxie/beego/orm"
func init() {
	orm.RegisterModel(new(UserRole))
}
type UserRole struct{
	Id int
	User    *User    `orm:"rel(fk)"`
	Role    *Role    `orm:"rel(fk)"`
}

func (ur UserRole) Create(uid int,role_id int)(int64,error){
	var user  User
	user.Id = uid
	orm.NewOrm().Read(&user, "id")

	var role Role
	role.Id = role_id
	orm.NewOrm().Read(&role, "id")

	o := orm.NewOrm()
	ins := &UserRole{User:&user,Role:&role}
	return o.Insert(ins)
}
func (ur UserRole) DeleteByUid(uid int64)(int64,error){
	o := orm.NewOrm()
	return o.QueryTable("user_role").Filter("user_id",uid).Delete()
}
func (ur UserRole) DeleteByRid(rid int64)(int64,error){
	o := orm.NewOrm()
	return o.QueryTable("user_role").Filter("role_id",rid).Delete()
}