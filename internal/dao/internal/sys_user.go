// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysUserDao is the data access object for table sys_user.
type SysUserDao struct {
	table   string         // table is the underlying table name of the DAO.
	group   string         // group is the database configuration group name of current DAO.
	columns SysUserColumns // columns contains all the column names of Table for convenient usage.
}

// SysUserColumns defines and stores column names for table sys_user.
type SysUserColumns struct {
	Id         string // 主键ID
	Username   string // 用户名
	Password   string // 密码
	JobNumber  string // 工号
	Email      string // 邮箱
	Phone      string // 电话
	Mobile     string // 手机
	OfficeId   string // 部门ID
	CompanyId  string // 公司ID
	UserType   string // 用户类型(0.普通 1.超管)
	LoginIp    string // 最后登陆IP
	LoginDate  string // 登陆时间
	CreateBy   string // 创建者
	CreateDate string // 创建时间
	UpdateBy   string // 更新者
	UpdateDate string // 更新时间
	Remark     string // 备注
	Status     string // 状态,0:删除,1:有效
}

//  sysUserColumns holds the columns for table sys_user.
var sysUserColumns = SysUserColumns{
	Id:         "id",
	Username:   "username",
	Password:   "password",
	JobNumber:  "job_number",
	Email:      "email",
	Phone:      "phone",
	Mobile:     "mobile",
	OfficeId:   "office_id",
	CompanyId:  "company_id",
	UserType:   "user_type",
	LoginIp:    "login_ip",
	LoginDate:  "login_date",
	CreateBy:   "create_by",
	CreateDate: "create_date",
	UpdateBy:   "update_by",
	UpdateDate: "update_date",
	Remark:     "remark",
	Status:     "status",
}

// NewSysUserDao creates and returns a new DAO object for table data access.
func NewSysUserDao() *SysUserDao {
	return &SysUserDao{
		group:   "default",
		table:   "sys_user",
		columns: sysUserColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SysUserDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SysUserDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SysUserDao) Columns() SysUserColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SysUserDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SysUserDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SysUserDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
