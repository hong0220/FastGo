package controller

import (
	"context"
	"github.com/hong0220/FastGo/internal/api"
	"github.com/hong0220/FastGo/internal/model/entity"
	"github.com/hong0220/FastGo/internal/service"
)

type SysDict struct{}

// 包装对外暴露对象
var CSysDict SysDict

// Creat 创建记录
func (c SysDict) Creat(ctx context.Context, req *api.SysDictCreatReq) (res *api.SysDictCreatRes, err error) {
	err = service.SysDict().Creat(ctx, req)
	if err != nil {
		return nil, err
	}
	res = &api.SysDictCreatRes{
		IsSuccess: true,
	}
	return
}

// Delete 根据id删除单条记录
func (c SysDict) Delete(ctx context.Context, req *api.SysDictDeleteReq) (res *api.SysDictDeleteRes, err error) {
	err = service.SysDict().Delete(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	res = &api.SysDictDeleteRes{
		IsSuccess: true,
	}
	return
}

// DeleteByIds 根据id数组批量删除记录
func (c SysDict) DeleteByIds(ctx context.Context, req *api.SysDictDeleteByIdsReq) (res *api.SysDictDeleteByIdsRes, err error) {
	err = service.SysDict().DeleteByIds(ctx, req.Ids)
	if err != nil {
		return nil, err
	}
	res = &api.SysDictDeleteByIdsRes{
		IsSuccess: true,
	}
	return
}

// Update 修改记录
func (c SysDict) Update(ctx context.Context, req *api.SysDictUpdateReq) (res *api.SysDictUpdateRes, err error) {
	err = service.SysDict().Update(ctx, req)
	if err != nil {
		return nil, err
	}
	res = &api.SysDictUpdateRes{
		IsSuccess: true,
	}
	return
}

// GetId 根据id查询记录
func (c SysDict) GetId(ctx context.Context, req *api.SysDictIdReq) (res *api.SysDictIdRes, err error) {
	var instance *entity.SysDict
	instance, err = service.SysDict().GetId(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	res = &api.SysDictIdRes{
		Instance: instance,
	}
	return
}

// GetList 分页查询
func (c SysDict) GetList(ctx context.Context, req *api.SysDictListReq) (res *api.SysDictListRes, err error) {
	list, err := service.SysDict().GetList(ctx, req.Page)
	if err != nil {
		return nil, err
	}
	res = &api.SysDictListRes{
		List: list,
	}
	return
}
