package util

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

// Format the key value into one string such that the key
// string should be of length 6 (since memory is the longest word)
// with 2 spaces in between the key and value.
func FormatKeyValue(key string, value string) string {
	lengthKey := utf8.RuneCountInString(key)

	if lengthKey < 6 {
		// Add the remaining chars by spaces plus two chars
		// for separation
		key += strings.Repeat(" ", 6-lengthKey+2)
	}

	return fmt.Sprint(key, value)
}
