package api

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/hong0220/FastGo/internal/model/entity"
	"github.com/hong0220/FastGo/pkg/common/base"
)

// 创建记录
type SysResourceCreatReq struct {
	g.Meta `path:"/sysresource" method:"post" summary:"新建" tags:"SysResource"`
	Form   *entity.SysResource `json:"form" in:"query"  dc:"Create对象"`
}
type SysResourceCreatRes struct {
	g.Meta    `mime:"json" example:"obj"`
	IsSuccess bool `json:"isSuccess" dc:"是否成功"`
}

// Delete 根据id删除单条记录
type SysResourceDeleteReq struct {
	g.Meta `path:"/sysresource" method:"delete" summary:"删除" tags:"SysResource"`
	Id     int `json:"id" in:"query" dc:"主键ID"`
}
type SysResourceDeleteRes struct {
	g.Meta    `mime:"json" example:"obj"`
	IsSuccess bool `json:"isSuccess" dc:"是否成功"`
}

// DeleteByIds 根据id数组批量删除记录
type SysResourceDeleteByIdsReq struct {
	g.Meta `path:"/sysresource/deleteByIds" method:"delete" summary:"批量删除" tags:"SysResource"`
	Ids    []int `json:"ids" in:"query" dc:"主键ID数组"`
}
type SysResourceDeleteByIdsRes struct {
	g.Meta    `mime:"json" example:"obj"`
	IsSuccess bool `json:"isSuccess" dc:"是否成功"`
}

// Update 修改记录
type SysResourceUpdateReq struct {
	g.Meta `path:"/sysresource" method:"put" summary:"修改" tags:"SysResource"`
	Form   *entity.SysResource `json:"form" in:"query" dc:"Update对象"`
}
type SysResourceUpdateRes struct {
	g.Meta    `mime:"json" example:"obj"`
	IsSuccess bool `json:"isSuccess" dc:"是否成功"`
}

// GetId 根据id查询记录
type SysResourceIdReq struct {
	g.Meta `path:"/sysresource" method:"get" summary:"查询" tags:"SysResource"`
	Id     int `json:"id" in:"query" dc:"主键ID"`
}
type SysResourceIdRes struct {
	g.Meta   `mime:"json" example:"obj"`
	Instance *entity.SysResource `json:"form" dc:"对象"`
}

// GetList 分页查询
type SysResourceListReq struct {
	g.Meta `path:"/sysresource/list" method:"get" summary:"分页查询" tags:"SysResource"`
	Page   *base.Page `json:"page" in:"query" dc:"分页数据"`
}
type SysResourceListRes struct {
	g.Meta `mime:"json" example:"obj"`
	List   []*entity.SysResource `json:"list" dc:"对象数组"`
}
