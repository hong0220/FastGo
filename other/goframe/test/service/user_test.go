package service

import (
	"errors"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/hong0220/FastGo/internal/api"
	logic "github.com/hong0220/FastGo/internal/logic/sys"
	"github.com/hong0220/FastGo/pkg/common/test"
	"testing"
)

func TestSysUser(t *testing.T) {
	test.ApiConfigSetUp()

	ctx := gctx.New()

	req := &api.SysUserCreatReq{}
	logic.NewSysUser().Creat(ctx, req)

	sysUser, _ := logic.NewSysUser().GetId(ctx, 1)
	if sysUser == nil {
		panic(errors.New("GetId fail"))
	}

	logic.NewSysUser().Delete(ctx, 1)

	//sysUserRes, _ = logic.NewSysUser().GetList(ctx, nil)
	//if sysUser != nil {
	//	panic(errors.New("GetId fail"))
	//}
}
