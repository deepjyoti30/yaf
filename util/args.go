package util

import "strings"

// Store default values for the commandline args
type DefaultValues struct {
	Align   []string
	Version string
}

func ArgsDefaultValues() DefaultValues {
	return DefaultValues{
		Align: []string{
			0: "center",
			1: "left",
			2: "right",
		},
		Version: "v0.0.3",
	}
}

// Parse the exclude fields string to an array of string
func ParseExcludeFields(excludeFields string) []string {
	return strings.Split(excludeFields, " ")
}
