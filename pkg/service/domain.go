package service

import (
	"github.com/bullteam/zeus/pkg/dao"
	"github.com/bullteam/zeus/pkg/dto"
	"github.com/bullteam/zeus/pkg/models"
)

type DomainService struct {
	dao *dao.DomainDao
}

func (s *DomainService) NewDomain(dto *dto.DomainAddDto) (d *models.Domain, err error) {
	return s.dao.NewDomain(dto)
}

func (s *DomainService) GetList(start int, limit int, q []string) ([]*models.Domain, int64) {
	return s.dao.List(start, limit, q)
}

func (s *DomainService) Insert(dto *dto.DomainAddDto) (int64, error) {
	return s.dao.Insert(dto)
}

func (s *DomainService) Update(id int, name string, callbackurl string, remark string) error {
	return s.dao.Update(id, name, callbackurl, remark)
}

func (s *DomainService) Delete(id int) error {
	return s.dao.Delete(id)
}

func (s *DomainService) GetDomain(id int) (models.Domain, error) {
	return s.dao.GetDomain(id)
}

func (s *DomainService) GetDomainByCode(code string) (domain models.Domain, err error) {
	return s.dao.GetDomainByCode(code)
}
