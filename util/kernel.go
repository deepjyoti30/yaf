package util

import (
	"fmt"
	"os"
	"regexp"
)

// Extract the kernel version by reading the /proc/version file
func GetKernelVersion() string {
	var procFile = "/proc/version"

	content, err := os.ReadFile(procFile)
	if err != nil {
		fmt.Println("couldn't read kernel version")
		return "not found"
	}

	return parseProc(string(content))
}

func parseProc(content string) string {
	re := regexp.MustCompile(`version\s(.*?)\s`)
	cleanRe := regexp.MustCompile(`version|\s`)

	versionExtracted := re.FindAllString(content, 1)
	return cleanRe.ReplaceAllString(versionExtracted[0], "")
}
