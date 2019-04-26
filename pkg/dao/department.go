package dao

import (
	"errors"
	"github.com/bullteam/zeus/pkg/dto"
	"github.com/bullteam/zeus/pkg/models"
	"github.com/bullteam/zeus/pkg/utils"
)

type DepartmentDao struct {
}

func (dao *DepartmentDao) List(start int, limit int, q []string) ([]*models.Department, int64) {
	o := GetOrmer()
	var depts []*models.Department
	qs := o.QueryTable("department").OrderBy("order_num")
	if len(q) > 0 {
		for k, v := range utils.TransformFieldsCdt(q, dto.DepartmentSearch) {
			qs = utils.TransformQset(qs, k, v.(string))
		}
	}
	//后期加入搜索条件可利用q参数
	_, err := qs.Limit(limit, start).RelatedSel().All(&depts)
	if err != nil {
		return depts, 0
	}
	total, _ := qs.Count()

	return depts, total
}

func (dao *DepartmentDao) GetDept(id int) (models.Department, error) {
	o := GetOrmer()
	dept := models.Department{
		Id: id,
	}
	err := o.Read(&dept)

	return dept, err
}

func (dao *DepartmentDao) GetDeptByName(name string) (dept models.Department, err error) {
	o := GetOrmer()
	err = o.QueryTable(&models.Department{}).Filter("name", name).One(&dept)

	return dept, err
}

func (dao *DepartmentDao) GetDeptByOtherName(id int, name string) (models.Department, error) {
	o := GetOrmer()
	var department models.Department
	err := o.Raw("select * from department where name = ? and id <> ?", name, id).QueryRow(&department)

	return department, err
}

func (dao *DepartmentDao) Insert(dto *dto.DepartmentAddDto) (int64, error) {
	o := GetOrmer()
	var dept models.Department
	dept.Name = dto.Name
	dept.ParentId = dto.ParentId
	dept.OrderNum = dto.OrderNum

	return o.Insert(&dept)
}

func (dao *DepartmentDao) Update(dto *dto.DepartmentEditDto) (int64, error) {
	o := GetOrmer()
	dept := models.Department{Id: dto.Id}
	if o.Read(&dept) == nil {
		dept.Name = dto.Name
		dept.ParentId = dto.ParentId
		dept.OrderNum = dto.OrderNum

		return o.Update(&dept, "Name", "ParentId", "OrderNum")
	}

	return 0, errors.New("update failed")
}

//删除
func (dao *DepartmentDao) Delete(id int) error {
	o := GetOrmer()
	dept := &models.Department{Id: id}
	if o.Read(dept) == nil {
		_, err := o.Delete(dept)
		if err != nil {
			return err
		}
	}

	return nil
}
