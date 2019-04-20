package service

import "github.com/bullteam/zeus/pkg/models"

type DomainService struct {
}

//GetList
func (r DomainService) GetList(start int, limit int, q []string) ([]*models.Domain, int64) {
	ds := models.Domain{}
	return ds.List(start, limit, q)
}
