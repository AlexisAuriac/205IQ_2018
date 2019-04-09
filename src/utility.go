package main

import "math"

// NormDistrib computes the normal distribution given u, s and x
func NormDistrib(u uint, s uint, x float64) float64 {
	var uf = float64(u)
	var sf = float64(s)

	return (math.Exp(-(math.Pow(x-uf, 2) / (2 * math.Pow(sf, 2))))) /
		(sf * math.Sqrt(2*math.Pi))
}
