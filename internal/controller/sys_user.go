package controller

import (
	"context"
	"github.com/hong0220/FastGo/internal/api"
	"github.com/hong0220/FastGo/internal/model/entity"
	"github.com/hong0220/FastGo/internal/service"
)

type SysUser struct{}

// 包装对外暴露对象
var CSysUser SysUser

// Creat 创建记录
func (c SysUser) Creat(ctx context.Context, req *api.SysUserCreatReq) (res *api.SysUserCreatRes, err error) {
	err = service.SysUser().Creat(ctx, req)
	if err != nil {
		return nil, err
	}
	res = &api.SysUserCreatRes{
		IsSuccess: true,
	}
	return
}

// Delete 根据id删除单条记录
func (c SysUser) Delete(ctx context.Context, req *api.SysUserDeleteReq) (res *api.SysUserDeleteRes, err error) {
	err = service.SysUser().Delete(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	res = &api.SysUserDeleteRes{
		IsSuccess: true,
	}
	return
}

// DeleteByIds 根据id数组批量删除记录
func (c SysUser) DeleteByIds(ctx context.Context, req *api.SysUserDeleteByIdsReq) (res *api.SysUserDeleteByIdsRes, err error) {
	err = service.SysUser().DeleteByIds(ctx, req.Ids)
	if err != nil {
		return nil, err
	}
	res = &api.SysUserDeleteByIdsRes{
		IsSuccess: true,
	}
	return
}

// Update 修改记录
func (c SysUser) Update(ctx context.Context, req *api.SysUserUpdateReq) (res *api.SysUserUpdateRes, err error) {
	err = service.SysUser().Update(ctx, req)
	if err != nil {
		return nil, err
	}
	res = &api.SysUserUpdateRes{
		IsSuccess: true,
	}
	return
}

// GetId 根据id查询记录
func (c SysUser) GetId(ctx context.Context, req *api.SysUserIdReq) (res *api.SysUserIdRes, err error) {
	var instance *entity.SysUser
	instance, err = service.SysUser().GetId(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	res = &api.SysUserIdRes{
		Instance: instance,
	}
	return
}

// GetList 分页查询
func (c SysUser) GetList(ctx context.Context, req *api.SysUserListReq) (res *api.SysUserListRes, err error) {
	list, err := service.SysUser().GetList(ctx, req.Page)
	if err != nil {
		return nil, err
	}
	res = &api.SysUserListRes{
		List: list,
	}
	return
}
