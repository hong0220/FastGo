package controller

import (
	"context"
	"github.com/hong0220/FastGo/internal/api"
	"github.com/hong0220/FastGo/internal/model/entity"
	"github.com/hong0220/FastGo/internal/service"
)

type SysRoleOffice struct{}

// 包装对外暴露对象
var CSysRoleOffice SysRoleOffice

// Creat 创建记录
func (c SysRoleOffice) Creat(ctx context.Context, req *api.SysRoleOfficeCreatReq) (res *api.SysRoleOfficeCreatRes, err error) {
	err = service.SysRoleOffice().Creat(ctx, req)
	if err != nil {
		return nil, err
	}
	res = &api.SysRoleOfficeCreatRes{
		IsSuccess: true,
	}
	return
}

// Delete 根据id删除单条记录
func (c SysRoleOffice) Delete(ctx context.Context, req *api.SysRoleOfficeDeleteReq) (res *api.SysRoleOfficeDeleteRes, err error) {
	err = service.SysRoleOffice().Delete(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	res = &api.SysRoleOfficeDeleteRes{
		IsSuccess: true,
	}
	return
}

// DeleteByIds 根据id数组批量删除记录
func (c SysRoleOffice) DeleteByIds(ctx context.Context, req *api.SysRoleOfficeDeleteByIdsReq) (res *api.SysRoleOfficeDeleteByIdsRes, err error) {
	err = service.SysRoleOffice().DeleteByIds(ctx, req.Ids)
	if err != nil {
		return nil, err
	}
	res = &api.SysRoleOfficeDeleteByIdsRes{
		IsSuccess: true,
	}
	return
}

// Update 修改记录
func (c SysRoleOffice) Update(ctx context.Context, req *api.SysRoleOfficeUpdateReq) (res *api.SysRoleOfficeUpdateRes, err error) {
	err = service.SysRoleOffice().Update(ctx, req)
	if err != nil {
		return nil, err
	}
	res = &api.SysRoleOfficeUpdateRes{
		IsSuccess: true,
	}
	return
}

// GetId 根据id查询记录
func (c SysRoleOffice) GetId(ctx context.Context, req *api.SysRoleOfficeIdReq) (res *api.SysRoleOfficeIdRes, err error) {
	var instance *entity.SysRoleOffice
	instance, err = service.SysRoleOffice().GetId(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	res = &api.SysRoleOfficeIdRes{
		Instance: instance,
	}
	return
}

// GetList 分页查询
func (c SysRoleOffice) GetList(ctx context.Context, req *api.SysRoleOfficeListReq) (res *api.SysRoleOfficeListRes, err error) {
	list, err := service.SysRoleOffice().GetList(ctx, req.Page)
	if err != nil {
		return nil, err
	}
	res = &api.SysRoleOfficeListRes{
		List: list,
	}
	return
}
