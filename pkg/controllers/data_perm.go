package controllers

import (
	"zeus/pkg/components"
	"zeus/pkg/dto"
	"zeus/pkg/models"
	"zeus/pkg/service"
)

type DataPermController struct {
	TokenCheckController
}

// 数据权限列表
func (c *DataPermController) List() {
	ds := service.DataPermService{}
	start, limit := c.GetPaginationParams()
	domainId, _ := c.GetInt("domain_id", 0)
	name := c.GetString("name")
	query := &models.DataPermQuery{
		DomainId: domainId,
		Name:     name,
		Pagination: &models.Pagination{
			Start: start,
			Limit: limit,
		},
	}
	data, total := ds.GetDataPermList(query)
	c.Resp(0, "success", map[string]interface{}{
		"result": data,
		"total":  total,
	})
}

// 数据权限详情
func (c *DataPermController) Show() {
	dataPermId, err := c.GetInt("id")
	if err != nil {
		c.Fail(components.ErrInputData)
		return
	}
	ds := service.DataPermService{}
	data, err := ds.GetById(dataPermId)
	if err != nil {
		c.Fail(components.ErrIdData)
		return
	}
	c.Resp(0, "success", data)
}

// 添加数据权限
func (c *DataPermController) Add() {
	dataPermAddDto := &dto.DataPermAddDto{}
	err := c.ParseAndValidateFirstErr(dataPermAddDto)
	if err != nil {
		c.Fail(components.ErrInvalidParams, err.Error())
		return
	}
	ds := service.DataPermService{}
	lastInsertId, err := ds.Insert(dataPermAddDto)
	if err != nil {
		c.Fail(components.ErrAddFail, err.Error())
		return
	}
	c.Resp(0, "success", map[string]interface{}{
		"id": lastInsertId,
	})
}

// 修改数据权限
func (c *DataPermController) Edit() {
	dataPermEditDto := &dto.DataPermEditDto{}
	err := c.ParseAndValidateFirstErr(dataPermEditDto)
	if err != nil {
		c.Fail(components.ErrInvalidParams, err.Error())
		return
	}
	ds := service.DataPermService{}

	err = ds.Update(dataPermEditDto)
	if err != nil {
		c.Fail(components.ErrEditFail, err.Error())
		return
	}
	c.Resp(0, "success", map[string]interface{}{})
}

// 删除数据权限
func (c *DataPermController) Del() {
	id, _ := c.GetInt("id", 0)
	if id > 0 {
		ds := service.DataPermService{}
		err := ds.Delete(id)
		if err != nil {
			c.Fail(components.ErrDelFail, err.Error())
			return
		}
		c.Resp(0, "success", map[string]interface{}{})
		return
	}
	c.Fail(components.ErrDelFail, "delete failed")
}
