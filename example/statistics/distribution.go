package main


import (
    "fmt"
    "math"
    statistics "github.com/RifanAmrozi/math-library/statistics"
)

func main() {
    lambda := 2.5 // Average rate of events
    fmt.Println("Poisson distribution probabilities for lambda =", lambda)
    fmt.Println("k\tProbability")

    for k := 0; k <= 10; k++ {
        probability := statistics.PoissonProbability(k, lambda)
        fmt.Printf("%d\t%.5f\n", k, probability)
    }

    // Verify the sum of probabilities equals 1
    totalProbability := 0.0
    for k := 0; k <= 10; k++ {
        totalProbability += statistics.PoissonProbability(k, lambda)
    }
    fmt.Println("Total Probability:", totalProbability)
    fmt.Println("Is total probability equal to 1?", math.Abs(totalProbability-1) < 1e-10)
}
