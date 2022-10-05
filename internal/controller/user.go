package controller

import (
	"context"
	"github.com/hong0220/FastGo/internal/api"
	"github.com/hong0220/FastGo/internal/model/entity"
	"github.com/hong0220/FastGo/internal/service"
)

type User struct{}

// 包装对外暴露对象
var cUser User

// Creat 创建记录
func (c User) Creat(ctx context.Context, req *api.UserCreatReq) (res *api.UserCreatRes, err error) {
	err = service.User().Creat(ctx, req)
	if err != nil {
		return nil, err
	}
	res = &api.UserCreatRes{
		IsSuccess: true,
	}
	return
}

// Delete 根据id删除单条记录
func (c User) Delete(ctx context.Context, req *api.UserDeleteReq) (res *api.UserDeleteRes, err error) {
	err = service.User().Delete(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	res = &api.UserDeleteRes{
		IsSuccess: true,
	}
	return
}

// DeleteByIds 根据id数组批量删除记录
func (c User) DeleteByIds(ctx context.Context, req *api.UserDeleteByIdsReq) (res *api.UserDeleteByIdsRes, err error) {
	err = service.User().DeleteByIds(ctx, req.Ids)
	if err != nil {
		return nil, err
	}
	res = &api.UserDeleteByIdsRes{
		IsSuccess: true,
	}
	return
}

// Update 修改记录
func (c User) Update(ctx context.Context, req *api.UserUpdateReq) (res *api.UserUpdateRes, err error) {
	err = service.User().Update(ctx, req)
	if err != nil {
		return nil, err
	}
	res = &api.UserUpdateRes{
		IsSuccess: true,
	}
	return
}

// GetId 根据id查询记录
func (c User) GetId(ctx context.Context, req *api.UserIdReq) (res *api.UserIdRes, err error) {
	var instance *entity.User
	instance, err = service.User().GetId(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	res = &api.UserIdRes{
		Instance: instance,
	}
	return
}

// GetList 分页查询
func (c User) GetList(ctx context.Context, req *api.UserListReq) (res *api.UserListRes, err error) {
	list, err := service.User().GetList(ctx, req.Page)
	if err != nil {
		return nil, err
	}
	res = &api.UserListRes{
		List: list,
	}
	return
}
