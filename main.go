package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/deepjyoti30/yaf/util"
)

// Declare vars
var (
	separator string
	align     string
)

func init() {
	flag.StringVar(&separator, "separator", "  ", "Separator to be used between the key and the value")
	flag.StringVar(&align, "align", "left", fmt.Sprint("Alignment of the content. Allowed values are: ", strings.Join(util.ArgsDefaultValues().Align, ", ")))
}

func main() {
	// Parse flags
	flag.Parse()

	var details = map[string]string{
		"username": util.GetUser(),
		"hostname": util.GetHostname(),
		"os":       util.GetDistroName(),
		"kernel":   util.GetKernalVersion(),
		"shell":    util.GetShell(),
		"uptime":   util.GetUptime(),
		"memory":   util.GetMemory(),
	}

	rightContent := util.GenerateContent(details, separator)

	for _, value := range rightContent {
		fmt.Println(value)
	}
}
