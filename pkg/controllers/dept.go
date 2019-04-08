package controllers

import (
	"github.com/bullteam/zeus/pkg/components"
	"github.com/bullteam/zeus/pkg/models"
	"github.com/bullteam/zeus/pkg/service"
	"github.com/bullteam/zeus/pkg/dto"
	"strings"
)

type DeptController struct {
	TokenCheckController
}

func (d *DeptController) List(){
	ds := service.DepartmentService{}
	start, _ := d.GetInt("start", 0)
	limit, _ := d.GetInt("limit", LIST_ROWS_PERPAGE)
	q := d.GetString("q")
	data,c := ds.GetList(start,limit,strings.Split(q,","))
	d.Resp(0, "success", map[string]interface{}{
		"result": data,
		"total":  c,
	})
}
//func (c *DeptController) List() {
//	page, page_err := c.GetInt("p")
//	if page_err != nil {
//		page = 1
//	}
//	offset, offset_err := c.GetInt("offset")
//	if offset_err != nil {
//		offset = 20
//	}
//
//	deptlist, cnt := models.Dept_list(page, offset)
//	c.Resp(0, "success", map[string]interface{}{
//		"deptlist": deptlist,
//		"total":      cnt,
//		"page":       page,
//	})
//}

func (c *DeptController) Post() {
	//form := models.DeptaddForm{}
	//if err := c.ParseForm(&form); err != nil {
	//	c.Fail(components.ErrInputData)
	//	return
	//}
	deptDto := &dto.DepartmentAddDto{}
	c.ParseAndValidate(deptDto)
	ds := service.DepartmentService{}

	if ds.CheckIfDeptExists(deptDto.Name){
		c.Fail(components.ErrDupRecord,"部门已存在")
		return
	}
	id,err := ds.Create(deptDto)
	if err != nil{
		c.Fail(components.ErrAddFail,err.Error())
	}
	//Dept, err := models.NewDept(&form)
	//if err != nil {
	//	c.Fail(components.ErrInputData)
	//	return
	//}
	//Dept.Insert()
	c.Resp(0, "success", map[string]interface{}{
		"id" : id,
	})
}

func (c *DeptController) Show() {
	//deptDto := &dto.DepartmentEditDto{}
	id, err := c.GetInt("id")
	if err != nil {
		c.Fail(components.ErrIdData)
		return
	}
	domain, err := models.GetDept(id)
	if err != nil {
		c.Fail(components.ErrInputData)
		return
	}
	c.Resp(0, "success", map[string]interface{}{
		"domain": domain,
	})
}

func (c *DeptController) Edit() {
	//id, err := c.GetInt("id")
	//if err != nil {
	//	c.Fail(components.ErrIdData)
	//	return
	//}
	//name := c.Input().Get("name")
	//parent_id,err := c.GetInt("parent_id")
	//if err != nil {
	//	c.Fail(components.ErrIdData)
	//	return
	//}
	//err = models.UpdateDept(id, name,parent_id)
	//if err != nil {
	//	c.Fail(components.ErrInputData)
	//	return
	//}
	//c.Resp(0, "success", map[string]interface{}{})
	deptDto := &dto.DepartmentEditDto{}
	c.ParseAndValidate(deptDto)
	ds := service.DepartmentService{}

	if ds.CheckIfOtherDeptExists(deptDto.Id,deptDto.Name){
		c.Fail(components.ErrDupRecord,"部门已存在")
		return
	}
	_,err := ds.Update(deptDto)
	if err != nil{
		c.Fail(components.ErrEditFail,err.Error())
	}
	//Dept, err := models.NewDept(&form)
	//if err != nil {
	//	c.Fail(components.ErrInputData)
	//	return
	//}
	//Dept.Insert()
	c.Resp(0, "success")
}

func (c *DeptController) Del() {
	id, err := c.GetInt("id")
	if err != nil {
		c.Fail(components.ErrInputData)
		return
	}
	deptService := service.DepartmentService{}
	if deptService.CheckIfPeopleInside(id){
		c.Fail(components.ErrDeptDel,components.ErrDeptDel.Moreinfo)
		return
	}
	err = models.DeleteDept(id)
	if err != nil {
		c.Fail(components.ErrInputData)
		return
	}
	c.Resp(0, "success", map[string]interface{}{})
}

func (c *DeptController) CheckNoMember(){
	id, err := c.GetInt("id")
	if err != nil {
		c.Fail(components.ErrInputData)
		return
	}
	deptService := service.DepartmentService{}
	if deptService.CheckIfPeopleInside(id){
		c.Fail(components.ErrDeptHasMember,components.ErrDeptHasMember.Moreinfo)
		return
	}
	c.Resp(0,"success")
}