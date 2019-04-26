package dao

import (
	"database/sql"
	"github.com/astaxie/beego/orm"
	"github.com/bullteam/zeus/pkg/dto"
	"github.com/bullteam/zeus/pkg/models"
	"github.com/bullteam/zeus/pkg/utils"
)

type RoleDao struct {
}

func (dao *RoleDao) List(start int, limit int, q []string) ([]models.Role, int64) {
	o := GetOrmer()
	var roles []models.Role
	qs := o.QueryTable("role")
	if len(q) > 0 {
		for k, v := range utils.TransformFieldsCdt(q, dto.RoleSearch) {
			qs = utils.TransformQset(qs, k, v.(string))
		}
	}
	//后期加入搜索条件可利用q参数
	_, err := qs.Limit(limit, start).RelatedSel().All(&roles)
	if err != nil {
		return roles, 0
	}
	c, _ := qs.Count()

	return roles, c
}

func (dao *RoleDao) GetRolesAndDomainByUid(uid string) []orm.Params {
	var roles []orm.Params
	o := GetOrmer()
	o.Raw(`select r.*,d.code as domain_code,d.name as domain_name,d.callbackurl as domain_url from role r
				inner join user_role ur on ur.role_id=r.id
				inner join domain d on d.id=r.domain_id  where ur.user_id=?`, uid).Values(&roles)

	return roles
}

func (dao *RoleDao) GetRoleById(id int) (*models.Role, error) {
	o := GetOrmer()
	v := &models.Role{}
	err := o.QueryTable("role").Filter("id", id).RelatedSel().One(v)
	if err == nil {
		return v, nil
	}

	return nil, err
}

func (dao *RoleDao) GetRolesByUid(uid int) []orm.Params {
	var roles []orm.Params
	o := GetOrmer()
	o.Raw(`SELECT r.name,r.id FROM user_role ur
		   INNER JOIN role r ON r.id =  ur.role_id WHERE ur.user_id = ?`, uid).Values(&roles)

	return roles
}

func (dao *RoleDao) GetRoleByDid(rid int, domainId int) (*models.Role, error) {
	o := GetOrmer()
	v := &models.Role{}
	err := o.QueryTable("role").Filter("id", rid).Filter("domain_id", domainId).RelatedSel().One(v)
	if err == nil {
		return v, nil
	}

	return nil, err
}
func (dao *RoleDao) Create(roleEntity models.RoleEntity) (sql.Result, error) {
	o := GetOrmer()
	qs, _ := o.Raw("insert into role (name,domain_id,role_name,remark,menu_ids,menu_ids_ele) values (?,?,?,?,?,?)").Prepare()
	result, err := qs.Exec(roleEntity.Name, roleEntity.DomainId, roleEntity.RoleName, roleEntity.Remark, roleEntity.MenuIds, roleEntity.MenuIdsEle)

	return result, err
}

func (dao *RoleDao) Update(roleEntity models.RoleEntity) error {
	o := orm.NewOrm()
	qs, _ := o.Raw("update role set name=?,domain_id=?,role_name=?,remark=?,menu_ids=?,menu_ids_ele=? where id=?").Prepare()
	_, err := qs.Exec(roleEntity.Name, roleEntity.DomainId, roleEntity.RoleName, roleEntity.Remark, roleEntity.MenuIds, roleEntity.MenuIdsEle, roleEntity.Id)

	return err
}
func (dao *RoleDao) Delete(roleEntity models.RoleEntity) error {
	o := orm.NewOrm()
	qs, _ := o.Raw("delete from role where id=? limit 1").Prepare()
	_, err := qs.Exec(roleEntity.Id)

	return err
}
