// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// SysTask is the golang structure of table sys_task for DAO operations like Where/Data.
type SysTask struct {
	g.Meta      `orm:"table:sys_task, do:true"`
	Id          interface{} //
	Name        interface{} //
	Cron        interface{} //
	BeanClass   interface{} //
	BeanName    interface{} //
	MethodName  interface{} //
	IsStart     interface{} //
	Description interface{} //
}
