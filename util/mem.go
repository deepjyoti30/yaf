package util

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"runtime"
	"strconv"
)

// Extract the memory details of the system.
// The memory will be returned in the format
// <percent used>% of <total in GB>
func GetMemory() string {
	memFile := "/proc/meminfo"

	content, err := os.ReadFile(memFile)
	if err == nil {
		return parseMemory(string(content))
	}

	var memErr error
	var memUsage string = ""
	switch runtime.GOOS {
	case "darwin":
		memUsage, memErr = detectDarwinMem()
	default:
		memErr = errors.New(fmt.Sprint("no OS specific method to get mem for: ", runtime.GOOS))
	}

	if memErr != nil {
		fmt.Println("couldn't read the memory file: ", memErr)
		return ""
	}

	return memUsage
}

// Parse the memory file to extract the available
// total memory of the system.
func parseMemory(content string) string {
	// Clean up the content to remove stuff like \n
	content = replaceSpecialChars(content, "\n|\\s")

	memAvailable := extractMemFieldIntoInt(content, "MemAvailable")
	memTotal := extractMemFieldIntoInt(content, "MemTotal")

	memUsedPercentage := int((float64(memTotal-memAvailable) / float64(memTotal)) * 100)
	return fmt.Sprint(memUsedPercentage, "% of ", (memTotal / 1000), "m")
}

// Extract the passed field into integer from the memory
// files content.
func extractMemFieldIntoInt(content string, field string) int {
	re := regexp.MustCompile(field + `:.*?kB`)
	matches := re.FindAllString(content, 1)

	if len(matches) < 1 {
		fmt.Println("error occurred while matching:", field)
		return 1
	}

	// Clean it up a bit to get the number
	cleanedMatch := replaceSpecialChars(matches[0], (field + `|:|kB`))

	// Convert to integer
	amountInKb, err := strconv.Atoi(cleanedMatch)
	if err != nil {
		fmt.Println("error occurred while converting", field, "to integer:", err)
		return 1
	}
	return amountInKb
}

// detectDarwinMem will detect the memory usage in Darwin
// OS and return a string that can be showed to the user.
func detectDarwinMem() (string, error) {
	// Detect the totalMem
	outputStr, err := runSysctlCmd("-n hw.memsize")
	if err != nil {
		return "", err
	}

	// Remove newlines
	outputStr = replaceSpecialChars(outputStr, "\n")

	// It will be an integer in bytes directly, we can convert to
	// int directly
	totalMemBytes, err := strconv.Atoi(outputStr)
	if err != nil {
		return "", err
	}

	// Convert the totalMem to MB
	totalMem := totalMemBytes / 1024 / 1024

	// Extract the used percentage using the memory pressure command
	memPressureCmd := "memory_pressure"
	cmd := exec.Command(memPressureCmd)
	stdout, err := cmd.Output()
	if err != nil {
		return "", err
	}

	memPressureStr := string(stdout)

	// Remove all newlines
	memPressureStr = replaceSpecialChars(memPressureStr, "\n")
	extractFreePercentRe := regexp.MustCompile(`.*?System-wide memory free percentage:\s|%`)
	memFreePercentageStr := extractFreePercentRe.ReplaceAllString(memPressureStr, "")
	memFreePercentage, err := strconv.Atoi(memFreePercentageStr)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%d of %dm", 100-memFreePercentage, totalMem), nil
}
