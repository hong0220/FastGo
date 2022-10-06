package controller

import (
	"context"
	"github.com/hong0220/FastGo/internal/api"
	"github.com/hong0220/FastGo/internal/model/entity"
	"github.com/hong0220/FastGo/internal/service"
)

type SysUserRole struct{}

// 包装对外暴露对象
var CSysUserRole SysUserRole

// Creat 创建记录
func (c SysUserRole) Creat(ctx context.Context, req *api.SysUserRoleCreatReq) (res *api.SysUserRoleCreatRes, err error) {
	err = service.SysUserRole().Creat(ctx, req)
	if err != nil {
		return nil, err
	}
	res = &api.SysUserRoleCreatRes{
		IsSuccess: true,
	}
	return
}

// Delete 根据id删除单条记录
func (c SysUserRole) Delete(ctx context.Context, req *api.SysUserRoleDeleteReq) (res *api.SysUserRoleDeleteRes, err error) {
	err = service.SysUserRole().Delete(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	res = &api.SysUserRoleDeleteRes{
		IsSuccess: true,
	}
	return
}

// DeleteByIds 根据id数组批量删除记录
func (c SysUserRole) DeleteByIds(ctx context.Context, req *api.SysUserRoleDeleteByIdsReq) (res *api.SysUserRoleDeleteByIdsRes, err error) {
	err = service.SysUserRole().DeleteByIds(ctx, req.Ids)
	if err != nil {
		return nil, err
	}
	res = &api.SysUserRoleDeleteByIdsRes{
		IsSuccess: true,
	}
	return
}

// Update 修改记录
func (c SysUserRole) Update(ctx context.Context, req *api.SysUserRoleUpdateReq) (res *api.SysUserRoleUpdateRes, err error) {
	err = service.SysUserRole().Update(ctx, req)
	if err != nil {
		return nil, err
	}
	res = &api.SysUserRoleUpdateRes{
		IsSuccess: true,
	}
	return
}

// GetId 根据id查询记录
func (c SysUserRole) GetId(ctx context.Context, req *api.SysUserRoleIdReq) (res *api.SysUserRoleIdRes, err error) {
	var instance *entity.SysUserRole
	instance, err = service.SysUserRole().GetId(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	res = &api.SysUserRoleIdRes{
		Instance: instance,
	}
	return
}

// GetList 分页查询
func (c SysUserRole) GetList(ctx context.Context, req *api.SysUserRoleListReq) (res *api.SysUserRoleListRes, err error) {
	list, err := service.SysUserRole().GetList(ctx, req.Page)
	if err != nil {
		return nil, err
	}
	res = &api.SysUserRoleListRes{
		List: list,
	}
	return
}
