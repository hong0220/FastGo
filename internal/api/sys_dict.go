package api

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/hong0220/FastGo/internal/model/entity"
	"github.com/hong0220/FastGo/pkg/common/base"
)

// 创建记录
type SysDictCreatReq struct {
	g.Meta `path:"/sysDict" method:"post" summary:"新建" tags:"SysDict"`
	Form   *entity.SysDict `json:"form" in:"query"  dc:"Create对象"`
}
type SysDictCreatRes struct {
	g.Meta    `mime:"json" example:"obj"`
	IsSuccess bool `json:"isSuccess" dc:"是否成功"`
}

// Delete 根据id删除单条记录
type SysDictDeleteReq struct {
	g.Meta `path:"/sysDict" method:"delete" summary:"删除" tags:"SysDict"`
	Id     int `json:"id" in:"query" dc:"主键ID"`
}
type SysDictDeleteRes struct {
	g.Meta    `mime:"json" example:"obj"`
	IsSuccess bool `json:"isSuccess" dc:"是否成功"`
}

// DeleteByIds 根据id数组批量删除记录
type SysDictDeleteByIdsReq struct {
	g.Meta `path:"/sysDict/deleteByIds" method:"delete" summary:"批量删除" tags:"SysDict"`
	Ids    []int `json:"ids" in:"query" dc:"主键ID数组"`
}
type SysDictDeleteByIdsRes struct {
	g.Meta    `mime:"json" example:"obj"`
	IsSuccess bool `json:"isSuccess" dc:"是否成功"`
}

// Update 修改记录
type SysDictUpdateReq struct {
	g.Meta `path:"/sysDict" method:"put" summary:"修改" tags:"SysDict"`
	Form   *entity.SysDict `json:"form" in:"query" dc:"Update对象"`
}
type SysDictUpdateRes struct {
	g.Meta    `mime:"json" example:"obj"`
	IsSuccess bool `json:"isSuccess" dc:"是否成功"`
}

// GetId 根据id查询记录
type SysDictIdReq struct {
	g.Meta `path:"/sysDict" method:"get" summary:"查询" tags:"SysDict"`
	Id     int `json:"id" in:"query" dc:"主键ID"`
}
type SysDictIdRes struct {
	g.Meta   `mime:"json" example:"obj"`
	Instance *entity.SysDict `json:"form" dc:"对象"`
}

// GetList 分页查询
type SysDictListReq struct {
	g.Meta `path:"/sysDict/list" method:"get" summary:"分页查询" tags:"SysDict"`
	Page   *base.Page `json:"page" in:"query" dc:"分页数据"`
}
type SysDictListRes struct {
	g.Meta `mime:"json" example:"obj"`
	List   []*entity.SysDict `json:"list" dc:"对象数组"`
}
