package main

import (
	"fmt"
	"os"
)

func main() {
	var params *Params = ParseArgv(os.Args[1:])

	if params.study == Density {
		for i := uint(0); i < 200; i++ {
			fmt.Printf("%d %.5f\n", i, NormDistrib(params.u, params.s, float64(i)))
		}
	} else if params.study == PercentageBelow {
		var res = 0.0
		for i := float64(0); i < float64(params.iq1); i += 0.01 {
			res += NormDistrib(params.u, params.s, i)
		}
		fmt.Printf("%.1f%% of people have an IQ inferior to %d\n",
			res,
			params.iq1)
	} else if params.study == PercentageBetween {
		var res = 0.0
		for i := float64(params.iq1); i < float64(params.iq2); i += 0.01 {
			res += NormDistrib(params.u, params.s, i)
		}
		fmt.Printf("%.1f%% of people have an IQ between %d and %d\n",
			res,
			params.iq1,
			params.iq2)
	}
}
