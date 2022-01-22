package util

import (
	"fmt"
	"strings"
	"unicode/utf8"

	"github.com/fatih/color"
)

var (
	boldGreen  func(a ...interface{}) string = color.New(color.Bold, color.FgGreen).SprintFunc()
	boldRed    func(a ...interface{}) string = color.New(color.FgRed).SprintFunc()
	boldYellow func(a ...interface{}) string = color.New(color.Bold, color.FgYellow).SprintFunc()
	green      func(a ...interface{}) string = color.New(color.FgGreen).SprintFunc()
)

// Format the key value into one string such that the key
// string should be of length 6 (since memory is the longest word)
// with 2 spaces in between the key and value.
func FormatKeyValue(key string, value string, separator string, keyPrefix string) string {
	lengthKey := utf8.RuneCountInString(key)

	if lengthKey < 6 {
		// Add the remaining chars by spaces
		key += strings.Repeat(" ", (6 - lengthKey))
	}

	// Add two spaces in between
	return fmt.Sprint(boldRed(keyPrefix), boldYellow(key), separator, green(value))
}

// Generate an array of strings to print line by line when
// fetch is called.
func GenerateContent(details map[string]string, separator string, keyPrefix string, fieldsToExclude []string) []string {
	// First line should be empty
	var lines = make([]string, 1)

	// Username and hostname
	lines = append(lines, fmt.Sprint(boldGreen(details["username"]), boldYellow("@"), boldGreen(details["hostname"])))

	// Add a separator line
	// Add 1 for the length of the `@`
	separatorCount := utf8.RuneCountInString(details["username"]) + utf8.RuneCountInString(details["hostname"]) + 1
	lines = append(lines, boldRed(strings.Repeat("━", separatorCount)))

	// Add an empty line
	lines = append(lines, "")

	// NOTE: Username and hostname will be removed from the exclude fields
	// array.

	// Delete the fields to exclude from the map
	for _, fieldToExclude := range fieldsToExclude {
		delete(details, fieldToExclude)
	}

	// Except username and hostname add rest into proper format
	for key, value := range details {
		lines = append(lines, FormatKeyValue(key, value, separator, keyPrefix))
	}

	// Add empty line at the end
	lines = append(lines, "")

	return lines
}
