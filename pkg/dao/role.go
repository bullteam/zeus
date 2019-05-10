package dao

import (
	"database/sql"
	"github.com/astaxie/beego/orm"
	"github.com/bullteam/zeus/pkg/dto"
	"github.com/bullteam/zeus/pkg/models"
	"github.com/bullteam/zeus/pkg/utils"
	"strconv"
	"strings"
)

type RoleDao struct {
}

func (dao *RoleDao) List(start int, limit int, q []string) ([]orm.Params, int64) {
	o := GetOrmer()
	var (
		roles []models.Role
		list  []orm.Params
	)
	qs := o.QueryTable("role").OrderBy("id")
	if len(q) > 0 {
		for k, v := range utils.TransformFieldsCdt(q, dto.RoleSearch) {
			qs = utils.TransformQset(qs, k, v.(string))
		}
	}
	//后期加入搜索条件可利用q参数
	_, err := qs.Limit(limit, start).RelatedSel().All(&roles)
	if len(roles) > 0 {
		for _, v := range roles {
			var dataPermIds []string
			tmp := map[string]interface{}{
				"id":           v.Id,
				"name":         v.Name,
				"role_name":    v.RoleName,
				"remark":       v.RoleName,
				"menu_ids":     v.MenuIds,
				"menu_ids_ele": v.MenuIdsEle,
			}
			tmpDomain := map[string]interface{}{
				"id":           v.Domain.Id,
				"name":         v.Domain.Name,
				"callbackurl":  v.Domain.Callbackurl,
				"remark":       v.Domain.Remark,
				"code":         v.Domain.Code,
				"created_time": v.Domain.CreateTime,
				"updated_time": v.Domain.LastUpdateTime,
			}
			tmp["domain"] = tmpDomain
			// 获取数据权限
			roleDataPermDao := RoleDataPermDao{}
			roleDataPerms, _ := roleDataPermDao.GetByRoleId(v.Id)
			if len(roleDataPerms) > 0 {
				for _, rdp := range roleDataPerms {
					dataPermIds = append(dataPermIds, rdp["id"].(string))
				}
			}

			if len(dataPermIds) > 0 {
				tmp["data_perm_ids"] = strings.Join(dataPermIds, ",")
			} else {
				tmp["data_perm_ids"] = ""
			}

			list = append(list, tmp)
		}
	}
	if err != nil {
		return list, 0
	}
	c, _ := qs.Count()

	return list, c
}

// 暂时没用到
func (dao *RoleDao) GetList(start int, limit int, q []string) ([]orm.Params, int) {
	o := GetOrmer()
	var (
		list      []orm.Params
		total     int
		roles     []orm.Params
		sqlstr    string
		countSql  string
		countData []orm.Params
	)
	roleTable := "role"
	domainTable := models.TableDomain

	sqlstr = "select r.*,d.id as did,d.name,d.callbackurl,d.remark,d.code from " + roleTable + " as r left join " + domainTable + " as d on r.domain_id=d.id"
	countSql = "select count(*) as total from " + roleTable + " as r left join " + domainTable + " as d on r.domain_id=d.id"

	_, _ = o.Raw(sqlstr).Values(&roles)
	_, _ = o.Raw(countSql).Values(&countData)

	if len(roles) > 0 {
		for _, v := range roles {
			v["domain"] = map[string]interface{}{
				"id":          v["did"],
				"name":        v["name"],
				"callbackurl": v["callbackurl"],
				"remark":      v["remark"],
				"code":        v["code"],
			}
			list = append(list, v)
		}
		totalStr := countData[0]["total"].(string)
		total, _ = strconv.Atoi(totalStr)
	}

	return list, total
}

func (dao *RoleDao) GetRolesAndDomainByUid(uid string) []orm.Params {
	var roles []orm.Params
	o := GetOrmer()
	_, _ = o.Raw(`select r.*,d.code as domain_code,d.name as domain_name,d.callbackurl as domain_url from role r
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
	_, _ = o.Raw(`SELECT r.name,r.id FROM user_role ur
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
