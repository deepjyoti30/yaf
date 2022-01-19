package util

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

// Extract the memory details of the system.
// The memory will be returned in the format
// <percent used>% of <total in GB>
func GetMemory() string {
	memFile := "/proc/meminfo"

	content, err := os.ReadFile(memFile)
	if err != nil {
		fmt.Println("couldn't read the memory file")
		return ""
	}

	return parseMemory(string(content))
}

// Parse the memory file to extract the available
// total memory of the system.
func parseMemory(content string) string {
	// Clean up the content to remove stuff like \n
	content = replaceSpecialChars(content, "\n|\\s")

	memAvailable := extractMemFieldIntoInt(content, "MemAvailable")
	memTotal := extractMemFieldIntoInt(content, "MemTotal")

	memUsedPercentage := int((float64(memTotal-memAvailable) / float64(memTotal)) * 100)
	return fmt.Sprint(memUsedPercentage, "% of ", (memTotal / 1000), "MB")
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
