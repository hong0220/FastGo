// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysLog is the golang structure for table sys_log.
type SysLog struct {
	Id          int64       `json:"id"          description:"编号"`
	Type        string      `json:"type"        description:"日志类型"`
	CreateBy    string      `json:"createBy"    description:"创建者"`
	CreateDate  *gtime.Time `json:"createDate"  description:"创建时间"`
	RemoteAddr  string      `json:"remoteAddr"  description:"操作IP地址"`
	UserAgent   string      `json:"userAgent"   description:"用户代理"`
	RequestUri  string      `json:"requestUri"  description:"请求URI"`
	Method      string      `json:"method"      description:"操作方式"`
	Params      string      `json:"params"      description:"操作提交的数据"`
	Exception   string      `json:"exception"   description:"异常信息"`
	Description string      `json:"description" description:"描述"`
}