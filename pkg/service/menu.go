package service

import (
	"github.com/astaxie/beego/orm"
	"zeus/pkg/dao"
	"zeus/pkg/dto"
)

type MenuService struct {
	dao *dao.MenuDao
}

func (ms *MenuService) List(domainId int) (menus []orm.Params) {
	return ms.dao.List(domainId)
}

func (ms MenuService) GetMenusByIds(ids string) []orm.Params {
	return ms.dao.GetMenusByIds(ids)
}

func (ms *MenuService) Insert(dto *dto.MenuAddDto) (int64, error) {
	return ms.dao.Insert(dto)
}

func (ms *MenuService) Update(dto *dto.MenuEditDto) error {
	return ms.dao.Update(dto)
}

func (ms *MenuService) Delete(id int) error {
	return ms.dao.Delete(id)
}
