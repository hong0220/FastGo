// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysUserRole is the golang structure of table sys_user_role for DAO operations like Where/Data.
type SysUserRole struct {
	g.Meta     `orm:"table:sys_user_role, do:true"`
	Id         interface{} // 主键ID
	RoleId     interface{} // 角色ID
	UserId     interface{} // 用户ID
	CreateBy   interface{} // 创建者
	CreateDate *gtime.Time // 创建时间
	UpdateBy   interface{} // 更新者
	UpdateDate *gtime.Time // 更新时间
	Status     interface{} // 状态,0:删除,1:有效
}