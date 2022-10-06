package file

import (
	"fmt"
	"os"
)

// 判断文件夹是否存在
func HasDir(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	// 排除文件不存在
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// 创建文件夹
func CreateDir(path string) {
	exist, err := HasDir(path)
	if err != nil {
		fmt.Printf("获取文件夹异常,err=%v\n", err)
		return
	}
	if exist {
		fmt.Println("文件夹已存在!")
	} else {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			fmt.Printf("创建目录异常,err=%v\n", err)
		} else {
			fmt.Println("创建成功!")
		}
	}
}
