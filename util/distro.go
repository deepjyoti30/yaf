package util

import (
	"os"
	"regexp"
)

// Get the Distro name
// Try to extract it from the lsb_release file
// If it's not there, try the os-release file
// Use the GOOS var as a fallback
func GetDistroName() string {
	var lsbFile = "/etc/lsb-release"
	//var osReleaseFile = "/etc/os-release"

	// Try the lsb file
	content, err := os.ReadFile(lsbFile)
	if err == nil {
		return parseLsbContent(string(content))
	}

	return ""
}

// Parse the LSB content, use regex to find the distro name
// and return it accordingly.
func parseLsbContent(content string) string {
	// Replace all newlines with nothing
	unwantedChar := regexp.MustCompile("\n|\"")
	content = unwantedChar.ReplaceAllString(content, "")

	re := regexp.MustCompile(".*?DISTRIB_DESCRIPTION=")
	cleanedStr := re.ReplaceAllString(content, "")

	return cleanedStr
}
