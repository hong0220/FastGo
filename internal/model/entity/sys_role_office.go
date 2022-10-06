// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysRoleOffice is the golang structure for table sys_role_office.
type SysRoleOffice struct {
	RoleId     int64       `json:"roleId"     description:"角色编号"`
	OfficeId   int64       `json:"officeId"   description:"机构编号"`
	Id         int64       `json:"id"         description:""`
	CreateBy   string      `json:"createBy"   description:""`
	CreateDate *gtime.Time `json:"createDate" description:""`
	UpdateBy   string      `json:"updateBy"   description:""`
	UpdateDate *gtime.Time `json:"updateDate" description:""`
	DelFlag    string      `json:"delFlag"    description:""`
}
