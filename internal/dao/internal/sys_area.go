// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysAreaDao is the data access object for table sys_area.
type SysAreaDao struct {
	table   string         // table is the underlying table name of the DAO.
	group   string         // group is the database configuration group name of current DAO.
	columns SysAreaColumns // columns contains all the column names of Table for convenient usage.
}

// SysAreaColumns defines and stores column names for table sys_area.
type SysAreaColumns struct {
	Id         string // 编号
	ParentId   string // 父级编号
	ParentIds  string // 所有父级编号
	Code       string // 区域编码
	Name       string // 区域名称
	Type       string // 区域类型
	CreateBy   string // 创建者
	CreateDate string // 创建时间
	UpdateBy   string // 更新者
	UpdateDate string // 更新时间
	Remarks    string // 备注信息
	DelFlag    string // 删除标记(0活null 正常 1,删除)
	Icon       string //
}

//  sysAreaColumns holds the columns for table sys_area.
var sysAreaColumns = SysAreaColumns{
	Id:         "id",
	ParentId:   "parent_id",
	ParentIds:  "parent_ids",
	Code:       "code",
	Name:       "name",
	Type:       "type",
	CreateBy:   "create_by",
	CreateDate: "create_date",
	UpdateBy:   "update_by",
	UpdateDate: "update_date",
	Remarks:    "remarks",
	DelFlag:    "del_flag",
	Icon:       "icon",
}

// NewSysAreaDao creates and returns a new DAO object for table data access.
func NewSysAreaDao() *SysAreaDao {
	return &SysAreaDao{
		group:   "default",
		table:   "sys_area",
		columns: sysAreaColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SysAreaDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SysAreaDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SysAreaDao) Columns() SysAreaColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SysAreaDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SysAreaDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SysAreaDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}