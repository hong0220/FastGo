// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysTaskDao is the data access object for table sys_task.
type SysTaskDao struct {
	table   string         // table is the underlying table name of the DAO.
	group   string         // group is the database configuration group name of current DAO.
	columns SysTaskColumns // columns contains all the column names of Table for convenient usage.
}

// SysTaskColumns defines and stores column names for table sys_task.
type SysTaskColumns struct {
	Id          string //
	Name        string //
	Cron        string //
	BeanClass   string //
	BeanName    string //
	MethodName  string //
	IsStart     string //
	Description string //
}

//  sysTaskColumns holds the columns for table sys_task.
var sysTaskColumns = SysTaskColumns{
	Id:          "id",
	Name:        "name",
	Cron:        "cron",
	BeanClass:   "bean_class",
	BeanName:    "bean_name",
	MethodName:  "method_name",
	IsStart:     "is_start",
	Description: "description",
}

// NewSysTaskDao creates and returns a new DAO object for table data access.
func NewSysTaskDao() *SysTaskDao {
	return &SysTaskDao{
		group:   "default",
		table:   "sys_task",
		columns: sysTaskColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SysTaskDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SysTaskDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SysTaskDao) Columns() SysTaskColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SysTaskDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SysTaskDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SysTaskDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
