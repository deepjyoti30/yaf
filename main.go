package main

import (
	"flag"
	"fmt"

	"github.com/fatih/color"

	"github.com/deepjyoti30/yaf/util"
)

// Declare vars
var (
	separator     string
	keyPrefix     string
	noColor       bool
	excludeFields string
)

func init() {
	flag.StringVar(&separator, "separator", "  ", "Separator to be used between the key and the value")
	flag.StringVar(&keyPrefix, "key-prefix", "▪ ", "Prefix to be set before the key is printed")
	flag.BoolVar(&noColor, "no-color", false, "Disable showing colors in the output")
	flag.StringVar(&excludeFields, "exclude", "", "Exclude the passed fields from output. Values should be space separated, eg: `disk os`")
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
		"disk":     util.GetDiskUsage(),
	}

	// Parse the fields to exclude string
	fieldsToExclude := util.ParseExcludeFields(excludeFields)

	rightContent := util.GenerateContent(details, separator, keyPrefix, fieldsToExclude)

	for _, value := range rightContent {
		fmt.Println(value)
	}
}
