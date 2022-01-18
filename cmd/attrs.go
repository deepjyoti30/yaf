package piratefetch

import (
	"fmt"
	"os"
)

// Get the name of the user by reading the $USER variable
func getUser() string {
	var username = os.Getenv("USER")
	return username
}

// Get the hostname by reading it from the /etc/hostname file
func getHostname() string {
	var hostnameFile = "/etc/hostname"

	content, err := os.ReadFile(hostnameFile)
	if err != nil {
		fmt.Println("error occurred while reading hostname")
		return ""
	}

	return string(content)
}
