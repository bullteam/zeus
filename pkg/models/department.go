package models

const TableDepartment = "department"

type Department struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	OrderNum int    `json:"order_num"`
	ParentId int    `json:"parent_id"`
	//CreateTime  time.Time `json:"create_time"`
	//UpdateTime time.Time `json:"update_time"`
}

func (dp *Department) TableName() string {
	return TableDepartment
}
