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
	var osReleaseFile = "/etc/os-release"

	// Try the lsb file
	content, err := os.ReadFile(lsbFile)
	if err == nil {
		return parseLsbContent(string(content))
	}

	// Try the Os release file
	content, err = os.ReadFile(osReleaseFile)
	if err == nil {
		return parseOsReleaseContent(string(content))
	}

	return ""
}

// Parse the LSB content, use regex to find the distro name
// and return it accordingly.
func parseLsbContent(content string) string {
	// Replace all newlines and " with nothing
	content = replaceSpecialChars(content)

	re := regexp.MustCompile(".*?DISTRIB_DESCRIPTION=")
	cleanedStr := re.ReplaceAllString(content, "")

	return cleanedStr
}

// Parse the OS Release content, use regex to find the distro name
// and return it accordingly.
func parseOsReleaseContent(content string) string {
	// Clean the data
	content = replaceSpecialChars(content)

	// Remove the content before the PRETTY_NAME field
	beforeMatch := regexp.MustCompile(".*?PRETTY_NAME=")
	content = beforeMatch.ReplaceAllString(content, "")

	// Remove the content after the PRETTY_NAME.
	// The ID field comes right after PRETTY_NAME
	afterMatch := regexp.MustCompile("ID.*?$")
	content = afterMatch.ReplaceAllString(content, "")

	return content
}

// Replace special chars to make the text clean
func replaceSpecialChars(content string) string {
	unwantedChar := regexp.MustCompile("\n|\"")
	content = unwantedChar.ReplaceAllString(content, "")
	return content
}
