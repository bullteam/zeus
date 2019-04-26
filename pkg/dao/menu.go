package dao

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/bullteam/zeus/pkg/dto"
	"github.com/bullteam/zeus/pkg/models"
	"strings"
)

type MenuDao struct {
}

func (dao *MenuDao) NewMenu(dto *dto.MenuAddDto) (m *models.Menu, err error) {
	menu := models.Menu{
		ParentId: dto.ParentId,
		DomainId: dto.DomainId,
		Name:     dto.Name,
		Url:      dto.Url,
		Perms:    dto.Perms,
		MenuType: dto.MenuType,
		Icon:     dto.Icon,
		OrderNum: dto.OrderNum,
	}
	return &menu, nil
}

func (dao *MenuDao) List(domainId int) (menus []orm.Params) {
	var menu []orm.Params
	o := GetOrmer()
	sql := "select id,parent_id,name,url,perms,menu_type,icon,order_num from menu where domain_id = ?"
	_, err := o.Raw(sql, domainId).Values(&menu)
	if err != nil {
		return menus
	}

	return menu
}

func (dao *MenuDao) GetMenusByIds(ids string) []orm.Params {
	var menus []orm.Params
	o := GetOrmer()
	fid := strings.Split(ids, ",")
	binds := strings.Repeat("?,", len(fid))
	prepare := fmt.Sprintf(`select * from menu where id in (%s) and menu_type=? order by order_num asc`, strings.Trim(binds, ","))
	_, err := o.Raw(prepare, fid, 1).Values(&menus)
	if err != nil {
		return menus
	}

	return menus
}

func (dao *MenuDao) Insert(dto *dto.MenuAddDto) (int64, error) {
	o := GetOrmer()
	var menu models.Menu
	menu.ParentId = dto.ParentId
	menu.DomainId = dto.DomainId
	menu.Name = dto.Name
	menu.Url = dto.Url
	menu.Perms = dto.Perms
	menu.MenuType = dto.MenuType
	menu.Icon = dto.Icon
	menu.OrderNum = dto.OrderNum

	return o.Insert(&menu)
}

func (dao *MenuDao) Update(dto *dto.MenuEditDto) error {
	o := GetOrmer()
	menu := models.Menu{Id: dto.Id}
	if o.Read(&menu) == nil {
		menu.ParentId = dto.ParentId
		menu.DomainId = dto.DomainId
		menu.Name = dto.Name
		menu.Url = dto.Url
		menu.Perms = dto.Perms
		menu.MenuType = dto.MenuType
		menu.Icon = dto.Icon
		menu.OrderNum = dto.OrderNum
		_, err := o.Update(&menu)
		if err != nil {
			return err
		}
	}

	return nil
}

func (dao *MenuDao) Delete(id int) error {
	o := GetOrmer()
	menu := &models.Menu{Id: id}
	if o.Read(menu) == nil {
		_, err := o.Delete(menu)
		if err != nil {
			return err
		}
	}

	return nil
}
