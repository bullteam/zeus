package dao

import (
	"errors"
	"github.com/astaxie/beego/orm"
	"zeus/pkg/dto"
	"zeus/pkg/models"
)

type DataPermDao struct {
}

// 获取数据权限列表
func (dao *DataPermDao) GetDataPermList(query *models.DataPermQuery) ([]models.DataPerm, int64) {
	var (
		dataPermList []models.DataPerm
		offset       int
	)
	qs := dataPermQueryConditions(query)

	total, _ := qs.Count()
	if query.Pagination != nil {
		offset = (query.Pagination.Start - 1) * query.Pagination.Limit
		qs = qs.Offset(offset).Limit(query.Pagination.Limit)
	}

	_, err := qs.OrderBy("order_num").RelatedSel().All(&dataPermList)
	if err != nil {
		return dataPermList, 0
	}

	return dataPermList, total
}

func (dao *DataPermDao) GetById(id int) (models.DataPerm, error) {
	o := GetOrmer()
	dataPerm := models.DataPerm{}
	err := o.QueryTable("data_perm").Filter("id",id).RelatedSel().One(&dataPerm)

	return dataPerm, err
}

func (dao *DataPermDao) GetByName(name string) (dataPerm models.DataPerm, err error) {
	o := GetOrmer()
	err = o.QueryTable(&models.DataPerm{}).Filter("name", name).One(&dataPerm)

	return dataPerm, err
}

func (dao *DataPermDao) GetByPerms(perms string) (dataPerm models.DataPerm, err error) {
	o := GetOrmer()
	err = o.QueryTable(&models.DataPerm{}).Filter("perms", perms).One(&dataPerm)

	return dataPerm, err
}

func (dao *DataPermDao) Insert(dto *dto.DataPermAddDto) (int64, error) {
	// 判断name是否存在
	_, err := dao.GetByName(dto.Name)
	if err == nil {

		return 0, errors.New("name is exist")
	}
	// 判断perms是否存在
	_, err = dao.GetByName(dto.Perms)
	if err == nil {

		return 0, errors.New("perms is exist")
	}

	o := GetOrmer()
	qs, _ := o.Raw("insert into " + models.TableDataPerm + " (domain_id,parent_id,name,perms,perms_rule,perms_type,order_num,remarks) values (?,?,?,?,?,?,?,?)").Prepare()
	result, err := qs.Exec(dto.DomainId, dto.ParentId, dto.Name, dto.Perms, dto.PermsRule, dto.PermsType, dto.OrderNum, dto.Remarks)
	id, _ := result.LastInsertId()

	return id, err
}

func (dao *DataPermDao) Update(dto *dto.DataPermEditDto) error {
	id := dto.Id
	dataPerm, err := dao.GetById(id)

	if err != nil {
		return errors.New("data perms is not exist")
	}

	// 判断name是否存在
	data, err := dao.GetByName(dto.Name)
	if err == nil && data.Id != dataPerm.Id {

		return errors.New("name is exist")
	}
	// 判断perms是否存在
	data, err = dao.GetByName(dto.Perms)
	if err == nil && data.Id != dataPerm.Id {

		return errors.New("perms is exist")
	}

	o := GetOrmer()
	dataPerm.Domain.Id = dto.DomainId
	dataPerm.ParentId = dto.ParentId
	dataPerm.Name = dto.Name
	dataPerm.Perms = dto.Perms
	dataPerm.PermsRule = dto.PermsRule
	dataPerm.PermsType = dto.PermsType
	dataPerm.OrderNum = dto.OrderNum
	dataPerm.Remarks = dto.Remarks

	_, err = o.Update(&dataPerm)

	return err
}

func (dao *DataPermDao) Delete(id int) error {
	o := GetOrmer()
	err := o.Begin()
	if err != nil {
		return err
	}

	num, err := o.Delete(&models.DataPerm{Id: id})
	if num == 0 {
		_ = o.Rollback()
		return errors.New("delete failed")
	}
	if err != nil {
		_ = o.Rollback()
	}
	// 删除已分配给角色的
	roleDataPermDao := RoleDataPermDao{}
	_, err = roleDataPermDao.DeleteByDataPermId(id)
	if err != nil {
		_ = o.Rollback()
	}
	_ = o.Commit()

	return nil
}

func dataPermQueryConditions(query *models.DataPermQuery) orm.QuerySeter {
	qs := GetOrmer().QueryTable(&models.DataPerm{}).Filter("domain_id", query.DomainId)

	if len(query.Name) > 0 {
		qs = qs.Filter("name__contains", query.Name)
	}

	return qs
}
