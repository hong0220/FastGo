// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysOffice is the golang structure of table sys_office for DAO operations like Where/Data.
type SysOffice struct {
	g.Meta     `orm:"table:sys_office, do:true"`
	Id         interface{} // 编号
	ParentId   interface{} // 父级编号
	ParentIds  interface{} // 所有父级编号
	AreaId     interface{} // 归属区域
	Code       interface{} // 区域编码
	Name       interface{} // 机构名称
	Type       interface{} // 机构类型
	Grade      interface{} // 机构等级
	Address    interface{} // 联系地址
	ZipCode    interface{} // 邮政编码
	Master     interface{} // 负责人
	Phone      interface{} // 电话
	Fax        interface{} // 传真
	Email      interface{} // 邮箱
	CreateBy   interface{} // 创建者
	CreateDate *gtime.Time // 创建时间
	UpdateBy   interface{} // 更新者
	UpdateDate *gtime.Time // 更新时间
	Remarks    interface{} // 备注信息
	DelFlag    interface{} // 删除标记
	Icon       interface{} //
}
