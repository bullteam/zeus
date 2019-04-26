package controllers

import (
	"github.com/bullteam/zeus/pkg/components"
	"github.com/bullteam/zeus/pkg/dto"
	"github.com/bullteam/zeus/pkg/service"
	"strings"
)

type DeptController struct {
	TokenCheckController
}

func (d *DeptController) List() {
	ds := service.DepartmentService{}
	start, _ := d.GetInt("start", 0)
	limit, _ := d.GetInt("limit", listRowsPerPage)
	q := d.GetString("q")
	data, c := ds.GetList(start, limit, strings.Split(q, ","))
	d.Resp(0, "success", map[string]interface{}{
		"result": data,
		"total":  c,
	})
}

func (c *DeptController) Post() {
	deptDto := &dto.DepartmentAddDto{}
	err := c.ParseAndValidateFirstErr(deptDto)
	if err != nil {
		c.Fail(components.ErrInvalidParams, err.Error())
		return
	}
	ds := service.DepartmentService{}

	if ds.CheckIfDeptExists(deptDto.Name) {
		c.Fail(components.ErrDupRecord, "部门已存在")
		return
	}
	id, err := ds.Insert(deptDto)
	if err != nil {
		c.Fail(components.ErrAddFail, err.Error())
	}

	c.Resp(0, "success", map[string]interface{}{
		"id": id,
	})
}

func (c *DeptController) Show() {
	id, err := c.GetInt("id")
	if err != nil {
		c.Fail(components.ErrIdData)
		return
	}
	ds := service.DepartmentService{}
	domain, err := ds.GetDept(id)
	if err != nil {
		c.Fail(components.ErrInputData)
		return
	}
	c.Resp(0, "success", map[string]interface{}{
		"domain": domain,
	})
}

func (c *DeptController) Edit() {
	deptDto := &dto.DepartmentEditDto{}
	err := c.ParseAndValidateFirstErr(deptDto)
	if err != nil {
		c.Fail(components.ErrInvalidParams, err.Error())
		return
	}
	ds := service.DepartmentService{}

	if ds.CheckIfOtherDeptExists(deptDto.Id, deptDto.Name) {
		c.Fail(components.ErrDupRecord, "部门已存在")
		return
	}
	_, err = ds.Update(deptDto)
	if err != nil {
		c.Fail(components.ErrEditFail, err.Error())
	}
	c.Resp(0, "success")
}

func (c *DeptController) Del() {
	id, err := c.GetInt("id")
	if err != nil {
		c.Fail(components.ErrInputData)
		return
	}
	ds := service.DepartmentService{}
	if ds.CheckIfPeopleInside(id) {
		c.Fail(components.ErrDeptDel, components.ErrDeptDel.Moreinfo)
		return
	}
	err = ds.Delete(id)
	if err != nil {
		c.Fail(components.ErrInputData)
		return
	}
	c.Resp(0, "success", map[string]interface{}{})
}

func (c *DeptController) CheckNoMember() {
	id, err := c.GetInt("id")
	if err != nil {
		c.Fail(components.ErrInputData)
		return
	}
	ds := service.DepartmentService{}
	if ds.CheckIfPeopleInside(id) {
		c.Fail(components.ErrDeptHasMember, components.ErrDeptHasMember.Moreinfo)
		return
	}
	c.Resp(0, "success")
}
