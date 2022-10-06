// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysOffice is the golang structure for table sys_office.
type SysOffice struct {
	Id         int64       `json:"id"         description:"编号"`
	ParentId   int64       `json:"parentId"   description:"父级编号"`
	ParentIds  string      `json:"parentIds"  description:"所有父级编号"`
	AreaId     int64       `json:"areaId"     description:"归属区域"`
	Code       string      `json:"code"       description:"区域编码"`
	Name       string      `json:"name"       description:"机构名称"`
	Type       string      `json:"type"       description:"机构类型"`
	Grade      string      `json:"grade"      description:"机构等级"`
	Address    string      `json:"address"    description:"联系地址"`
	ZipCode    string      `json:"zipCode"    description:"邮政编码"`
	Master     string      `json:"master"     description:"负责人"`
	Phone      string      `json:"phone"      description:"电话"`
	Fax        string      `json:"fax"        description:"传真"`
	Email      string      `json:"email"      description:"邮箱"`
	CreateBy   string      `json:"createBy"   description:"创建者"`
	CreateDate *gtime.Time `json:"createDate" description:"创建时间"`
	UpdateBy   string      `json:"updateBy"   description:"更新者"`
	UpdateDate *gtime.Time `json:"updateDate" description:"更新时间"`
	Remarks    string      `json:"remarks"    description:"备注信息"`
	DelFlag    string      `json:"delFlag"    description:"删除标记"`
	Icon       string      `json:"icon"       description:""`
}