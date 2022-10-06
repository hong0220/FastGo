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

type IUser interface {
	Creat(ctx context.Context, req *api.UserCreatReq) (err error)
	Delete(ctx context.Context, id int) (err error)
	DeleteByIds(ctx context.Context, ids []int) (err error)
	Update(ctx context.Context, req *api.UserUpdateReq) (err error)
	GetId(ctx context.Context, id int) (res *entity.User, err error)
	GetList(ctx context.Context, page *base.Page) (res []*entity.User, err error)
}

var localUser IUser

func User() IUser {
	if localUser == nil {
		panic("implement not found for interface IUser, forgot register?")
	}
	return localUser
}

func RegisterUser(i IUser) {
	localUser = i
}