// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// SysTask is the golang structure for table sys_task.
type SysTask struct {
	Id          int64  `json:"id"          description:""`
	Name        string `json:"name"        description:""`
	Cron        string `json:"cron"        description:""`
	BeanClass   string `json:"beanClass"   description:""`
	BeanName    string `json:"beanName"    description:""`
	MethodName  string `json:"methodName"  description:""`
	IsStart     int    `json:"isStart"     description:""`
	Description string `json:"description" description:""`
}
