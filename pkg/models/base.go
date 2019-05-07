package models

import "github.com/astaxie/beego/orm"

type Pagination struct {
	Start int
	Limit int
}

func init() {
	orm.RegisterModel(
		new(DataPerm),
		new(Department),
		new(Domain),
		new(Menu),
		new(Role),
		new(RoleDataPerm),
		new(User),
		new(UserRole),
		new(UserOAuth),
	)
}
