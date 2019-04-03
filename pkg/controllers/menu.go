package controllers

import (
	"zeus/pkg/models"
	"zeus/pkg/components"
)


type MenuController struct {
	TokenCheckController
}

func (c *MenuController) List()  {
	domain_id, err := c.GetInt("domain_id")
	if err != nil {
		domain_id = 1
	}
	menu := models.Menu_list(domain_id)
	c.Resp(0,"sucess",map[string]interface{}{
        "result" : menu,
	})
}

func (c *MenuController) Add()  {
	form := models.MenuaddForm{}
	if err := c.ParseForm(&form); err != nil {
		c.Fail(components.ErrInputData)
		return
	}
	Menu, err := models.NewMenu(&form)
	if err != nil {
		c.Fail(components.ErrInputData)
		return
	}
	Menu.Insert()
	c.Resp(0,"success",map[string]interface{}{

	})
}

func (c *MenuController) Edit()  {
	id,err := c.GetInt("id")
	if err != nil {
		c.Fail(components.ErrIdData)
		return
	}
	parent_id,err := c.GetInt("parent_id")
	if err != nil {
		c.Fail(components.ErrIdData)
		return
	}
	domain_id,err := c.GetInt("domain_id")
	if err != nil {
		c.Fail(components.ErrIdData)
		return
	}
	name := c.Input().Get("name")
	url := c.Input().Get("url")
	perms := c.Input().Get("perms")
	menu_type,err := c.GetInt("menu_type")
	if err != nil {
		c.Fail(components.ErrIdData)
		return
	}
	icon := c.Input().Get("icon")
	order_num,err := c.GetInt("order_num")
	if err != nil {
		c.Fail(components.ErrIdData)
		return
	}
	err = models.UpdateMenu(id,parent_id,domain_id,name,url,perms,menu_type,icon,order_num)
	if err != nil {
		c.Fail(components.ErrInputData)
		return
	}
	c.Resp(0,"success",map[string]interface{}{})
}

/**删除菜单**/
func (c *MenuController) Del() {
	id,err := c.GetInt("id")
	if err != nil {
		c.Fail(components.ErrInputData)
		return
	}
	err = models.DeleteMenu(id)
	if err != nil {
		c.Fail(components.ErrInputData)
		return
	}
	c.Resp(0, "success", map[string]interface{}{})
}