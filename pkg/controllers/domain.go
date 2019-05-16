package controllers

import (
	"strings"
	"zeus/pkg/components"
	"zeus/pkg/dto"
	"zeus/pkg/service"
)

type DomainController struct {
	TokenCheckController
}

func (c *DomainController) List() {
	ds := service.DomainService{}
	start, _ := c.GetInt("start", 0)
	limit, _ := c.GetInt("limit", listRowsPerPage)
	q := c.GetString("q")
	data, _c := ds.GetList(start, limit, strings.Split(q, ","))
	c.Resp(0, "success", map[string]interface{}{
		"result": data,
		"total":  _c,
	})
}

func (c *DomainController) Post() {
	domainAddDto := &dto.DomainAddDto{}
	err := c.ParseAndValidateFirstErr(domainAddDto)
	if err != nil {
		c.Fail(components.ErrInvalidParams, err.Error())
		return
	}
	ds := service.DomainService{}
	lastInsertId, err := ds.Insert(domainAddDto)
	if err != nil {
		c.Fail(components.ErrAddFail, err.Error())
		return
	}
	c.Resp(0, "success", map[string]interface{}{
		"id": lastInsertId,
	})
}

func (c *DomainController) Show() {
	id, err := c.GetInt("id")
	if err != nil {
		c.Fail(components.ErrIdData)
		return
	}
	ds := service.DomainService{}
	domain, err := ds.GetDomain(id)
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
	ds := service.DomainService{}
	err = ds.Update(id, name, callbackurl, remark)
	if err != nil {
		c.Fail(components.ErrInputData)
		return
	}
	c.Resp(0, "success", map[string]interface{}{})
}

func (c *DomainController) Del() {
	id, err := c.GetInt("id")
	if err != nil {
		c.Fail(components.ErrInputData)
		return
	}
	ds := service.DomainService{}
	err = ds.Delete(id)
	if err != nil {
		c.Fail(components.ErrInputData)
		return
	}
	c.Resp(0, "success", map[string]interface{}{})
}
