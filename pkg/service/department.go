package service

import (
	"fmt"
	"github.com/bullteam/zeus/pkg/dao"
	"github.com/bullteam/zeus/pkg/dto"
	"github.com/bullteam/zeus/pkg/models"
)

type DepartmentService struct {
	dao *dao.DepartmentDao
}

func (s *DepartmentService) GetList(start int, limit int, q []string) ([]*models.Department, int64) {
	return s.dao.List(start, limit, q)
}

func (s *DepartmentService) GetDept(id int) (models.Department, error) {
	return s.dao.GetDept(id)
}

func (s *DepartmentService) CheckIfDeptExists(name string) bool {
	_, err := s.dao.GetDeptByName(name)

	return err == nil
}

func (s *DepartmentService) CheckIfOtherDeptExists(id int, name string) bool {
	_, err := s.dao.GetDeptByOtherName(id, name)

	return err == nil
}

//did is department id
func (s *DepartmentService) CheckIfPeopleInside(did int) bool {
	user := &UserService{}
	_, _c := user.GetList(0, 1, []string{fmt.Sprintf("d=%d", did)})
	return _c > 0
}

func (s *DepartmentService) Insert(dto *dto.DepartmentAddDto) (int64, error) {
	return s.dao.Insert(dto)
}

func (s *DepartmentService) Update(dto *dto.DepartmentEditDto) (int64, error) {
	return s.dao.Update(dto)
}

func (s *DepartmentService) Delete(id int) error {
	return s.dao.Delete(id)
}
