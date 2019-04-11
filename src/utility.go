package main

import "math"

// NormDistrib computes the normal distribution given u, s and x
func NormDistrib(u float64, s float64, x float64) float64 {
	return (math.Exp(-(math.Pow(x-u, 2) / (2 * math.Pow(s, 2))))) /
		(s * math.Sqrt(2*math.Pi))
}

// CumulNormDistrib computes the cumulative normal distribution
func CumulNormDistrib(u float64, s float64, x float64) float64 {
	return (1 + math.Erf((x-u)/(s*math.Sqrt(2)))) / 2
}
