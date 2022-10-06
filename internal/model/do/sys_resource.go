// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysResource is the golang structure of table sys_resource for DAO operations like Where/Data.
type SysResource struct {
	g.Meta        `orm:"table:sys_resource, do:true"`
	Id            interface{} //
	Name          interface{} // 资源名称
	Common        interface{} // 是否是公共资源(0.不是 1.是)
	Icon          interface{} // 图标
	Sort          interface{} // 排序号
	ParentId      interface{} // 父级id
	Type          interface{} // 类型(0.菜单 1.按钮)
	Url           interface{} // 链接
	Description   interface{} // 描述
	Status        interface{} // 状态(0.正常 1.禁用)
	ParentIds     interface{} // 父级集合
	CreateDate    *gtime.Time //
	UpdateDate    *gtime.Time //
	CreateBy      interface{} //
	UpdateBy      interface{} //
	DelFlag       interface{} //
	PermissionStr interface{} //
}
