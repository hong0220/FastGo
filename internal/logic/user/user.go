package user

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/hong0220/FastGo/internal/api"
	"github.com/hong0220/FastGo/internal/dao"
	"github.com/hong0220/FastGo/internal/model/entity"
	"github.com/hong0220/FastGo/internal/service"
	"github.com/hong0220/FastGo/pkg/common/base"
)

type sUser struct{}

func init() {
	service.RegisterUser(New())
}

// 包装对外暴露对象
func New() *sUser {
	return &sUser{}
}

// 创建记录
func (s sUser) Creat(ctx context.Context, req *api.UserCreatReq) (err error) {
	res1, err1 := dao.User.Ctx(ctx).Data(req).InsertIgnore()
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
func (s sUser) Delete(ctx context.Context, id int) (err error) {
	res1, err1 := dao.User.Ctx(ctx).WhereIn(dao.User.Columns().Id, id).Delete()
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
func (s sUser) DeleteByIds(ctx context.Context, ids []int) (err error) {
	res1, err1 := dao.User.Ctx(ctx).WhereIn(dao.User.Columns().Id, ids).Delete()
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
func (s sUser) Update(ctx context.Context, req *api.UserUpdateReq) (err error) {
	res1, err1 := dao.User.Ctx(ctx).Data(req).Where(dao.User.Columns().Id, req.Form.Id).Update()
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
func (s sUser) GetId(ctx context.Context, id int) (res *entity.User, err error) {
	var one *entity.User
	err = dao.User.Ctx(ctx).Where(dao.User.Columns().Id, id).Scan(&one)
	if err != nil {
		return nil, err
	}
	res = one
	return
}

// GetList 分页查询
func (s sUser) GetList(ctx context.Context, page *base.Page) (res []*entity.User, err error) {
	var many []*entity.User
	start := (page.PageNum - 1) * 10
	limit := page.PageSize
	err = dao.User.Ctx(ctx).Limit(start, limit).Scan(&many)
	if err != nil {
		return nil, err
	}
	res = many
	return
}
