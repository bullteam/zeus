package controllers

import (
	//"fmt"
	"github.com/bullteam/zeus/pkg/components"
	"github.com/bullteam/zeus/pkg/dto"
	"github.com/bullteam/zeus/pkg/service"
	"strings"
)

type RoleController struct {
	TokenCheckController
}

//角色列表
func (r *RoleController) List() {
	rs := service.RoleService{}
	start, _ := r.GetInt("start", 0)
	limit, _ := r.GetInt("limit", LIST_ROWS_PERPAGE)
	q := r.GetString("q")
	data, c := rs.GetList(start, limit, strings.Split(q, ","))
	r.Resp(0, "success", map[string]interface{}{
		"result": data,
		"total":  c,
	})
}

//角色详情
func (r *RoleController) Show() {
	rs := service.RoleService{}
	role_id, err := r.GetInt("id")
	if err != nil {
		r.Fail(components.ErrInputData)
		return
	}
	//get data
	data, perms, data_err := rs.GetRoleById(role_id)

	//err data
	if data_err != nil {
		r.Fail(components.ErrIdData)
		//err_obj := components.ErrIdData
		//err_obj.Moreinfo = fmt.Sprintf("%v", data_err)
		//r.Fail(err_obj)
		return
	}
	r.Resp(0, "success", map[string]interface{}{
		"detail": data,
		"perms":  perms,
	})
}

//角色添加
func (r *RoleController) Add() {
	rs := service.RoleService{}
	roleDto := &dto.RoleDto{}
	r.ParseAndValidate(roleDto)
	result, err := rs.CreateRole(roleDto)
	if err != nil {
		r.Fail(components.ErrDupRecord, "角色已存在")
		return
	}
	menu_ids := r.GetString("menu_ids")
	if menu_ids != "" {
		id, _ := result.LastInsertId()
		if err := rs.AssignPerm(roleDto.DomainId, int(id), menu_ids); err != nil {
			r.Fail(components.ErrRoleAssignFail, err.Error())
			return
		}
	}
	r.Resp(0, "success")
}

//角色更新
func (r *RoleController) Edit() {
	rs := service.RoleService{}
	roleDto := &dto.RoleDto{}
	r.ParseAndValidate(roleDto)
	if roleDto.Id <= 0 {
		r.Fail(components.ErrIdData)
	}
	err := rs.UpdateRole(roleDto)
	if err != nil {
		r.Fail(components.ErrEditFail, err.Error())
		return
	}
	menu_ids := r.GetString("menu_ids")
	if menu_ids != "" {
		if err := rs.AssignPerm(roleDto.DomainId, roleDto.Id, menu_ids); err != nil {
			r.Fail(components.ErrRoleAssignFail, err.Error())
			return
		}
	}
	r.Resp(0, "success")
	//r.Resp(0, "success")
}

//角色删除
func (r *RoleController) Del() {
	id, err := r.GetInt("id")
	if err != nil || id < 1 {
		r.Fail(components.ErrIdData)
	}
	rs := service.RoleService{}
	err = rs.DeleteRole(id)
	if err != nil {
		r.Fail(components.ErrDelFail, err.Error())
		return
	}

	r.Resp(0, "success")
}

//角色分配
func (r *RoleController) Assign() {
	domain_id, err := r.GetInt("domain_id")
	if err != nil || domain_id < 1 {
		r.Fail(components.ErrIdData)
	}
	role_id, err := r.GetInt("role_id")
	if err != nil || role_id < 1 {
		r.Fail(components.ErrIdData)
	}
	menu_ids := r.GetString("menu_ids")
	if menu_ids == "" {
		r.Fail(components.ErrIdData)
	}
	us := service.RoleService{}
	if err := us.AssignPerm(domain_id, role_id, menu_ids); err != nil {
		r.Fail(components.ErrRoleAssignFail, err.Error())
		return
	}
	r.Resp(0, "success")
}
