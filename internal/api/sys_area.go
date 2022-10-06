package api

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/hong0220/FastGo/internal/model/entity"
	"github.com/hong0220/FastGo/pkg/common/base"
)

// 创建记录
type SysAreaCreatReq struct {
	g.Meta `path:"/sysarea" method:"post" summary:"新建" tags:"SysArea"`
	Form   *entity.SysArea `json:"form" in:"query"  dc:"Create对象"`
}
type SysAreaCreatRes struct {
	g.Meta    `mime:"json" example:"obj"`
	IsSuccess bool `json:"isSuccess" dc:"是否成功"`
}

// Delete 根据id删除单条记录
type SysAreaDeleteReq struct {
	g.Meta `path:"/sysarea" method:"delete" summary:"删除" tags:"SysArea"`
	Id     int `json:"id" in:"query" dc:"主键ID"`
}
type SysAreaDeleteRes struct {
	g.Meta    `mime:"json" example:"obj"`
	IsSuccess bool `json:"isSuccess" dc:"是否成功"`
}

// DeleteByIds 根据id数组批量删除记录
type SysAreaDeleteByIdsReq struct {
	g.Meta `path:"/sysarea/deleteByIds" method:"delete" summary:"批量删除" tags:"SysArea"`
	Ids    []int `json:"ids" in:"query" dc:"主键ID数组"`
}
type SysAreaDeleteByIdsRes struct {
	g.Meta    `mime:"json" example:"obj"`
	IsSuccess bool `json:"isSuccess" dc:"是否成功"`
}

// Update 修改记录
type SysAreaUpdateReq struct {
	g.Meta `path:"/sysarea" method:"put" summary:"修改" tags:"SysArea"`
	Form   *entity.SysArea `json:"form" in:"query" dc:"Update对象"`
}
type SysAreaUpdateRes struct {
	g.Meta    `mime:"json" example:"obj"`
	IsSuccess bool `json:"isSuccess" dc:"是否成功"`
}

// GetId 根据id查询记录
type SysAreaIdReq struct {
	g.Meta `path:"/sysarea" method:"get" summary:"查询" tags:"SysArea"`
	Id     int `json:"id" in:"query" dc:"主键ID"`
}
type SysAreaIdRes struct {
	g.Meta   `mime:"json" example:"obj"`
	Instance *entity.SysArea `json:"form" dc:"对象"`
}

// GetList 分页查询
type SysAreaListReq struct {
	g.Meta `path:"/sysarea/list" method:"get" summary:"分页查询" tags:"SysArea"`
	Page   *base.Page `json:"page" in:"query" dc:"分页数据"`
}
type SysAreaListRes struct {
	g.Meta `mime:"json" example:"obj"`
	List   []*entity.SysArea `json:"list" dc:"对象数组"`
}
