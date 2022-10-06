// ================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"

	"github.com/hong0220/FastGo/internal/api"
	"github.com/hong0220/FastGo/internal/model/entity"
	"github.com/hong0220/FastGo/pkg/common/base"
)

type ISysUser interface {
	Creat(ctx context.Context, req *api.SysUserCreatReq) (err error)
	Delete(ctx context.Context, id int) (err error)
	DeleteByIds(ctx context.Context, ids []int) (err error)
	Update(ctx context.Context, req *api.SysUserUpdateReq) (err error)
	GetId(ctx context.Context, id int) (res *entity.SysUser, err error)
	GetList(ctx context.Context, page *base.Page) (res []*entity.SysUser, err error)
}

var localSysUser ISysUser

func SysUser() ISysUser {
	if localSysUser == nil {
		panic("implement not found for interface ISysUser, forgot register?")
	}
	return localSysUser
}

func RegisterSysUser(i ISysUser) {
	localSysUser = i
}