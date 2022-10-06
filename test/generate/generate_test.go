package generate

import (
	"fmt"
	"github.com/hong0220/FastGo/test"
	"io/ioutil"
	"strings"
	"testing"
)

func TestGenerate(t *testing.T) {
	array := []string{"sys_area", "sys_dict", "sys_log",
		"sys_office", "sys_resource", "sys_role",
		"sys_role_office", "sys_resource", "sys_task",
		"sys_user", "sys_user_role"}
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
		GenerateApi(entityName, entityPath, fileName)
		GenerateLogic(entityName, entityPath, fileName)
		GenerateController(entityName, fileName)
	}
}

func GenerateApi(entityName string, entityPath string, fileName string) {
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

func GenerateLogic(entityName string, entityPath string, fileName string) {
	inFilePath := test.GetProjectPath() + "/manifest/generate/logic.txt"
	data, err := ioutil.ReadFile(inFilePath)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Print(string(data))

	content := strings.Replace(string(data), "{entityName}", entityName, -1)
	content = strings.Replace(content, "{entityPath}", entityPath, -1)
	fmt.Print(content)

	outFilePath := test.GetProjectPath() + "/internal/logic/sys/" + fileName + ".go"
	err = ioutil.WriteFile(outFilePath, []byte(content), 0777)
	if err != nil {
		fmt.Println(err)
	}
}

func GenerateController(entityName string, fileName string) {
	inFilePath := test.GetProjectPath() + "/manifest/generate/controller.txt"
	data, err := ioutil.ReadFile(inFilePath)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Print(string(data))

	content := strings.Replace(string(data), "{entityName}", entityName, -1)
	fmt.Print(content)

	outFilePath := test.GetProjectPath() + "/internal/controller/" + fileName + ".go"
	err = ioutil.WriteFile(outFilePath, []byte(content), 0777)
	if err != nil {
		fmt.Println(err)
	}
}
