package util

import (
	"fmt"
	"os"
)

// Get the name of the user by reading the $USER variable
func GetUser() string {
	var username = os.Getenv("USER")
	return username
}

// Get the hostname by reading it from the /etc/hostname file
func GetHostname() string {
	var hostnameFile = "/etc/hostname"

	content, err := os.ReadFile(hostnameFile)
	if err != nil {
		fmt.Println("error occurred while reading hostname")
		return ""
	}

	return string(content)
}
