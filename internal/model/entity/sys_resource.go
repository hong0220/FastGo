// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysResource is the golang structure for table sys_resource.
type SysResource struct {
	Id            int64       `json:"id"            description:""`
	Name          string      `json:"name"          description:"资源名称"`
	Common        string      `json:"common"        description:"是否是公共资源(0.不是 1.是)"`
	Icon          string      `json:"icon"          description:"图标"`
	Sort          int         `json:"sort"          description:"排序号"`
	ParentId      int64       `json:"parentId"      description:"父级id"`
	Type          string      `json:"type"          description:"类型(0.菜单 1.按钮)"`
	Url           string      `json:"url"           description:"链接"`
	Description   string      `json:"description"   description:"描述"`
	Status        string      `json:"status"        description:"状态(0.正常 1.禁用)"`
	ParentIds     string      `json:"parentIds"     description:"父级集合"`
	CreateDate    *gtime.Time `json:"createDate"    description:""`
	UpdateDate    *gtime.Time `json:"updateDate"    description:""`
	CreateBy      string      `json:"createBy"      description:""`
	UpdateBy      string      `json:"updateBy"      description:""`
	DelFlag       string      `json:"delFlag"       description:""`
	PermissionStr string      `json:"permissionStr" description:""`
}
