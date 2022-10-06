package api

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/hong0220/FastGo/internal/model/entity"
	"github.com/hong0220/FastGo/pkg/common/base"
)

// 创建记录
type SysRoleCreatReq struct {
	g.Meta `path:"/sysrole" method:"post" summary:"新建" tags:"SysRole"`
	Form   *entity.SysRole `json:"form" in:"query"  dc:"Create对象"`
}
type SysRoleCreatRes struct {
	g.Meta    `mime:"json" example:"obj"`
	IsSuccess bool `json:"isSuccess" dc:"是否成功"`
}

// Delete 根据id删除单条记录
type SysRoleDeleteReq struct {
	g.Meta `path:"/sysrole" method:"delete" summary:"删除" tags:"SysRole"`
	Id     int `json:"id" in:"query" dc:"主键ID"`
}
type SysRoleDeleteRes struct {
	g.Meta    `mime:"json" example:"obj"`
	IsSuccess bool `json:"isSuccess" dc:"是否成功"`
}

// DeleteByIds 根据id数组批量删除记录
type SysRoleDeleteByIdsReq struct {
	g.Meta `path:"/sysrole/deleteByIds" method:"delete" summary:"批量删除" tags:"SysRole"`
	Ids    []int `json:"ids" in:"query" dc:"主键ID数组"`
}
type SysRoleDeleteByIdsRes struct {
	g.Meta    `mime:"json" example:"obj"`
	IsSuccess bool `json:"isSuccess" dc:"是否成功"`
}

// Update 修改记录
type SysRoleUpdateReq struct {
	g.Meta `path:"/sysrole" method:"put" summary:"修改" tags:"SysRole"`
	Form   *entity.SysRole `json:"form" in:"query" dc:"Update对象"`
}
type SysRoleUpdateRes struct {
	g.Meta    `mime:"json" example:"obj"`
	IsSuccess bool `json:"isSuccess" dc:"是否成功"`
}

// GetId 根据id查询记录
type SysRoleIdReq struct {
	g.Meta `path:"/sysrole" method:"get" summary:"查询" tags:"SysRole"`
	Id     int `json:"id" in:"query" dc:"主键ID"`
}
type SysRoleIdRes struct {
	g.Meta   `mime:"json" example:"obj"`
	Instance *entity.SysRole `json:"form" dc:"对象"`
}

// GetList 分页查询
type SysRoleListReq struct {
	g.Meta `path:"/sysrole/list" method:"get" summary:"分页查询" tags:"SysRole"`
	Page   *base.Page `json:"page" in:"query" dc:"分页数据"`
}
type SysRoleListRes struct {
	g.Meta `mime:"json" example:"obj"`
	List   []*entity.SysRole `json:"list" dc:"对象数组"`
}
