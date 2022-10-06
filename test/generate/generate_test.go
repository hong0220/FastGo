package generate

import (
	"fmt"
	"github.com/hong0220/FastGo/pkg/common/file"
	"github.com/hong0220/FastGo/pkg/common/test"
	"io/ioutil"
	"strings"
	"testing"
)

func TestGenerate(t *testing.T) {
	// todo包名

	array := []string{"sys_area", "sys_dict", "sys_log",
		"sys_office", "sys_resource", "sys_role",
		"sys_role_office", "sys_resource", "sys_task",
		"sys_user", "sys_user_role"}

	packageName := "sys"
	for _, item := range array {
		dictArray := strings.Split(item, "_")
		entityName := ""
		entityPath := ""
		if len(dictArray) == 2 {
			entityName = strings.ToUpper(dictArray[0][:1]) + dictArray[0][1:] + strings.ToUpper(dictArray[1][:1]) + dictArray[1][1:]
			entityPath = dictArray[0] + dictArray[1]
		} else if len(dictArray) == 3 {
			entityName = strings.ToUpper(dictArray[0][:1]) + dictArray[0][1:] + strings.ToUpper(dictArray[1][:1]) + dictArray[1][1:] + strings.ToUpper(dictArray[2][:1]) + dictArray[2][1:]
			entityPath = dictArray[0] + dictArray[1] + dictArray[2]
		}
		fileName := item
		GenerateApi(packageName, entityName, entityPath, fileName)
		GenerateLogic(packageName, entityName, entityPath, fileName)
		GenerateController(packageName, entityName, entityPath, fileName)
	}
}

func GenerateApi(packageName string, entityName string, entityPath string, fileName string) {
	inFilePath := test.GetProjectPath() + "/manifest/generate/api.txt"
	data, err := ioutil.ReadFile(inFilePath)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Print(string(data))

	content := strings.Replace(string(data), "{entityName}", entityName, -1)
	content = strings.Replace(content, "{entityPath}", entityPath, -1)
	fmt.Print(content)

	outFilePath := test.GetProjectPath() + "/internal/api/" + fileName + ".go"
	err = ioutil.WriteFile(outFilePath, []byte(content), 0777)
	if err != nil {
		fmt.Println(err)
	}
}

func GenerateLogic(packageName string, entityName string, entityPath string, fileName string) {
	inFilePath := test.GetProjectPath() + "/manifest/generate/logic.txt"
	data, err := ioutil.ReadFile(inFilePath)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Print(string(data))

	content := strings.Replace(string(data), "{entityName}", entityName, -1)
	content = strings.Replace(content, "{entityPath}", entityPath, -1)
	fmt.Print(content)

	filePath := test.GetProjectPath() + "/internal/logic/" + packageName
	file.CreateDir(filePath)
	outFilePath := filePath + "/" + fileName + ".go"
	err = ioutil.WriteFile(outFilePath, []byte(content), 0777)
	if err != nil {
		fmt.Println(err)
	}
}

func GenerateController(packageName string, entityName string, entityPath string, fileName string) {
	inFilePath := test.GetProjectPath() + "/manifest/generate/controller.txt"
	data, err := ioutil.ReadFile(inFilePath)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Print(string(data))

	content := strings.Replace(string(data), "{entityName}", entityName, -1)
	fmt.Print(content)

	filePath := test.GetProjectPath() + "/internal/controller/" + packageName
	file.CreateDir(filePath)
	outFilePath := filePath + "/" + fileName + ".go"
	err = ioutil.WriteFile(outFilePath, []byte(content), 0777)
	if err != nil {
		fmt.Println(err)
	}
}
