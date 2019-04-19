package main

import (
	"fmt"
	"os"
)

func studyDensity(u float64, s float64) {
	for i := 0.0; i <= 200; i++ {
		fmt.Printf("%d %.5f\n", int(i), NormDistrib(u, s, i))
	}
}

func studyIQBelow(u float64, s float64, iq float64) {
	fmt.Printf("%.1f%% of people have an IQ inferior to %d\n",
		CumulNormDistrib(u, s, iq)*100-
			CumulNormDistrib(u, s, 0)*100,
		int(iq))
}

func studyIQBetween(u float64, s float64, iq1 float64, iq2 float64) {
	fmt.Printf("%.1f%% of people have an IQ between %d and %d\n",
		CumulNormDistrib(u, s, iq2)*100-
			CumulNormDistrib(u, s, iq1)*100,
		int(iq1),
		int(iq2))
}

func main() {
	var params *Params = ParseArgv(os.Args[1:])

	switch params.study {
	case Density:
		studyDensity(params.u, params.s)
	case PercentageBelow:
		studyIQBelow(params.u, params.s, params.iq1)
	case PercentageBetween:
		studyIQBetween(params.u, params.s, params.iq1, params.iq2)
	}
}
