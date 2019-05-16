package controllers

import (
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
	limit, _ := r.GetInt("limit", listRowsPerPage)
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
	roleId, err := r.GetInt("id")
	if err != nil {
		r.Fail(components.ErrInputData)
		return
	}
	//get data
	data, perms, err := rs.GetRoleById(roleId)

	// 获取数据权限列表
	dataPerms, _ := rs.GetRoleDataPermsByRoleId(roleId)

	//err data
	if err != nil {
		r.Fail(components.ErrIdData)
		return
	}
	r.Resp(0, "success", map[string]interface{}{
		"detail":     data,
		"perms":      perms,
		"data_perms": dataPerms,
	})
}

//角色添加
func (r *RoleController) Add() {
	rs := service.RoleService{}
	roleDto := &dto.RoleDto{}
	err := r.ParseAndValidateFirstErr(roleDto)
	if err != nil {
		r.Fail(components.ErrInvalidParams, err.Error())
		return
	}
	result, err := rs.CreateRole(roleDto)
	if err != nil {
		r.Fail(components.ErrDupRecord, "角色已存在")
		return
	}
	roleId, _ := result.LastInsertId()
	// 分配功能权限
	menuIds := r.GetString("menu_ids")
	if menuIds != "" {
		if err := rs.AssignPerm(roleDto.DomainId, int(roleId), menuIds); err != nil {
			r.Fail(components.ErrRoleAssignFail, err.Error())
			return
		}
	}
	// 分配数据权限
	dataPermIds := r.GetString("data_perm_ids")
	if len(dataPermIds) > 0 {
		_ = rs.AssignDataPerm(int(roleId), dataPermIds)
	}
	r.Resp(0, "success")
}

//角色更新
func (r *RoleController) Edit() {
	rs := service.RoleService{}
	roleDto := &dto.RoleDto{}
	err := r.ParseAndValidateFirstErr(roleDto)
	if roleDto.Id <= 0 {
		r.Fail(components.ErrIdData)
		return
	}
	if err != nil {
		r.Fail(components.ErrInvalidParams, err.Error())
		return
	}
	err = rs.UpdateRole(roleDto)
	if err != nil {
		r.Fail(components.ErrEditFail, err.Error())
		return
	}
	// 分配功能权限
	menuIds := r.GetString("menu_ids")
	if menuIds != "" {
		if err := rs.AssignPerm(roleDto.DomainId, roleDto.Id, menuIds); err != nil {
			r.Fail(components.ErrRoleAssignFail, err.Error())
			return
		}
	}
	// 分配数据权限
	dataPermIds := r.GetString("data_perm_ids")
	if len(dataPermIds) > 0 {
		_ = rs.AssignDataPerm(roleDto.Id, dataPermIds)
	}
	r.Resp(0, "success")
}

//角色删除
func (r *RoleController) Del() {
	id, err := r.GetInt("id")
	if err != nil || id < 1 {
		r.Fail(components.ErrIdData)
		return
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
	domainId, err := r.GetInt("domain_id")
	if err != nil || domainId < 1 {
		r.Fail(components.ErrIdData)
		return
	}
	roleId, err := r.GetInt("role_id")
	if err != nil || roleId < 1 {
		r.Fail(components.ErrIdData)
		return
	}
	menuIds := r.GetString("menu_ids")
	if menuIds == "" {
		r.Fail(components.ErrIdData)
		return
	}
	us := service.RoleService{}
	if err := us.AssignPerm(domainId, roleId, menuIds); err != nil {
		r.Fail(components.ErrRoleAssignFail, err.Error())
		return
	}
	r.Resp(0, "success")
}
