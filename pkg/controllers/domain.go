package controllers

import (
	"github.com/bullteam/zeus/pkg/components"
	"github.com/bullteam/zeus/pkg/models"
	"github.com/bullteam/zeus/pkg/service"
	"strings"
)

type DomainController struct {
	TokenCheckController
}

//func (c *DomainController) List() {
//	page, page_err := c.GetInt("p")
//	if page_err != nil {
//		page = 1
//	}
//	offset, offset_err := c.GetInt("offset")
//	if offset_err != nil {
//		offset = 20
//	}
//
//	domainlist, cnt := models.Domain_list(page, offset)
//	c.Resp(0, "success", map[string]interface{}{
//		"result": domainlist,
//		"total":      cnt,
//		"page":       page,
//	})
//}
func (c *DomainController) List() {
	ds := service.DomainService{}
	start, _ := c.GetInt("start", 0)
	limit, _ := c.GetInt("limit", LIST_ROWS_PERPAGE)
	q := c.GetString("q")
	data,_c := ds.GetList(start,limit,strings.Split(q,","))
	c.Resp(0, "success", map[string]interface{}{
		"result": data,
		"total":  _c,
	})
}

func (c *DomainController) Post() {
	form := models.DomainaddForm{}
	if err := c.ParseForm(&form); err != nil {
		c.Fail(components.ErrInputData)
		return
	}
	Domain, err := models.NewDomain(&form)
	if err != nil {
		c.Fail(components.ErrInputData)
		return
	}
	Domain.Insert()
	c.Resp(0, "success", map[string]interface{}{})
}

func (c *DomainController) Show() {
	id, err := c.GetInt("id")
	if err != nil {
		c.Fail(components.ErrIdData)
		return
	}
	domain, err := models.GetDomain(id)
	if err != nil {
		c.Fail(components.ErrInputData)
		return
	}
	c.Resp(0, "success", map[string]interface{}{
		"domain": domain,
	})
}

func (c *DomainController) Edit() {
	id, err := c.GetInt("id")
	name := c.Input().Get("name")
	callbackurl := c.Input().Get("callbackurl")
	remark := c.Input().Get("remark")
	if err != nil {
		c.Fail(components.ErrIdData)
		return
	}
	err = models.UpdateDomain(id, name, callbackurl, remark)
	if err != nil {
		c.Fail(components.ErrInputData)
		return
	}
	c.Resp(0, "success", map[string]interface{}{})
}

/**删除域**/
func (c *DomainController) Del() {
	id, err := c.GetInt("id")
	if err != nil {
		c.Fail(components.ErrInputData)
		return
	}
	err = models.DeleteDomain(id)
	if err != nil {
		c.Fail(components.ErrInputData)
		return
	}
	c.Resp(0, "success", map[string]interface{}{})
}
