package controller

import (
	"context"
	"github.com/hong0220/FastGo/internal/api"
	"github.com/hong0220/FastGo/internal/model/entity"
	"github.com/hong0220/FastGo/internal/service"
)

type SysTask struct{}

// 包装对外暴露对象
var CSysTask SysTask

// Creat 创建记录
func (c SysTask) Creat(ctx context.Context, req *api.SysTaskCreatReq) (res *api.SysTaskCreatRes, err error) {
	err = service.SysTask().Creat(ctx, req)
	if err != nil {
		return nil, err
	}
	res = &api.SysTaskCreatRes{
		IsSuccess: true,
	}
	return
}

// Delete 根据id删除单条记录
func (c SysTask) Delete(ctx context.Context, req *api.SysTaskDeleteReq) (res *api.SysTaskDeleteRes, err error) {
	err = service.SysTask().Delete(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	res = &api.SysTaskDeleteRes{
		IsSuccess: true,
	}
	return
}

// DeleteByIds 根据id数组批量删除记录
func (c SysTask) DeleteByIds(ctx context.Context, req *api.SysTaskDeleteByIdsReq) (res *api.SysTaskDeleteByIdsRes, err error) {
	err = service.SysTask().DeleteByIds(ctx, req.Ids)
	if err != nil {
		return nil, err
	}
	res = &api.SysTaskDeleteByIdsRes{
		IsSuccess: true,
	}
	return
}

// Update 修改记录
func (c SysTask) Update(ctx context.Context, req *api.SysTaskUpdateReq) (res *api.SysTaskUpdateRes, err error) {
	err = service.SysTask().Update(ctx, req)
	if err != nil {
		return nil, err
	}
	res = &api.SysTaskUpdateRes{
		IsSuccess: true,
	}
	return
}

// GetId 根据id查询记录
func (c SysTask) GetId(ctx context.Context, req *api.SysTaskIdReq) (res *api.SysTaskIdRes, err error) {
	var instance *entity.SysTask
	instance, err = service.SysTask().GetId(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	res = &api.SysTaskIdRes{
		Instance: instance,
	}
	return
}

// GetList 分页查询
func (c SysTask) GetList(ctx context.Context, req *api.SysTaskListReq) (res *api.SysTaskListRes, err error) {
	list, err := service.SysTask().GetList(ctx, req.Page)
	if err != nil {
		return nil, err
	}
	res = &api.SysTaskListRes{
		List: list,
	}
	return
}
