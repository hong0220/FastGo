package api

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/hong0220/FastGo/internal/model/entity"
	"github.com/hong0220/FastGo/pkg/common/base"
)

// 创建记录
type SysUserCreatReq struct {
	g.Meta `path:"/sysuser" method:"post" summary:"新建" tags:"SysUser"`
	Form   *entity.SysUser `json:"form" in:"query"  dc:"Create对象"`
}
type SysUserCreatRes struct {
	g.Meta    `mime:"json" example:"obj"`
	IsSuccess bool `json:"isSuccess" dc:"是否成功"`
}

// Delete 根据id删除单条记录
type SysUserDeleteReq struct {
	g.Meta `path:"/sysuser" method:"delete" summary:"删除" tags:"SysUser"`
	Id     int `json:"id" in:"query" dc:"主键ID"`
}
type SysUserDeleteRes struct {
	g.Meta    `mime:"json" example:"obj"`
	IsSuccess bool `json:"isSuccess" dc:"是否成功"`
}

// DeleteByIds 根据id数组批量删除记录
type SysUserDeleteByIdsReq struct {
	g.Meta `path:"/sysuser/deleteByIds" method:"delete" summary:"批量删除" tags:"SysUser"`
	Ids    []int `json:"ids" in:"query" dc:"主键ID数组"`
}
type SysUserDeleteByIdsRes struct {
	g.Meta    `mime:"json" example:"obj"`
	IsSuccess bool `json:"isSuccess" dc:"是否成功"`
}

// Update 修改记录
type SysUserUpdateReq struct {
	g.Meta `path:"/sysuser" method:"put" summary:"修改" tags:"SysUser"`
	Form   *entity.SysUser `json:"form" in:"query" dc:"Update对象"`
}
type SysUserUpdateRes struct {
	g.Meta    `mime:"json" example:"obj"`
	IsSuccess bool `json:"isSuccess" dc:"是否成功"`
}

// GetId 根据id查询记录
type SysUserIdReq struct {
	g.Meta `path:"/sysuser" method:"get" summary:"查询" tags:"SysUser"`
	Id     int `json:"id" in:"query" dc:"主键ID"`
}
type SysUserIdRes struct {
	g.Meta   `mime:"json" example:"obj"`
	Instance *entity.SysUser `json:"form" dc:"对象"`
}

// GetList 分页查询
type SysUserListReq struct {
	g.Meta `path:"/sysuser/list" method:"get" summary:"分页查询" tags:"SysUser"`
	Page   *base.Page `json:"page" in:"query" dc:"分页数据"`
}
type SysUserListRes struct {
	g.Meta `mime:"json" example:"obj"`
	List   []*entity.SysUser `json:"list" dc:"对象数组"`
}
