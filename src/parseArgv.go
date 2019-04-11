package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

// FAILURE is the constant representing the failure of the program
const FAILURE int = 84

// SUCCESS is the constant representing the success of the program
const SUCCESS int = 0

// usage prints the usage of the program
func usage() {
	fmt.Println("USAGE")
	fmt.Printf("\t%s u s [IQ1] [IQ2]\n", os.Args[0])
	fmt.Println()
	fmt.Println("DESCRIPTION")
	fmt.Println("\tu\t\tmean")
	fmt.Println("\ts\t\tstandard deviation")
	fmt.Println("\tIQ1\t\tminimum IQ")
	fmt.Println("\tIQ2\t\tmaximum IQ")
}

// determinStudy determins the kind of study the programs must do
func determinStudy(argv []string, params *Params) {
	switch len(argv) {
	case 2:
		params.study = Density
	case 3:
		params.study = PercentageBelow
	case 4:
		params.study = PercentageBetween
	default:
		usage()
		os.Exit(FAILURE)
	}
}

// getParameters tests if parameters are of the right
// type and transforms them in unsigned ints
func getParameters(argv []string, params *Params) {
	ruint, _ := regexp.Compile("^[1-9]\\d*$")
	argNames := []string{"mean", "standard deviation", "minimum IQ", "maximum IQ"}
	argConv := []*float64{&params.u, &params.s, &params.iq1, &params.iq2}

	for ndx, arg := range argv {
		if !ruint.MatchString(arg) {
			fmt.Fprintf(os.Stderr, "%s: Invalid %s\n", arg, argNames[ndx])
			os.Exit(FAILURE)
		}
		conv, _ := strconv.ParseUint(arg, 10, 64)
		*argConv[ndx] = float64(conv)
	}
}

// checkValues checks if the values given as argument make sense
func checkValues(params *Params) {
	if params.study == PercentageBetween && params.iq1 >= params.iq2 {
		fmt.Fprintln(os.Stderr, "Error: IQ minimum superior or equal to IQ maximum")
		os.Exit(FAILURE)
	}
}

// ParseArgv parses the arguments of the program and returns a constant
func ParseArgv(argv []string) *Params {
	var params = &Params{u: 0, s: 0, iq1: 0, iq2: 200, study: Undefined}

	if len(argv) == 1 && (argv[0] == "-h" || argv[0] == "--help") {
		usage()
		os.Exit(SUCCESS)
	}

	determinStudy(argv, params)
	getParameters(argv, params)
	checkValues(params)

	return params
}
