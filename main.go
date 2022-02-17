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
	version       bool
)

func init() {
	flag.StringVar(&separator, "separator", "  ", "Separator to be used between the key and the value")
	flag.StringVar(&keyPrefix, "key-prefix", "▪ ", "Prefix to be set before the key is printed")
	flag.BoolVar(&noColor, "no-color", false, "Disable showing colors in the output")
	flag.StringVar(&excludeFields, "exclude", "username hostname", "Exclude the passed fields from output. Values should be space separated, eg: `disk os`")
	flag.BoolVar(&version, "version", false, "Print current version of yaf installed")
}

func main() {
	// Parse flags
	flag.Parse()

	// Disable the color if flag is passed
	if noColor {
		color.NoColor = true
	}

	// If to show version or not
	if version {
		// Print the version and exit.
		fmt.Printf("yaf %s\n", util.ArgsDefaultValues().Version)
		return
	}

	var details = map[string]util.GetterFunc{
		"username": util.GetUser,
		"hostname": util.GetHostname,
		"os":       util.GetDistroName,
		"kernel":   util.GetKernelVersion,
		"shell":    util.GetShell,
		"uptime":   util.GetUptime,
		"memory":   util.GetMemory,
		"disk":     util.GetDiskUsage,
	}

	// Parse the fields to exclude string
	fieldsToExclude := util.ParseExcludeFields(excludeFields)

	rightContent := util.GenerateContent(details, separator, keyPrefix, fieldsToExclude)

	for _, value := range rightContent {
		fmt.Println(value)
	}
}
