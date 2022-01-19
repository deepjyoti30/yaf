package main

import (
	"fmt"

	"github.com/deepjyoti30/piratefetch/util"
)

func main() {
	var details = map[string]string{
		"username": util.GetUser(),
		"hostname": util.GetHostname(),
		"distro":   util.GetDistroName(),
		"kernel":   util.GetKernalVersion(),
		"shell":    util.GetShell(),
		"uptime":   util.GetUptime(),
		"memory":   util.GetMemory(),
	}
	fmt.Println(details)
	fmt.Println(util.GetPirateAscii())
}
