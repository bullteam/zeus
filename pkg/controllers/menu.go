package controllers

import (
	"zeus/pkg/components"
	"zeus/pkg/dto"
	"zeus/pkg/service"
)

type MenuController struct {
	TokenCheckController
}

func (c *MenuController) List() {
	domainId, err := c.GetInt("domain_id")
	if err != nil {
		domainId = 1
	}
	s := service.MenuService{}
	menu := s.List(domainId)
	c.Resp(0, "success", map[string]interface{}{
		"result": menu,
	})
}

func (c *MenuController) Add() {
	menuAddDto := &dto.MenuAddDto{}
	err := c.ParseAndValidateFirstErr(menuAddDto)
	if err != nil {
		c.Fail(components.ErrInvalidParams, err.Error())
		return
	}
	s := service.MenuService{}
	lastInsertId, err := s.Insert(menuAddDto)
	if err != nil {
		c.Fail(components.ErrAddFail)
		return
	}
	c.Resp(0, "success", map[string]interface{}{
		"id": lastInsertId,
	})
}

func (c *MenuController) Edit() {
	menuEditDto := &dto.MenuEditDto{}
	err := c.ParseAndValidateFirstErr(menuEditDto)
	if err != nil {
		c.Fail(components.ErrInvalidParams, err.Error())
		return
	}
	s := service.MenuService{}
	err = s.Update(menuEditDto)
	if err != nil {
		c.Fail(components.ErrEditFail)
		return
	}

	c.Resp(0, "success", map[string]interface{}{})
}

func (c *MenuController) Del() {
	id, err := c.GetInt("id")
	if err != nil {
		c.Fail(components.ErrInputData)
		return
	}
	s := service.MenuService{}
	err = s.Delete(id)
	if err != nil {
		c.Fail(components.ErrInputData)
		return
	}
	c.Resp(0, "success", map[string]interface{}{})
}
