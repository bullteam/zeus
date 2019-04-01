package models

import (
	"github.com/astaxie/beego/orm"
	"database/sql"
	"github.com/bullteam/zeus/utils"
	"github.com/bullteam/zeus/dto"
)

func init() {
	orm.RegisterModel(new(Role))
}

//for列表,有关联关系
type Role struct {
	Id int `json:"id"`
	//DomainId    string `json:"domain_id"`
	Name     string  `json:"name"`
	Domain   *Domain `json:"domain" orm:"rel(one)"`
	RoleName string  `json:"role_name"`
	Remark   string  `json:"remark"`
	Users    []*User `json:"users" orm:"reverse(many)"`
	MenuIds  string  `json:"menu_ids"`
	MenuIdsEle string `json:"menu_ids_ele"`
}
//for更新创建
type RoleEntity struct {
	Id int `json:"id"`
	//DomainId    string `json:"domain_id"`
	Name     string `json:"name"`
	DomainId  int  `json:"domain_id"`
	RoleName string `json:"role_name"`
	Remark   string `json:"remark"`
	MenuIds  string `json:"menu_ids"`
	MenuIdsEle string `json:"menu_ids_ele"`
}
func (r Role) List(start int, limit int, q []string) ([]Role, int64) {
	o := orm.NewOrm()
	var roles []Role
	qs := o.QueryTable("role")
	//qs.RelatedSel("domain","domain_id","id")
	if len(q) > 0 {
		for k, v := range utils.TransformFieldsCdt(q, dto.ROLE_SEARCH) {
			qs = utils.TransformQset(qs,k,v.(string))
			//qs = qs.Filter(k, v)
		}
	}
	//后期加入搜索条件可利用q参数
	qs.Limit(limit, start).RelatedSel().All(&roles)
	c, _ := qs.Count()
	return roles, c
}
func (r Role) GetRolesAndDomainByUid(uid string) []orm.Params {
	var roles []orm.Params
	o := orm.NewOrm()
	o.Raw(`select r.*,d.code as domain_code,d.name as domain_name,d.callbackurl as domain_url from role r
				inner join user_role ur on ur.role_id=r.id
				inner join domain d on d.id=r.domain_id  where ur.user_id=?`, uid).Values(&roles)
	return roles
}
func (r Role)GetRoleById(id int) (*Role, error) {
	o := orm.NewOrm()
	v := &Role{}
	err1:= o.QueryTable("role").Filter("id",id).RelatedSel().One(v)
	if err1 == nil {
		return v, nil
	}
	return nil, err1
}

func (r Role) GetRolesByUid(uid int) []orm.Params {
	var roles []orm.Params
	o := orm.NewOrm()
	o.Raw(`SELECT r.name,r.id FROM user_role ur
		   INNER JOIN role r ON r.id =  ur.role_id WHERE ur.user_id = ?`, uid).Values(&roles)
	return roles
}
func (r Role) GetRoleByDid(rid int,domain_id int) (*Role, error){
	o := orm.NewOrm()
	v := &Role{}
	err:= o.QueryTable("role").Filter("id",rid).Filter("domain_id",domain_id).RelatedSel().One(v)
	if err == nil {
		return v, nil
	}
	return nil, err
}
func (r Role) Create(roleEntity RoleEntity)(sql.Result,error){
	o := orm.NewOrm()
	qs,_ := o.Raw("insert into role (name,domain_id,role_name,remark,menu_ids,menu_ids_ele) values (?,?,?,?,?,?)").Prepare()
	result,err := qs.Exec(roleEntity.Name,roleEntity.DomainId,roleEntity.RoleName,roleEntity.Remark,roleEntity.MenuIds,roleEntity.MenuIdsEle)
	return result,err
}

func (r Role) Update(roleEntity RoleEntity) error{
	o := orm.NewOrm()
	qs,_ := o.Raw("update role set name=?,domain_id=?,role_name=?,remark=?,menu_ids=?,menu_ids_ele=? where id=?").Prepare()
	_,err := qs.Exec(roleEntity.Name,roleEntity.DomainId,roleEntity.RoleName,roleEntity.Remark,roleEntity.MenuIds,roleEntity.MenuIdsEle,roleEntity.Id)
	return err
}
func (r Role) Delete(roleEntity RoleEntity) error{
	o := orm.NewOrm()
	qs,_ := o.Raw("delete from role where id=? limit 1").Prepare()
	_,err := qs.Exec(roleEntity.Id)
	return err
}