package api

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/hong0220/FastGo/internal/model/entity"
	"github.com/hong0220/FastGo/pkg/common/base"
)

// 创建记录
type SysLogCreatReq struct {
	g.Meta `path:"/syslog" method:"post" summary:"新建" tags:"SysLog"`
	Form   *entity.SysLog `json:"form" in:"query"  dc:"Create对象"`
}
type SysLogCreatRes struct {
	g.Meta    `mime:"json" example:"obj"`
	IsSuccess bool `json:"isSuccess" dc:"是否成功"`
}

// Delete 根据id删除单条记录
type SysLogDeleteReq struct {
	g.Meta `path:"/syslog" method:"delete" summary:"删除" tags:"SysLog"`
	Id     int `json:"id" in:"query" dc:"主键ID"`
}
type SysLogDeleteRes struct {
	g.Meta    `mime:"json" example:"obj"`
	IsSuccess bool `json:"isSuccess" dc:"是否成功"`
}

// DeleteByIds 根据id数组批量删除记录
type SysLogDeleteByIdsReq struct {
	g.Meta `path:"/syslog/deleteByIds" method:"delete" summary:"批量删除" tags:"SysLog"`
	Ids    []int `json:"ids" in:"query" dc:"主键ID数组"`
}
type SysLogDeleteByIdsRes struct {
	g.Meta    `mime:"json" example:"obj"`
	IsSuccess bool `json:"isSuccess" dc:"是否成功"`
}

// Update 修改记录
type SysLogUpdateReq struct {
	g.Meta `path:"/syslog" method:"put" summary:"修改" tags:"SysLog"`
	Form   *entity.SysLog `json:"form" in:"query" dc:"Update对象"`
}
type SysLogUpdateRes struct {
	g.Meta    `mime:"json" example:"obj"`
	IsSuccess bool `json:"isSuccess" dc:"是否成功"`
}

// GetId 根据id查询记录
type SysLogIdReq struct {
	g.Meta `path:"/syslog" method:"get" summary:"查询" tags:"SysLog"`
	Id     int `json:"id" in:"query" dc:"主键ID"`
}
type SysLogIdRes struct {
	g.Meta   `mime:"json" example:"obj"`
	Instance *entity.SysLog `json:"form" dc:"对象"`
}

// GetList 分页查询
type SysLogListReq struct {
	g.Meta `path:"/syslog/list" method:"get" summary:"分页查询" tags:"SysLog"`
	Page   *base.Page `json:"page" in:"query" dc:"分页数据"`
}
type SysLogListRes struct {
	g.Meta `mime:"json" example:"obj"`
	List   []*entity.SysLog `json:"list" dc:"对象数组"`
}
