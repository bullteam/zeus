package models

const TableDataPerm = "data_perm"

type DataPerm struct {
	Id        int     `json:"id"`        // 自增ID
	Name      string  `json:"name"`      // 名称
	Perms     string  `json:"perms"`     // 数据权限标识
	OrderNum  int     `json:"order_num"` // 排序字段
	Menu      *Menu   `orm:"rel(one)" json:"menu"`
	Domain    *Domain `orm:"rel(one)" json:"domain"`
	PermsRule string  `json:"perms_rule"` // 数据规则
}

type DataPermQuery struct {
	DomainId   int
	Name       string
	Pagination *Pagination
}

func (dp *DataPerm) TableName() string {
	return TableDataPerm
}
