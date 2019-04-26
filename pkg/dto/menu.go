package dto

type MenuAddDto struct {
	ParentId int    `form:"parent_id"`
	DomainId int    `form:"domain_id" valid:"Required"`
	Name     string `form:"name" valid:"Required"`
	Url      string `form:"url"`
	Perms    string `form:"perms"`
	MenuType int    `form:"menu_type" valid:"Required"`
	Icon     string `form:"icon"`
	OrderNum int    `form:"order_num" valid:"Required"`
}

type MenuEditDto struct {
	Id       int    `form:"id" valid:"Required"`
	ParentId int    `form:"parent_id"`
	DomainId int    `form:"domain_id" valid:"Required"`
	Name     string `form:"name" valid:"Required"`
	Url      string `form:"url"`
	Perms    string `form:"perms"`
	MenuType int    `form:"menu_type" valid:"Required"`
	Icon     string `form:"icon"`
	OrderNum int    `form:"order_num" valid:"Required"`
}
