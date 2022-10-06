package logic

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/hong0220/FastGo/internal/api"
	"github.com/hong0220/FastGo/internal/dao"
	"github.com/hong0220/FastGo/internal/model/entity"
	"github.com/hong0220/FastGo/internal/service"
	"github.com/hong0220/FastGo/pkg/common/base"
)

type sSysRole struct{}

func init() {
	service.RegisterSysRole(NewSysRole())
}

// 包装对外暴露对象
func NewSysRole() *sSysRole {
	return &sSysRole{}
}

// 创建记录
func (s sSysRole) Creat(ctx context.Context, req *api.SysRoleCreatReq) (err error) {
	res1, err1 := dao.SysRole.Ctx(ctx).Data(req).InsertIgnore()
	if err1 != nil {
		return err1
	}
	res2, err2 := res1.RowsAffected()
	if err2 != nil {
		return err2
	}
	if res2 < 1 {
		err = gerror.New("影响行数小于1")
		return err
	}
	return
}

// Delete 根据id删除单条记录
func (s sSysRole) Delete(ctx context.Context, id int) (err error) {
	res1, err1 := dao.SysRole.Ctx(ctx).WhereIn(dao.SysRole.Columns().Id, id).Delete()
	if err1 != nil {
		return err1
	}
	res2, err2 := res1.RowsAffected()
	if err2 != nil {
		return err2
	}
	if res2 == 0 {
		err = gerror.New("影响行数为0")
		return err
	}
	return
}

// DeleteByIds 根据id数组批量删除记录
func (s sSysRole) DeleteByIds(ctx context.Context, ids []int) (err error) {
	res1, err1 := dao.SysRole.Ctx(ctx).WhereIn(dao.SysRole.Columns().Id, ids).Delete()
	if err1 != nil {
		return err1
	}
	res2, err2 := res1.RowsAffected()
	if err2 != nil {
		return err2
	}
	if res2 == 0 {
		err = gerror.New("影响行数为0")
		return err
	}
	return
}

// Update 修改记录
func (s sSysRole) Update(ctx context.Context, req *api.SysRoleUpdateReq) (err error) {
	res1, err1 := dao.SysRole.Ctx(ctx).Data(req).Where(dao.SysRole.Columns().Id, req.Form.Id).Update()
	if err1 != nil {
		return err1
	}
	res2, err2 := res1.RowsAffected()
	if err2 != nil {
		return err2
	}
	if res2 == 0 {
		err = gerror.New("影响行数为0")
		return err
	}
	return
}

// GetId 根据id查询记录
func (s sSysRole) GetId(ctx context.Context, id int) (res *entity.SysRole, err error) {
	var one *entity.SysRole
	err = dao.SysRole.Ctx(ctx).Where(dao.SysRole.Columns().Id, id).Scan(&one)
	if err != nil {
		return nil, err
	}
	res = one
	return
}

// GetList 分页查询
func (s sSysRole) GetList(ctx context.Context, page *base.Page) (res []*entity.SysRole, err error) {
	var many []*entity.SysRole
	start := (page.PageNum - 1) * 10
	limit := page.PageSize
	err = dao.SysRole.Ctx(ctx).Limit(start, limit).Scan(&many)
	if err != nil {
		return nil, err
	}
	res = many
	return
}
