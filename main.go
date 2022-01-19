package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/fatih/color"

	"github.com/deepjyoti30/yaf/util"
)

// Declare vars
var (
	separator string
	keyPrefix string
	align     string
	noColor   bool
)

func init() {
	flag.StringVar(&separator, "separator", "  ", "Separator to be used between the key and the value")
	flag.StringVar(&keyPrefix, "key-prefix", "▪ ", "Prefix to be set before the key is printed")
	flag.BoolVar(&noColor, "no-color", false, "Disable showing colors in the output")
	flag.StringVar(&align, "align", "left", fmt.Sprint("Alignment of the content. Allowed values are: ", strings.Join(util.ArgsDefaultValues().Align, ", ")))
}

func main() {
	// Parse flags
	flag.Parse()

	// Disable the color if flag is passed
	if noColor {
		color.NoColor = true
	}

	var details = map[string]string{
		"username": util.GetUser(),
		"hostname": util.GetHostname(),
		"os":       util.GetDistroName(),
		"kernel":   util.GetKernalVersion(),
		"shell":    util.GetShell(),
		"uptime":   util.GetUptime(),
		"memory":   util.GetMemory(),
	}

	rightContent := util.GenerateContent(details, separator, keyPrefix)

	for _, value := range rightContent {
		fmt.Println(value)
	}
}
