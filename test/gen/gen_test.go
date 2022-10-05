package gen

import (
	"fmt"
	"github.com/gogf/gf/v2/test/gtest"
	"github.com/hong0220/FastGo/pkg/common/base"
	"io/ioutil"
	"strings"
	"testing"
)

func Test_Gen_Api(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		inFilePath := base.GetProjectPath() + "/manifest/gen/api.txt"
		data, err := ioutil.ReadFile(inFilePath)
		if err != nil {
			fmt.Println(err)
		}
		//fmt.Print(string(data))

		content := strings.Replace(string(data), "{entityName}", "User", -1)
		content = strings.Replace(content, "{entityPath}", "user", -1)
		fmt.Print(content)

		outFilePath := base.GetProjectPath() + "/internal/api/user.go"
		err = ioutil.WriteFile(outFilePath, []byte(content), 0777)
		if err != nil {
			fmt.Println(err)
		}
	})
}

func Test_Gen_Logic(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		inFilePath := base.GetProjectPath() + "/manifest/gen/logic.txt"
		data, err := ioutil.ReadFile(inFilePath)
		if err != nil {
			fmt.Println(err)
		}
		//fmt.Print(string(data))

		content := strings.Replace(string(data), "{entityName}", "User", -1)
		content = strings.Replace(content, "{entityPath}", "user", -1)
		fmt.Print(content)

		outFilePath := base.GetProjectPath() + "/internal/logic/user/user.go"
		err = ioutil.WriteFile(outFilePath, []byte(content), 0777)
		if err != nil {
			fmt.Println(err)
		}
	})
}

func Test_Gen_Controller(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		inFilePath := base.GetProjectPath() + "/manifest/gen/controller.txt"
		data, err := ioutil.ReadFile(inFilePath)
		if err != nil {
			fmt.Println(err)
		}
		//fmt.Print(string(data))

		content := strings.Replace(string(data), "{entityName}", "User", -1)
		fmt.Print(content)

		outFilePath := base.GetProjectPath() + "/internal/controller/user.go"
		err = ioutil.WriteFile(outFilePath, []byte(content), 0777)
		if err != nil {
			fmt.Println(err)
		}
	})
}
