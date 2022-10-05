package api

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/hong0220/FastGo/internal/model/entity"
	"github.com/hong0220/FastGo/pkg/common/base"
)

// 创建记录
type UserCreatReq struct {
	g.Meta `path:"/user" method:"post" summary:"新建" tags:"User"`
	Form   *entity.User `json:"form" in:"query"  dc:"Create对象"`
}
type UserCreatRes struct {
	g.Meta    `mime:"json" example:"obj"`
	IsSuccess bool `json:"isSuccess" dc:"是否成功"`
}

// Delete 根据id删除单条记录
type UserDeleteReq struct {
	g.Meta `path:"/user" method:"delete" summary:"删除" tags:"User"`
	Id     int `json:"id" in:"query" dc:"主键ID"`
}
type UserDeleteRes struct {
	g.Meta    `mime:"json" example:"obj"`
	IsSuccess bool `json:"isSuccess" dc:"是否成功"`
}

// DeleteByIds 根据id数组批量删除记录
type UserDeleteByIdsReq struct {
	g.Meta `path:"/user/deleteByIds" method:"delete" summary:"批量删除" tags:"User"`
	Ids    []int `json:"ids" in:"query" dc:"主键ID数组"`
}
type UserDeleteByIdsRes struct {
	g.Meta    `mime:"json" example:"obj"`
	IsSuccess bool `json:"isSuccess" dc:"是否成功"`
}

// Update 修改记录
type UserUpdateReq struct {
	g.Meta `path:"/user" method:"put" summary:"修改" tags:"User"`
	Form   *entity.User `json:"form" in:"query" dc:"Update对象"`
}
type UserUpdateRes struct {
	g.Meta    `mime:"json" example:"obj"`
	IsSuccess bool `json:"isSuccess" dc:"是否成功"`
}

// GetId 根据id查询记录
type UserIdReq struct {
	g.Meta `path:"/user" method:"get" summary:"查询" tags:"User"`
	Id     int `json:"id" in:"query" dc:"主键ID"`
}
type UserIdRes struct {
	g.Meta   `mime:"json" example:"obj"`
	Instance *entity.User `json:"form" dc:"对象"`
}

// GetList 分页查询
type UserListReq struct {
	g.Meta `path:"/user/list" method:"get" summary:"分页查询" tags:"User"`
	Page   *base.Page `json:"page" in:"query" dc:"分页数据"`
}
type UserListRes struct {
	g.Meta `mime:"json" example:"obj"`
	List   []*entity.User `json:"list" dc:"对象数组"`
}
