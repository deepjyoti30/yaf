package main

import (
	"fmt"
	"strings"
	"unicode/utf8"

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

	rightContent := generateRightContent(details)

	for _, value := range rightContent {
		fmt.Println(value)
	}
}

// Generate an array of strings to print line by line when
// fetch is called.
func generateRightContent(details map[string]string) []string {
	// First line should be empty
	var lines = make([]string, 1)

	// Username and hostname
	lines = append(lines, fmt.Sprint(details["username"], "@", details["hostname"]))

	// Add a separator line
	lines = append(lines, strings.Repeat("=", utf8.RuneCountInString(lines[1])))

	// Add an empty line
	lines = append(lines, "")

	// Except username and hostname add rest into proper format
	for key, value := range details {
		if key == "username" || key == "hostname" {
			continue
		}

		lines = append(lines, util.FormatKeyValue(key, value))
	}

	return lines
}
