package dto

type DataPermAddDto struct {
	DomainId int    `form:"domain_id" valid:"Required"`          // 项目域id
	MenuId   int    `form:"menu_id" valid:"Required"`            // 菜单ID
	Name     string `form:"name" valid:"Required;MaxSize(50)"`   // 名称
	Perms    string `form:"perms" valid:"Required;MaxSize(100)"` // 数据权限标识
	OrderNum int    `form:"order_num" valid:"Required"`          // 排序字段
}

type DataPermEditDto struct {
	Id       int    `form:"id" valid:"Required"`
	DomainId int    `form:"domain_id" valid:"Required"`          // 项目域id
	MenuId   int    `form:"menu_id" valid:"Required"`            // 菜单ID
	Name     string `form:"name" valid:"Required;MaxSize(50)"`   // 名称
	Perms    string `form:"perms" valid:"Required;MaxSize(100)"` // 数据权限标识
	OrderNum int    `form:"order_num" valid:"Required"`          // 排序字段
}
