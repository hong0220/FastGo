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

type ISysUserRole interface {
	Creat(ctx context.Context, req *api.SysUserRoleCreatReq) (err error)
	Delete(ctx context.Context, id int) (err error)
	DeleteByIds(ctx context.Context, ids []int) (err error)
	Update(ctx context.Context, req *api.SysUserRoleUpdateReq) (err error)
	GetId(ctx context.Context, id int) (res *entity.SysUserRole, err error)
	GetList(ctx context.Context, page *base.Page) (res []*entity.SysUserRole, err error)
}

var localSysUserRole ISysUserRole

func SysUserRole() ISysUserRole {
	if localSysUserRole == nil {
		panic("implement not found for interface ISysUserRole, forgot register?")
	}
	return localSysUserRole
}

func RegisterSysUserRole(i ISysUserRole) {
	localSysUserRole = i
}
