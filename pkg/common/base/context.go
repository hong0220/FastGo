package base

import (
	"os"
	"strings"
)

func GetProjectPath() string {
	projectName := "FastGo"
	path, _ := os.Getwd()
	path = path[0 : strings.Index(path, projectName)+len(projectName)]
	return path
}
