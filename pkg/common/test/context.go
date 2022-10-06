package test

import (
	// 注意这边
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcfg"
	"os"
	"strings"
)

func GetProjectPath() string {
	projectName := "FastGo"
	path, _ := os.Getwd()
	path = path[0 : strings.Index(path, projectName)+len(projectName)]
	return path
}

func ApiConfigSetUp() {
	path := GetProjectPath() + "/manifest/config/config.local.yaml"
	g.Cfg().GetAdapter().(*gcfg.AdapterFile).SetFileName(path)
}
