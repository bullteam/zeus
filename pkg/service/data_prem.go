package service

import (
	"github.com/bullteam/zeus/pkg/dao"
	"github.com/bullteam/zeus/pkg/dto"
	"github.com/bullteam/zeus/pkg/models"
)

type DataPermService struct {
	dao *dao.DataPermDao
}

func (s *DataPermService) GetDataPermList(query *models.DataPermQuery) ([]*models.DataPerm, int64) {
	return s.dao.GetDataPermList(query)
}

func (s *DataPermService) Insert(dto *dto.DataPermAddDto) (int64, error) {
	return s.dao.Insert(dto)
}

func (s *DataPermService) Update(dto *dto.DataPermEditDto) error {
	return s.dao.Update(dto)
}

func (s *DataPermService) Delete(id int) error {
	return s.dao.Delete(id)
}
