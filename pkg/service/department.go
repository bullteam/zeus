package service

import (
	"github.com/bullteam/zeus/pkg/models"
	"fmt"
	"github.com/bullteam/zeus/pkg/dto"
)

type DepartmentService struct {

}
//GetList
func (r DepartmentService) GetList(start int, limit int, q []string) ([]*models.Department, int64) {
	ds := models.Department{}
	return ds.List(start, limit, q)
}
func (d DepartmentService) CheckIfDeptExists(name string)  bool{
	m := models.Department{}
	_,err := m.GetDeptByName(name)
	return err == nil
}
func (d DepartmentService) CheckIfOtherDeptExists(id int,name string)  bool{
	m := models.Department{}
	_,err := m.GetDeptByOtherName(id,name)
	return err == nil
}
//did is department id
func (d DepartmentService) CheckIfPeopleInside(did int) bool{
	user := &models.User{}
	_,_c := user.List(0,1,[]string{fmt.Sprintf("d=%d",did)})
	return _c > 0
}

func (d DepartmentService) Create(dto *dto.DepartmentAddDto)(int64,error){
	dept := &models.Department{
		Name : dto.Name,
		Parent_id : dto.Parent_id,
		Order_num: dto.Order_num,
	}
	return dept.Insert()
}

func (d DepartmentService) Update(dto *dto.DepartmentEditDto)(int64,error){
	dept := &models.Department{
		Id : dto.Id,
		Name : dto.Name,
		Order_num: dto.Order_num,
	}
	return dept.Update()
}