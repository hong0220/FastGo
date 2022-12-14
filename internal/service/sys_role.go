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

type ISysRole interface {
	Creat(ctx context.Context, req *api.SysRoleCreatReq) (err error)
	Delete(ctx context.Context, id int) (err error)
	DeleteByIds(ctx context.Context, ids []int) (err error)
	Update(ctx context.Context, req *api.SysRoleUpdateReq) (err error)
	GetId(ctx context.Context, id int) (res *entity.SysRole, err error)
	GetList(ctx context.Context, page *base.Page) (res []*entity.SysRole, err error)
}

var localSysRole ISysRole

func SysRole() ISysRole {
	if localSysRole == nil {
		panic("implement not found for interface ISysRole, forgot register?")
	}
	return localSysRole
}

func RegisterSysRole(i ISysRole) {
	localSysRole = i
}
