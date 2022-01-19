package main

import (
	"fmt"

	"github.com/deepjyoti30/yaf/util"
)

func main() {
	var details = map[string]string{
		"username": util.GetUser(),
		"hostname": util.GetHostname(),
		"os":       util.GetDistroName(),
		"kernel":   util.GetKernalVersion(),
		"shell":    util.GetShell(),
		"uptime":   util.GetUptime(),
		"memory":   util.GetMemory(),
	}

	rightContent := util.GenerateContent(details)

	for _, value := range rightContent {
		fmt.Println(value)
	}
}
