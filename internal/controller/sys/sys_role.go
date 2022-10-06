package controller

import (
	"context"
	"github.com/hong0220/FastGo/internal/api"
	"github.com/hong0220/FastGo/internal/model/entity"
	"github.com/hong0220/FastGo/internal/service"
)

type SysRole struct{}

// 包装对外暴露对象
var CSysRole SysRole

// Creat 创建记录
func (c SysRole) Creat(ctx context.Context, req *api.SysRoleCreatReq) (res *api.SysRoleCreatRes, err error) {
	err = service.SysRole().Creat(ctx, req)
	if err != nil {
		return nil, err
	}
	res = &api.SysRoleCreatRes{
		IsSuccess: true,
	}
	return
}

// Delete 根据id删除单条记录
func (c SysRole) Delete(ctx context.Context, req *api.SysRoleDeleteReq) (res *api.SysRoleDeleteRes, err error) {
	err = service.SysRole().Delete(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	res = &api.SysRoleDeleteRes{
		IsSuccess: true,
	}
	return
}

// DeleteByIds 根据id数组批量删除记录
func (c SysRole) DeleteByIds(ctx context.Context, req *api.SysRoleDeleteByIdsReq) (res *api.SysRoleDeleteByIdsRes, err error) {
	err = service.SysRole().DeleteByIds(ctx, req.Ids)
	if err != nil {
		return nil, err
	}
	res = &api.SysRoleDeleteByIdsRes{
		IsSuccess: true,
	}
	return
}

// Update 修改记录
func (c SysRole) Update(ctx context.Context, req *api.SysRoleUpdateReq) (res *api.SysRoleUpdateRes, err error) {
	err = service.SysRole().Update(ctx, req)
	if err != nil {
		return nil, err
	}
	res = &api.SysRoleUpdateRes{
		IsSuccess: true,
	}
	return
}

// GetId 根据id查询记录
func (c SysRole) GetId(ctx context.Context, req *api.SysRoleIdReq) (res *api.SysRoleIdRes, err error) {
	var instance *entity.SysRole
	instance, err = service.SysRole().GetId(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	res = &api.SysRoleIdRes{
		Instance: instance,
	}
	return
}

// GetList 分页查询
func (c SysRole) GetList(ctx context.Context, req *api.SysRoleListReq) (res *api.SysRoleListRes, err error) {
	list, err := service.SysRole().GetList(ctx, req.Page)
	if err != nil {
		return nil, err
	}
	res = &api.SysRoleListRes{
		List: list,
	}
	return
}
