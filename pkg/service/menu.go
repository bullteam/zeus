package service

import (
	"github.com/bullteam/zeus/pkg/models"
	"github.com/astaxie/beego/orm"
)

type MenuService struct{

}

func (ms MenuService) GetMenusByIds(ids string) []orm.Params{
	menu := models.Menu{}
	return menu.GetMenusByIds(ids)
}
