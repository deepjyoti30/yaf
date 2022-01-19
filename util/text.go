package util

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

// Format the key value into one string such that the key
// string should be of length 6 (since memory is the longest word)
// with 2 spaces in between the key and value.
func FormatKeyValue(key string, value string, separator string) string {
	lengthKey := utf8.RuneCountInString(key)

	if lengthKey < 6 {
		// Add the remaining chars by spaces
		key += strings.Repeat(" ", (6 - lengthKey))
	}

	// Add two spaces in between
	return fmt.Sprint("▪ ", key, separator, value)
}

// Generate an array of strings to print line by line when
// fetch is called.
func GenerateContent(details map[string]string, separator string) []string {
	// First line should be empty
	var lines = make([]string, 1)

	// Username and hostname
	lines = append(lines, fmt.Sprint(details["username"], "@", details["hostname"]))

	// Add a separator line
	lines = append(lines, strings.Repeat("━", utf8.RuneCountInString(lines[1])))

	// Add an empty line
	lines = append(lines, "")

	// Except username and hostname add rest into proper format
	for key, value := range details {
		if key == "username" || key == "hostname" {
			continue
		}

		lines = append(lines, FormatKeyValue(key, value, separator))
	}

	return lines
}
