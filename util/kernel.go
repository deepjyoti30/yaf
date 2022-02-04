package util

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"runtime"
)

// Extract the kernel version by reading the /proc/version file
func GetKernelVersion() string {
	var procFile = "/proc/version"

	content, err := os.ReadFile(procFile)
	if err == nil {
		return parseProc(string(content))
	}

	// Use OS specific methods to extract kernel
	var kernelVersion = ""
	var kernelErr error
	switch runtime.GOOS {
	case "darwin":
		kernelVersion, kernelErr = detectDarwinKernel()
	default:
		kernelErr = errors.New(fmt.Sprint("no OS specific method to get kernel for: ", runtime.GOOS))
	}

	fmt.Println("couldn't read kernel version: ", kernelErr)
	return kernelVersion
}

func parseProc(content string) string {
	re := regexp.MustCompile(`version\s(.*?)\s`)
	cleanRe := regexp.MustCompile(`version|\s`)

	versionExtracted := re.FindAllString(content, 1)
	return cleanRe.ReplaceAllString(versionExtracted[0], "")
}

// detectDarwinKernel detects the kernel version in darwin
// OS. This is only called if the generic way to extract kernel
// fails.
func detectDarwinKernel() (string, error) {
	return "", nil
}
