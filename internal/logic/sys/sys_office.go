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

type sSysOffice struct{}

func init() {
	service.RegisterSysOffice(NewSysOffice())
}

// 包装对外暴露对象
func NewSysOffice() *sSysOffice {
	return &sSysOffice{}
}

// 创建记录
func (s sSysOffice) Creat(ctx context.Context, req *api.SysOfficeCreatReq) (err error) {
	res1, err1 := dao.SysOffice.Ctx(ctx).Data(req).InsertIgnore()
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
func (s sSysOffice) Delete(ctx context.Context, id int) (err error) {
	res1, err1 := dao.SysOffice.Ctx(ctx).WhereIn(dao.SysOffice.Columns().Id, id).Delete()
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
func (s sSysOffice) DeleteByIds(ctx context.Context, ids []int) (err error) {
	res1, err1 := dao.SysOffice.Ctx(ctx).WhereIn(dao.SysOffice.Columns().Id, ids).Delete()
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
func (s sSysOffice) Update(ctx context.Context, req *api.SysOfficeUpdateReq) (err error) {
	res1, err1 := dao.SysOffice.Ctx(ctx).Data(req).Where(dao.SysOffice.Columns().Id, req.Form.Id).Update()
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
func (s sSysOffice) GetId(ctx context.Context, id int) (res *entity.SysOffice, err error) {
	var one *entity.SysOffice
	err = dao.SysOffice.Ctx(ctx).Where(dao.SysOffice.Columns().Id, id).Scan(&one)
	if err != nil {
		return nil, err
	}
	res = one
	return
}

// GetList 分页查询
func (s sSysOffice) GetList(ctx context.Context, page *base.Page) (res []*entity.SysOffice, err error) {
	var many []*entity.SysOffice
	start := (page.PageNum - 1) * 10
	limit := page.PageSize
	err = dao.SysOffice.Ctx(ctx).Limit(start, limit).Scan(&many)
	if err != nil {
		return nil, err
	}
	res = many
	return
}
