package controller

import (
	"context"
	"github.com/hong0220/FastGo/internal/api"
	"github.com/hong0220/FastGo/internal/model/entity"
	"github.com/hong0220/FastGo/internal/service"
)

type SysResource struct{}

// 包装对外暴露对象
var CSysResource SysResource

// Creat 创建记录
func (c SysResource) Creat(ctx context.Context, req *api.SysResourceCreatReq) (res *api.SysResourceCreatRes, err error) {
	err = service.SysResource().Creat(ctx, req)
	if err != nil {
		return nil, err
	}
	res = &api.SysResourceCreatRes{
		IsSuccess: true,
	}
	return
}

// Delete 根据id删除单条记录
func (c SysResource) Delete(ctx context.Context, req *api.SysResourceDeleteReq) (res *api.SysResourceDeleteRes, err error) {
	err = service.SysResource().Delete(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	res = &api.SysResourceDeleteRes{
		IsSuccess: true,
	}
	return
}

// DeleteByIds 根据id数组批量删除记录
func (c SysResource) DeleteByIds(ctx context.Context, req *api.SysResourceDeleteByIdsReq) (res *api.SysResourceDeleteByIdsRes, err error) {
	err = service.SysResource().DeleteByIds(ctx, req.Ids)
	if err != nil {
		return nil, err
	}
	res = &api.SysResourceDeleteByIdsRes{
		IsSuccess: true,
	}
	return
}

// Update 修改记录
func (c SysResource) Update(ctx context.Context, req *api.SysResourceUpdateReq) (res *api.SysResourceUpdateRes, err error) {
	err = service.SysResource().Update(ctx, req)
	if err != nil {
		return nil, err
	}
	res = &api.SysResourceUpdateRes{
		IsSuccess: true,
	}
	return
}

// GetId 根据id查询记录
func (c SysResource) GetId(ctx context.Context, req *api.SysResourceIdReq) (res *api.SysResourceIdRes, err error) {
	var instance *entity.SysResource
	instance, err = service.SysResource().GetId(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	res = &api.SysResourceIdRes{
		Instance: instance,
	}
	return
}

// GetList 分页查询
func (c SysResource) GetList(ctx context.Context, req *api.SysResourceListReq) (res *api.SysResourceListRes, err error) {
	list, err := service.SysResource().GetList(ctx, req.Page)
	if err != nil {
		return nil, err
	}
	res = &api.SysResourceListRes{
		List: list,
	}
	return
}
