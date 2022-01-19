package util

// Store default values for the commandline args
type DefaultValues struct {
	Align []string
}

func ArgsDefaultValues() DefaultValues {
	return DefaultValues{
		Align: []string{
			0: "center",
			1: "left",
			2: "right",
		},
	}
}
