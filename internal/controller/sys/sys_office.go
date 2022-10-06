package controller

import (
	"context"
	"github.com/hong0220/FastGo/internal/api"
	"github.com/hong0220/FastGo/internal/model/entity"
	"github.com/hong0220/FastGo/internal/service"
)

type SysOffice struct{}

// 包装对外暴露对象
var CSysOffice SysOffice

// Creat 创建记录
func (c SysOffice) Creat(ctx context.Context, req *api.SysOfficeCreatReq) (res *api.SysOfficeCreatRes, err error) {
	err = service.SysOffice().Creat(ctx, req)
	if err != nil {
		return nil, err
	}
	res = &api.SysOfficeCreatRes{
		IsSuccess: true,
	}
	return
}

// Delete 根据id删除单条记录
func (c SysOffice) Delete(ctx context.Context, req *api.SysOfficeDeleteReq) (res *api.SysOfficeDeleteRes, err error) {
	err = service.SysOffice().Delete(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	res = &api.SysOfficeDeleteRes{
		IsSuccess: true,
	}
	return
}

// DeleteByIds 根据id数组批量删除记录
func (c SysOffice) DeleteByIds(ctx context.Context, req *api.SysOfficeDeleteByIdsReq) (res *api.SysOfficeDeleteByIdsRes, err error) {
	err = service.SysOffice().DeleteByIds(ctx, req.Ids)
	if err != nil {
		return nil, err
	}
	res = &api.SysOfficeDeleteByIdsRes{
		IsSuccess: true,
	}
	return
}

// Update 修改记录
func (c SysOffice) Update(ctx context.Context, req *api.SysOfficeUpdateReq) (res *api.SysOfficeUpdateRes, err error) {
	err = service.SysOffice().Update(ctx, req)
	if err != nil {
		return nil, err
	}
	res = &api.SysOfficeUpdateRes{
		IsSuccess: true,
	}
	return
}

// GetId 根据id查询记录
func (c SysOffice) GetId(ctx context.Context, req *api.SysOfficeIdReq) (res *api.SysOfficeIdRes, err error) {
	var instance *entity.SysOffice
	instance, err = service.SysOffice().GetId(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	res = &api.SysOfficeIdRes{
		Instance: instance,
	}
	return
}

// GetList 分页查询
func (c SysOffice) GetList(ctx context.Context, req *api.SysOfficeListReq) (res *api.SysOfficeListRes, err error) {
	list, err := service.SysOffice().GetList(ctx, req.Page)
	if err != nil {
		return nil, err
	}
	res = &api.SysOfficeListRes{
		List: list,
	}
	return
}
