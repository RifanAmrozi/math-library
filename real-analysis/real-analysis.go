package real_analysis

import "math"

// RiemannSumLeftEndpoint calculates the Riemann sum using left endpoints.
// It takes a function f, a lower limit a, an upper limit b, and the number of subintervals n.
// It returns the approximation of the definite integral of f over [a, b].
func RiemannSumLeftEndpoint(f func(float64) float64, a, b float64, n int) float64 {
    delta := (b - a) / float64(n)
    sum := 0.0
    for i := 0; i < n; i++ {
        x := a + float64(i)*delta
        sum += f(x)
    }
    return delta * sum
}

// RiemannSumRightEndpoint calculates the Riemann sum using right endpoints.
// It takes a function f, a lower limit a, an upper limit b, and the number of subintervals n.
// It returns the approximation of the definite integral of f over [a, b].
func RiemannSumRightEndpoint(f func(float64) float64, a, b float64, n int) float64 {
    delta := (b - a) / float64(n)
    sum := 0.0
    for i := 1; i <= n; i++ {
        x := a + float64(i)*delta
        sum += f(x)
    }
    return delta * sum
}