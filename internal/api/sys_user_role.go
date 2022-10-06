package api

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/hong0220/FastGo/internal/model/entity"
	"github.com/hong0220/FastGo/pkg/common/base"
)

// 创建记录
type SysUserRoleCreatReq struct {
	g.Meta `path:"/sysuserrole" method:"post" summary:"新建" tags:"SysUserRole"`
	Form   *entity.SysUserRole `json:"form" in:"query"  dc:"Create对象"`
}
type SysUserRoleCreatRes struct {
	g.Meta    `mime:"json" example:"obj"`
	IsSuccess bool `json:"isSuccess" dc:"是否成功"`
}

// Delete 根据id删除单条记录
type SysUserRoleDeleteReq struct {
	g.Meta `path:"/sysuserrole" method:"delete" summary:"删除" tags:"SysUserRole"`
	Id     int `json:"id" in:"query" dc:"主键ID"`
}
type SysUserRoleDeleteRes struct {
	g.Meta    `mime:"json" example:"obj"`
	IsSuccess bool `json:"isSuccess" dc:"是否成功"`
}

// DeleteByIds 根据id数组批量删除记录
type SysUserRoleDeleteByIdsReq struct {
	g.Meta `path:"/sysuserrole/deleteByIds" method:"delete" summary:"批量删除" tags:"SysUserRole"`
	Ids    []int `json:"ids" in:"query" dc:"主键ID数组"`
}
type SysUserRoleDeleteByIdsRes struct {
	g.Meta    `mime:"json" example:"obj"`
	IsSuccess bool `json:"isSuccess" dc:"是否成功"`
}

// Update 修改记录
type SysUserRoleUpdateReq struct {
	g.Meta `path:"/sysuserrole" method:"put" summary:"修改" tags:"SysUserRole"`
	Form   *entity.SysUserRole `json:"form" in:"query" dc:"Update对象"`
}
type SysUserRoleUpdateRes struct {
	g.Meta    `mime:"json" example:"obj"`
	IsSuccess bool `json:"isSuccess" dc:"是否成功"`
}

// GetId 根据id查询记录
type SysUserRoleIdReq struct {
	g.Meta `path:"/sysuserrole" method:"get" summary:"查询" tags:"SysUserRole"`
	Id     int `json:"id" in:"query" dc:"主键ID"`
}
type SysUserRoleIdRes struct {
	g.Meta   `mime:"json" example:"obj"`
	Instance *entity.SysUserRole `json:"form" dc:"对象"`
}

// GetList 分页查询
type SysUserRoleListReq struct {
	g.Meta `path:"/sysuserrole/list" method:"get" summary:"分页查询" tags:"SysUserRole"`
	Page   *base.Page `json:"page" in:"query" dc:"分页数据"`
}
type SysUserRoleListRes struct {
	g.Meta `mime:"json" example:"obj"`
	List   []*entity.SysUserRole `json:"list" dc:"对象数组"`
}
