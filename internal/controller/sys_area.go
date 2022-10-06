package controller

import (
	"context"
	"github.com/hong0220/FastGo/internal/api"
	"github.com/hong0220/FastGo/internal/model/entity"
	"github.com/hong0220/FastGo/internal/service"
)

type SysArea struct{}

// 包装对外暴露对象
var CSysArea SysArea

// Creat 创建记录
func (c SysArea) Creat(ctx context.Context, req *api.SysAreaCreatReq) (res *api.SysAreaCreatRes, err error) {
	err = service.SysArea().Creat(ctx, req)
	if err != nil {
		return nil, err
	}
	res = &api.SysAreaCreatRes{
		IsSuccess: true,
	}
	return
}

// Delete 根据id删除单条记录
func (c SysArea) Delete(ctx context.Context, req *api.SysAreaDeleteReq) (res *api.SysAreaDeleteRes, err error) {
	err = service.SysArea().Delete(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	res = &api.SysAreaDeleteRes{
		IsSuccess: true,
	}
	return
}

// DeleteByIds 根据id数组批量删除记录
func (c SysArea) DeleteByIds(ctx context.Context, req *api.SysAreaDeleteByIdsReq) (res *api.SysAreaDeleteByIdsRes, err error) {
	err = service.SysArea().DeleteByIds(ctx, req.Ids)
	if err != nil {
		return nil, err
	}
	res = &api.SysAreaDeleteByIdsRes{
		IsSuccess: true,
	}
	return
}

// Update 修改记录
func (c SysArea) Update(ctx context.Context, req *api.SysAreaUpdateReq) (res *api.SysAreaUpdateRes, err error) {
	err = service.SysArea().Update(ctx, req)
	if err != nil {
		return nil, err
	}
	res = &api.SysAreaUpdateRes{
		IsSuccess: true,
	}
	return
}

// GetId 根据id查询记录
func (c SysArea) GetId(ctx context.Context, req *api.SysAreaIdReq) (res *api.SysAreaIdRes, err error) {
	var instance *entity.SysArea
	instance, err = service.SysArea().GetId(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	res = &api.SysAreaIdRes{
		Instance: instance,
	}
	return
}

// GetList 分页查询
func (c SysArea) GetList(ctx context.Context, req *api.SysAreaListReq) (res *api.SysAreaListRes, err error) {
	list, err := service.SysArea().GetList(ctx, req.Page)
	if err != nil {
		return nil, err
	}
	res = &api.SysAreaListRes{
		List: list,
	}
	return
}
