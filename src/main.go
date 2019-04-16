package main

import (
	"fmt"
	"os"
)

func main() {
	var params *Params = ParseArgv(os.Args[1:])

	if params.study == Density {
		for i := 0.0; i <= 200; i++ {
			fmt.Printf("%d %.5f\n", int(i), NormDistrib(params.u, params.s, i))
		}
	} else if params.study == PercentageBelow {
		fmt.Printf("%.1f%% of people have an IQ inferior to %d\n",
			CumulNormDistrib(params.u, params.s, params.iq1)*100-
				CumulNormDistrib(params.u, params.s, 0)*100,
			int(params.iq1))
	} else if params.study == PercentageBetween {
		fmt.Printf("%.1f%% of people have an IQ between %d and %d\n",
			CumulNormDistrib(params.u, params.s, params.iq2)*100-
				CumulNormDistrib(params.u, params.s, params.iq1)*100,
			int(params.iq1),
			int(params.iq2))
	}
}
