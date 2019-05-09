package service

import (
	"github.com/astaxie/beego/orm"
	"github.com/bullteam/zeus/pkg/dao"
	"github.com/bullteam/zeus/pkg/dto"
	"github.com/bullteam/zeus/pkg/models"
)

type DataPermService struct {
	dao *dao.DataPermDao
}

func (s *DataPermService) GetDataPermList(query *models.DataPermQuery) (dataPerms []orm.Params, total int64) {
	dpRows, total := s.dao.GetDataPermList(query)

	if total > 0 {
		for _, dpRow := range dpRows {
			dataPerms = append(dataPerms, map[string]interface{}{
				"id": dpRow.Id,
				"name": dpRow.Name,
				"perms": dpRow.Perms,
				"perms_rule": dpRow.PermsRule,
				"perms_type": dpRow.PermsType,
				"order_num": dpRow.OrderNum,
				"parent_id": dpRow.ParentId,
				"domain_id": dpRow.Domain.Id,
				"domain_name": dpRow.Domain.Name,
				"domain_code": dpRow.Domain.Code,
				"remarks": dpRow.Remarks,
			})
		}
	}

	return dataPerms, total
}

func (s *DataPermService) GetById(dataPermId int) (models.DataPerm, error) {
	return s.dao.GetById(dataPermId)
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
