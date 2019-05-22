package dto

type DataPermAddDto struct {
	DomainId  int    `form:"domain_id" valid:"Required"`        // 项目域id
	ParentId  int    `form:"parent_id" valid:"Min(0)"`          // 父级id
	Name      string `form:"name" valid:"Required;MaxSize(50)"` // 名称
	Perms     string `form:"perms" valid:"MaxSize(100)"`        // 数据权限key
	PermsRule string `form:"perms_rule"`                        // 数据权限规则
	PermsType int    `form:"perms_type" valid:"Required"`       // 类型 1=分类 2=数据权限
	OrderNum  int    `form:"order_num" valid:"Required"`        // 排序字段
	Remarks   string `form:"remarks" valid:"MaxSize(500)"`      // 说明
}

type DataPermEditDto struct {
	Id        int    `form:"id" valid:"Required"`
	DomainId  int    `form:"domain_id" valid:"Required"`        // 项目域id
	ParentId  int    `form:"parent_id" valid:"Min(0)"`          // 菜单ID
	Name      string `form:"name" valid:"Required;MaxSize(50)"` // 名称
	Perms     string `form:"perms" valid:"MaxSize(100)"`        // 数据权限key
	PermsRule string `form:"perms_rule" valid:"Required"`       // 数据权限规则
	PermsType int    `form:"perms_type" valid:"Required"`       // 类型 1=分类 2=数据权限
	OrderNum  int    `form:"order_num" valid:"Required"`        // 排序字段
	Remarks   string `form:"remarks" valid:"MaxSize(500)"`      // 说明
}
