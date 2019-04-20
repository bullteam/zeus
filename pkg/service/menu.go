package service

import (
	"github.com/astaxie/beego/orm"
	"github.com/bullteam/zeus/pkg/models"
)

type MenuService struct {
}

func (ms MenuService) GetMenusByIds(ids string) []orm.Params {
	menu := models.Menu{}
	return menu.GetMenusByIds(ids)
}
