package api

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/hong0220/FastGo/internal/model/entity"
	"github.com/hong0220/FastGo/pkg/common/base"
)

// 创建记录
type SysOfficeCreatReq struct {
	g.Meta `path:"/sysoffice" method:"post" summary:"新建" tags:"SysOffice"`
	Form   *entity.SysOffice `json:"form" in:"query"  dc:"Create对象"`
}
type SysOfficeCreatRes struct {
	g.Meta    `mime:"json" example:"obj"`
	IsSuccess bool `json:"isSuccess" dc:"是否成功"`
}

// Delete 根据id删除单条记录
type SysOfficeDeleteReq struct {
	g.Meta `path:"/sysoffice" method:"delete" summary:"删除" tags:"SysOffice"`
	Id     int `json:"id" in:"query" dc:"主键ID"`
}
type SysOfficeDeleteRes struct {
	g.Meta    `mime:"json" example:"obj"`
	IsSuccess bool `json:"isSuccess" dc:"是否成功"`
}

// DeleteByIds 根据id数组批量删除记录
type SysOfficeDeleteByIdsReq struct {
	g.Meta `path:"/sysoffice/deleteByIds" method:"delete" summary:"批量删除" tags:"SysOffice"`
	Ids    []int `json:"ids" in:"query" dc:"主键ID数组"`
}
type SysOfficeDeleteByIdsRes struct {
	g.Meta    `mime:"json" example:"obj"`
	IsSuccess bool `json:"isSuccess" dc:"是否成功"`
}

// Update 修改记录
type SysOfficeUpdateReq struct {
	g.Meta `path:"/sysoffice" method:"put" summary:"修改" tags:"SysOffice"`
	Form   *entity.SysOffice `json:"form" in:"query" dc:"Update对象"`
}
type SysOfficeUpdateRes struct {
	g.Meta    `mime:"json" example:"obj"`
	IsSuccess bool `json:"isSuccess" dc:"是否成功"`
}

// GetId 根据id查询记录
type SysOfficeIdReq struct {
	g.Meta `path:"/sysoffice" method:"get" summary:"查询" tags:"SysOffice"`
	Id     int `json:"id" in:"query" dc:"主键ID"`
}
type SysOfficeIdRes struct {
	g.Meta   `mime:"json" example:"obj"`
	Instance *entity.SysOffice `json:"form" dc:"对象"`
}

// GetList 分页查询
type SysOfficeListReq struct {
	g.Meta `path:"/sysoffice/list" method:"get" summary:"分页查询" tags:"SysOffice"`
	Page   *base.Page `json:"page" in:"query" dc:"分页数据"`
}
type SysOfficeListRes struct {
	g.Meta `mime:"json" example:"obj"`
	List   []*entity.SysOffice `json:"list" dc:"对象数组"`
}
