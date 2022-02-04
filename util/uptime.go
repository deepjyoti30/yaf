package util

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"time"
)

// Get the uptime of the system based on the /proc/uptime file
func GetUptime() string {
	var procFile = "/proc/uptime"

	content, err := os.ReadFile(procFile)
	if err == nil {
		return parseUptimeFile(string(content))
	}

	// Try the fallbacks depending on Go detected OS
	var nativeUptimeErr error
	var nativeUptime string
	switch runtime.GOOS {
	case "darwin":
		nativeUptime, nativeUptimeErr = detectDarwinUptime()
	default:
		nativeUptime = ""
		nativeUptimeErr = errors.New(fmt.Sprint("no uptime config defined for: ", runtime.GOOS))
	}

	if nativeUptimeErr != nil {
		fmt.Println("couldn't find uptime for system: ", nativeUptimeErr)
		return ""
	}

	return nativeUptime
}

// parseUptimeFile parses the output of the /proc/uptime file
// to make it readable
func parseUptimeFile(content string) string {
	// Find the uptime in seconds
	splittedContent := strings.Split(string(content), " ")
	uptimeParsed, parseErr := strconv.ParseFloat(splittedContent[0], 64)
	if parseErr != nil {
		fmt.Println("error occurred while parsing uptime string to float")
		return ""
	}

	// Convert to integer
	uptime := int(uptimeParsed)
	return makeDurationReadable(uptime)
}

// detectDarwinUptime detects the uptime in Darwin OS
// by using the sysctl cmd and parsing the string.
func detectDarwinUptime() (string, error) {
	// Run the sysctl command to get the uptime string
	sysctlCmdStr := "sysctl"
	args := "-n kern.boottime"
	cmd := exec.Command(sysctlCmdStr, strings.Split(args, " ")...)
	stdout, err := cmd.Output()

	if err != nil {
		return "", err
	}

	outputInStr := string(stdout)

	extractBootTime := regexp.MustCompile(`=\s\d+,`)
	cleanRe := regexp.MustCompile(`=|\s|,`)

	matches := extractBootTime.FindAllString(outputInStr, 1)

	if len(matches) < 1 {
		return "", errors.New("error while parsing sysctl output string")
	}

	cleanedMatch := cleanRe.ReplaceAllString(matches[0], "")

	bootTimeInt, err := strconv.ParseInt(cleanedMatch, 10, 64)
	if err != nil {
		return "", err
	}

	bootTime := time.Unix(bootTimeInt, 0)
	duration := time.Since(bootTime)

	return makeDurationReadable(int(duration.Seconds())), nil

}

// makeDurationReadable converts the passed hours and minutes to
// make them human readable
func makeDurationReadable(uptime int) string {
	indexToTimeMap := map[string]int{
		"M": 0,
		"d": 0,
		"h": 0,
		"m": 0,
	}

	// Convert seconds to minutes
	indexToTimeMap["m"] = uptime / 60
	if indexToTimeMap["m"] > 60 {
		indexToTimeMap["h"] = indexToTimeMap["m"] / 60
		indexToTimeMap["m"] = indexToTimeMap["m"] % 60
	}

	if indexToTimeMap["h"] > 24 {
		indexToTimeMap["d"] = indexToTimeMap["h"] / 24
		indexToTimeMap["h"] = indexToTimeMap["h"] % 24
	}

	if indexToTimeMap["d"] > 30 {
		indexToTimeMap["M"] = indexToTimeMap["d"] / 30
		indexToTimeMap["d"] = indexToTimeMap["d"] % 30
	}

	const maxDuration = 2
	var count = 0
	durationArr := make([]string, 0)

	for timeStr, time := range indexToTimeMap {
		if count == maxDuration {
			break
		}

		if time == 0 {
			continue
		}

		durationArr = append(durationArr, fmt.Sprintf("%d%s", time, timeStr))

		count += 1
	}

	return strings.Join(durationArr, " ")
}
