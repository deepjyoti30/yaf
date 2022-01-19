package util

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Get the name of the user by reading the $USER variable
func GetUser() string {
	var username = os.Getenv("USER")
	return username
}

// Get the hostname by reading it from the /etc/hostname file
// If the file is not accessible try using Go's os to get the
// hostname as a fallback.
func GetHostname() string {
	var hostnameFile = "/etc/hostname"

	content, err := os.ReadFile(hostnameFile)
	if err == nil {
		return string(content)
	}

	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println("error reading hostname", err)
		return ""
	}

	return hostname
}

// Get the shell by reading the $SHELL variable
func GetShell() string {
	var shell = os.Getenv("SHELL")

	// Split the string on / and return the last item
	splittedContent := strings.Split(shell, "/")
	lenSplitted := len(splittedContent)

	// Make sure string is not empty
	if lenSplitted < 1 {
		fmt.Println("error occurred while reading shell")
		return "not found"
	}

	// Return the last item
	return splittedContent[lenSplitted-1]
}

// Get the uptime of the system based on the /proc/uptime file
func GetUptime() string {
	var procFile = "/proc/uptime"

	content, err := os.ReadFile(procFile)
	if err != nil {
		fmt.Println("error occurred while reading uptime file")
		return ""
	}

	// Find the uptime in seconds
	splittedContent := strings.Split(string(content), " ")
	uptimeParsed, parseErr := strconv.ParseFloat(splittedContent[0], 64)
	if parseErr != nil {
		fmt.Println("error occurred while parsing uptime string to float")
		return ""
	}

	// Convert to integer
	uptime := int(uptimeParsed)

	hours := uptime / (60 * 60)
	minutes := (uptime % (60 * 60)) / 60

	return fmt.Sprint(hours, "h ", minutes, "m")
}
