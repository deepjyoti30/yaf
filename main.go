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
	}
	fmt.Println(details)
}
