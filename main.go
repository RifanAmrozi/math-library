package main

import (
    "fmt"
    integration "finite-difference"
)

func main() {
    x := []float64{0, 1, 2, 3, 4}
    y := []float64{0, 1, 4, 9, 16}

    areaTrapezoidal := integration.TrapezoidalRule(x, y)
    areaTriangle := integration.TriangleMethod(x, y)

    fmt.Println("Area under the curve using trapezoidal rule:", areaTrapezoidal)
    fmt.Println("Area under the curve using triangle (midpoint) method:", areaTriangle)
}
