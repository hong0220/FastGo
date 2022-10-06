package controller

import (
	"context"
	"github.com/hong0220/FastGo/internal/api"
	"github.com/hong0220/FastGo/internal/model/entity"
	"github.com/hong0220/FastGo/internal/service"
)

type SysLog struct{}

// 包装对外暴露对象
var CSysLog SysLog

// Creat 创建记录
func (c SysLog) Creat(ctx context.Context, req *api.SysLogCreatReq) (res *api.SysLogCreatRes, err error) {
	err = service.SysLog().Creat(ctx, req)
	if err != nil {
		return nil, err
	}
	res = &api.SysLogCreatRes{
		IsSuccess: true,
	}
	return
}

// Delete 根据id删除单条记录
func (c SysLog) Delete(ctx context.Context, req *api.SysLogDeleteReq) (res *api.SysLogDeleteRes, err error) {
	err = service.SysLog().Delete(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	res = &api.SysLogDeleteRes{
		IsSuccess: true,
	}
	return
}

// DeleteByIds 根据id数组批量删除记录
func (c SysLog) DeleteByIds(ctx context.Context, req *api.SysLogDeleteByIdsReq) (res *api.SysLogDeleteByIdsRes, err error) {
	err = service.SysLog().DeleteByIds(ctx, req.Ids)
	if err != nil {
		return nil, err
	}
	res = &api.SysLogDeleteByIdsRes{
		IsSuccess: true,
	}
	return
}

// Update 修改记录
func (c SysLog) Update(ctx context.Context, req *api.SysLogUpdateReq) (res *api.SysLogUpdateRes, err error) {
	err = service.SysLog().Update(ctx, req)
	if err != nil {
		return nil, err
	}
	res = &api.SysLogUpdateRes{
		IsSuccess: true,
	}
	return
}

// GetId 根据id查询记录
func (c SysLog) GetId(ctx context.Context, req *api.SysLogIdReq) (res *api.SysLogIdRes, err error) {
	var instance *entity.SysLog
	instance, err = service.SysLog().GetId(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	res = &api.SysLogIdRes{
		Instance: instance,
	}
	return
}

// GetList 分页查询
func (c SysLog) GetList(ctx context.Context, req *api.SysLogListReq) (res *api.SysLogListRes, err error) {
	list, err := service.SysLog().GetList(ctx, req.Page)
	if err != nil {
		return nil, err
	}
	res = &api.SysLogListRes{
		List: list,
	}
	return
}
