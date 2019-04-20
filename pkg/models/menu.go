package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"strings"
	"time"
)

type MenuaddForm struct {
	Parent_id int    `form:"parent_id"`
	Domain_id int    `form:"domain_id"`
	Name      string `form:"name"`
	Url       string `form:"url"`
	Perms     string `form:"perms"`
	Menu_type int    `form:"menu_type"`
	Icon      string `form:"icon"`
	Order_num int    `form:"order_num"`
}

type Menu struct {
	Id               int
	Parent_id        int
	Domain_id        int
	Name             string
	Url              string
	Perms            string
	Menu_type        int
	Icon             string
	Order_num        int
	Create_time      string
	Last_update_time time.Time
}

func init() {
	orm.RegisterModel(new(Menu))
}

func NewMenu(menuaddform *MenuaddForm) (m *Menu, err error) {
	menu := Menu{
		Parent_id: menuaddform.Parent_id,
		Domain_id: menuaddform.Domain_id,
		Name:      menuaddform.Name,
		Url:       menuaddform.Url,
		Perms:     menuaddform.Perms,
		Menu_type: menuaddform.Menu_type,
		Icon:      menuaddform.Icon,
		Order_num: menuaddform.Order_num,
	}
	return &menu, nil
}
func (m *Menu) Insert() {
	o := orm.NewOrm()
	o.Using("default")
	_, err := o.Raw("insert into menu(parent_id,domain_id,name,url,perms,menu_type,icon,order_num) values (?,?,?,?,?,?,?,?)", m.Parent_id, m.Domain_id, m.Name, m.Url, m.Perms, m.Menu_type, m.Icon, m.Order_num).Exec()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("insert ok")
	}
}

func Menu_list(domain_id int) (menus []orm.Params) {
	var menu []orm.Params
	o := orm.NewOrm()
	o.Using("default")
	o.Raw("select id,parent_id,name,url,perms,menu_type,icon,order_num from menu where domain_id = ?", domain_id).Values(&menu)
	return menu
}

//修改
func UpdateMenu(id int, parent_id int, domain_id int, name string, url string, perms string, menu_type int, icon string, order_num int) error {
	o := orm.NewOrm()
	menu := &Menu{Id: id}
	if o.Read(menu) == nil {
		menu.Parent_id = parent_id
		menu.Domain_id = domain_id
		menu.Name = name
		menu.Url = url
		menu.Perms = perms
		menu.Menu_type = menu_type
		menu.Icon = icon
		menu.Order_num = order_num
		menu.Last_update_time = time.Now()
		_, err := o.Update(menu)
		if err != nil {
			return err
		}
	}
	return nil
}

//删除
func DeleteMenu(id int) error {
	o := orm.NewOrm()
	menu := &Menu{Id: id}
	if o.Read(menu) == nil {
		_, err := o.Delete(menu)
		if err != nil {
			return err
		}
	}
	return nil
}

func (m Menu) GetMenusByIds(ids string) []orm.Params {
	var menus []orm.Params
	o := orm.NewOrm()
	fid := strings.Split(ids, ",")
	binds := strings.Repeat("?,", len(fid))
	prepare := fmt.Sprintf(`select * from menu where id in (%s) and menu_type=? order by order_num asc`, strings.Trim(binds, ","))
	o.Raw(prepare, fid, 1).Values(&menus)
	return menus
}
