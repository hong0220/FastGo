package service

import (
	"fmt"
	"github.com/gogf/gf/v2/os/gctx"
	logic "github.com/hong0220/FastGo/internal/logic/sys"
	"github.com/hong0220/FastGo/pkg/common/base"
	"testing"
)

func TestGetId(t *testing.T) {
	base.ApiConfigSetUp()

	ctx := gctx.New()
	fmt.Println(logic.NewSysUser().GetId(ctx, 2))
}
