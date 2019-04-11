package main

// Study is an enum that determins the kind of study the programs must do
type Study int

// Enumeration of studies
const (
	Density           Study = 0
	PercentageBelow   Study = 1
	PercentageBetween Study = 2
	Undefined         Study = 3
)

// Params is a structure containing the parsed argument sent to the program
type Params struct {
	u     float64
	s     float64
	iq1   float64
	iq2   float64
	study Study
}
