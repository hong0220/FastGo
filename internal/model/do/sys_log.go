// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysLog is the golang structure of table sys_log for DAO operations like Where/Data.
type SysLog struct {
	g.Meta      `orm:"table:sys_log, do:true"`
	Id          interface{} // 编号
	Type        interface{} // 日志类型
	CreateBy    interface{} // 创建者
	CreateDate  *gtime.Time // 创建时间
	RemoteAddr  interface{} // 操作IP地址
	UserAgent   interface{} // 用户代理
	RequestUri  interface{} // 请求URI
	Method      interface{} // 操作方式
	Params      interface{} // 操作提交的数据
	Exception   interface{} // 异常信息
	Description interface{} // 描述
}
