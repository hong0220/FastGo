package gen

import (
	"fmt"
	"github.com/gogf/gf/v2/test/gtest"
	"github.com/hong0220/FastGo/pkg/common/base"
	"io/ioutil"
	"strings"
	"testing"
)

func TestGenApi(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		inFilePath := base.GetProjectPath() + "/manifest/gen/api.txt"
		data, err := ioutil.ReadFile(inFilePath)
		if err != nil {
			fmt.Println(err)
		}
		//fmt.Print(string(data))

		content := strings.Replace(string(data), "{entityName}", "SysDict", -1)
		content = strings.Replace(content, "{entityPath}", "sysDict", -1)
		fmt.Print(content)

		outFilePath := base.GetProjectPath() + "/internal/api/sys_dict.go"
		err = ioutil.WriteFile(outFilePath, []byte(content), 0777)
		if err != nil {
			fmt.Println(err)
		}
	})
}

func TestGenLogic(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		inFilePath := base.GetProjectPath() + "/manifest/gen/logic.txt"
		data, err := ioutil.ReadFile(inFilePath)
		if err != nil {
			fmt.Println(err)
		}
		//fmt.Print(string(data))

		content := strings.Replace(string(data), "{entityName}", "SysDict", -1)
		content = strings.Replace(content, "{entityPath}", "sysDict", -1)
		fmt.Print(content)

		outFilePath := base.GetProjectPath() + "/internal/logic/sys/sys_dict.go"
		err = ioutil.WriteFile(outFilePath, []byte(content), 0777)
		if err != nil {
			fmt.Println(err)
		}
	})
}

func TestGenController(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		inFilePath := base.GetProjectPath() + "/manifest/gen/controller.txt"
		data, err := ioutil.ReadFile(inFilePath)
		if err != nil {
			fmt.Println(err)
		}
		//fmt.Print(string(data))

		content := strings.Replace(string(data), "{entityName}", "SysDict", -1)
		fmt.Print(content)

		outFilePath := base.GetProjectPath() + "/internal/controller/sys_dict.go"
		err = ioutil.WriteFile(outFilePath, []byte(content), 0777)
		if err != nil {
			fmt.Println(err)
		}
	})
}
