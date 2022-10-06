package api

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/hong0220/FastGo/internal/model/entity"
	"github.com/hong0220/FastGo/pkg/common/base"
)

// 创建记录
type SysTaskCreatReq struct {
	g.Meta `path:"/systask" method:"post" summary:"新建" tags:"SysTask"`
	Form   *entity.SysTask `json:"form" in:"query"  dc:"Create对象"`
}
type SysTaskCreatRes struct {
	g.Meta    `mime:"json" example:"obj"`
	IsSuccess bool `json:"isSuccess" dc:"是否成功"`
}

// Delete 根据id删除单条记录
type SysTaskDeleteReq struct {
	g.Meta `path:"/systask" method:"delete" summary:"删除" tags:"SysTask"`
	Id     int `json:"id" in:"query" dc:"主键ID"`
}
type SysTaskDeleteRes struct {
	g.Meta    `mime:"json" example:"obj"`
	IsSuccess bool `json:"isSuccess" dc:"是否成功"`
}

// DeleteByIds 根据id数组批量删除记录
type SysTaskDeleteByIdsReq struct {
	g.Meta `path:"/systask/deleteByIds" method:"delete" summary:"批量删除" tags:"SysTask"`
	Ids    []int `json:"ids" in:"query" dc:"主键ID数组"`
}
type SysTaskDeleteByIdsRes struct {
	g.Meta    `mime:"json" example:"obj"`
	IsSuccess bool `json:"isSuccess" dc:"是否成功"`
}

// Update 修改记录
type SysTaskUpdateReq struct {
	g.Meta `path:"/systask" method:"put" summary:"修改" tags:"SysTask"`
	Form   *entity.SysTask `json:"form" in:"query" dc:"Update对象"`
}
type SysTaskUpdateRes struct {
	g.Meta    `mime:"json" example:"obj"`
	IsSuccess bool `json:"isSuccess" dc:"是否成功"`
}

// GetId 根据id查询记录
type SysTaskIdReq struct {
	g.Meta `path:"/systask" method:"get" summary:"查询" tags:"SysTask"`
	Id     int `json:"id" in:"query" dc:"主键ID"`
}
type SysTaskIdRes struct {
	g.Meta   `mime:"json" example:"obj"`
	Instance *entity.SysTask `json:"form" dc:"对象"`
}

// GetList 分页查询
type SysTaskListReq struct {
	g.Meta `path:"/systask/list" method:"get" summary:"分页查询" tags:"SysTask"`
	Page   *base.Page `json:"page" in:"query" dc:"分页数据"`
}
type SysTaskListRes struct {
	g.Meta `mime:"json" example:"obj"`
	List   []*entity.SysTask `json:"list" dc:"对象数组"`
}
