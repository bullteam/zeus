package models

//for列表,有关联关系
type Role struct {
	Id         int     `json:"id"`
	Name       string  `json:"name"`
	Domain     *Domain `json:"domain" orm:"rel(one)"`
	RoleName   string  `json:"role_name"`
	Remark     string  `json:"remark"`
	Users      []*User `json:"users" orm:"reverse(many)"`
	MenuIds    string  `json:"menu_ids"`
	MenuIdsEle string  `json:"menu_ids_ele"`
}

//for更新创建
type RoleEntity struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	DomainId   int    `json:"domain_id"`
	RoleName   string `json:"role_name"`
	Remark     string `json:"remark"`
	MenuIds    string `json:"menu_ids"`
	MenuIdsEle string `json:"menu_ids_ele"`
}
