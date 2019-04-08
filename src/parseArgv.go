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

func usage() {
	fmt.Println("USAGE")
	fmt.Println("\t./205IQ u s [IQ1] [IQ2]")
	fmt.Println()
	fmt.Println("DESCRIPTION")
	fmt.Println("\tu\t\tmean")
	fmt.Println("\ts\t\tstandard deviation")
	fmt.Println("\tIQ1\t\tminimum IQ")
	fmt.Println("\tIQ2\t\tmaximum IQ")
}

// GetParameters tests if parameters are of the right
// type and transforms them in unsigned ints
func GetParameters(argv []string, argConv []*uint) {
	ruint, _ := regexp.Compile("^[1-9]\\d*$")
	argNames := []string{"mean", "standard deviation", "minimum IQ", "maximum IQ"}

	for ndx, arg := range argv {
		if !ruint.MatchString(arg) {
			fmt.Fprintf(os.Stderr, "%s: Invalid %s\n", arg, argNames[ndx])
			os.Exit(FAILURE)
		}
		conv, _ := strconv.ParseUint(arg, 10, 64)
		*argConv[ndx] = uint(conv)
	}
}

// ParseArgv parses the arguments of the program and returns a constant
func ParseArgv(argv []string) (u uint, s uint, iq1 uint, iq2 uint) {
	iq1 = 0
	iq2 = 200

	if len(argv) == 1 && (argv[0] == "-h" || argv[0] == "--help") {
		usage()
		os.Exit(SUCCESS)
	} else if len(argv) < 2 || len(argv) > 4 {
		usage()
		os.Exit(FAILURE)
	}

	GetParameters(argv, []*uint{&u, &s, &iq1, &iq2})

	// if !r.MatchString(argv[0]) {
	// 	fmt.Fprint(os.Stderr, argv[0], ": Invalid constant\n")
	// 	os.Exit(FAILURE)
	// }

	// a, _ := strconv.ParseFloat(argv[0], 64)

	// if a < 0 || a > 2.5 {
	// 	fmt.Fprint(os.Stderr, argv[0], ": Invalid constant\n")
	// 	os.Exit(FAILURE)
	// }
	return u, s, iq1, iq2
}
