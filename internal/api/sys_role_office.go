package api

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/hong0220/FastGo/internal/model/entity"
	"github.com/hong0220/FastGo/pkg/common/base"
)

// 创建记录
type SysRoleOfficeCreatReq struct {
	g.Meta `path:"/sysroleoffice" method:"post" summary:"新建" tags:"SysRoleOffice"`
	Form   *entity.SysRoleOffice `json:"form" in:"query"  dc:"Create对象"`
}
type SysRoleOfficeCreatRes struct {
	g.Meta    `mime:"json" example:"obj"`
	IsSuccess bool `json:"isSuccess" dc:"是否成功"`
}

// Delete 根据id删除单条记录
type SysRoleOfficeDeleteReq struct {
	g.Meta `path:"/sysroleoffice" method:"delete" summary:"删除" tags:"SysRoleOffice"`
	Id     int `json:"id" in:"query" dc:"主键ID"`
}
type SysRoleOfficeDeleteRes struct {
	g.Meta    `mime:"json" example:"obj"`
	IsSuccess bool `json:"isSuccess" dc:"是否成功"`
}

// DeleteByIds 根据id数组批量删除记录
type SysRoleOfficeDeleteByIdsReq struct {
	g.Meta `path:"/sysroleoffice/deleteByIds" method:"delete" summary:"批量删除" tags:"SysRoleOffice"`
	Ids    []int `json:"ids" in:"query" dc:"主键ID数组"`
}
type SysRoleOfficeDeleteByIdsRes struct {
	g.Meta    `mime:"json" example:"obj"`
	IsSuccess bool `json:"isSuccess" dc:"是否成功"`
}

// Update 修改记录
type SysRoleOfficeUpdateReq struct {
	g.Meta `path:"/sysroleoffice" method:"put" summary:"修改" tags:"SysRoleOffice"`
	Form   *entity.SysRoleOffice `json:"form" in:"query" dc:"Update对象"`
}
type SysRoleOfficeUpdateRes struct {
	g.Meta    `mime:"json" example:"obj"`
	IsSuccess bool `json:"isSuccess" dc:"是否成功"`
}

// GetId 根据id查询记录
type SysRoleOfficeIdReq struct {
	g.Meta `path:"/sysroleoffice" method:"get" summary:"查询" tags:"SysRoleOffice"`
	Id     int `json:"id" in:"query" dc:"主键ID"`
}
type SysRoleOfficeIdRes struct {
	g.Meta   `mime:"json" example:"obj"`
	Instance *entity.SysRoleOffice `json:"form" dc:"对象"`
}

// GetList 分页查询
type SysRoleOfficeListReq struct {
	g.Meta `path:"/sysroleoffice/list" method:"get" summary:"分页查询" tags:"SysRoleOffice"`
	Page   *base.Page `json:"page" in:"query" dc:"分页数据"`
}
type SysRoleOfficeListRes struct {
	g.Meta `mime:"json" example:"obj"`
	List   []*entity.SysRoleOffice `json:"list" dc:"对象数组"`
}
