package statistics

import (
    "math"
	combinatorics "github.com/RifanAmrozi/math-library/combinatorics"
)

// PoissonProbability calculates the probability of observing k events in a fixed interval
// given the average rate lambda (λ) of events occurring in that interval.
func PoissonProbability(k int, lambda float64) float64 {
    if k < 0 {
        return 0
    }
    return (math.Pow(lambda, float64(k)) * math.Exp(-lambda)) / float64(combinatorics.Factorial(k))
}