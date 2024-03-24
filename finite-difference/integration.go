package finite_difference

// TrapezoidalRule calculates the area under the curve using the trapezoidal rule.
// It takes slices of x and y values representing discrete data points and returns the calculated area.
func TrapezoidalRule(x, y []float64) float64 {
    n := len(x)
    if n != len(y) || n < 2 {
        return 0
    }

    area := 0.0
    for i := 1; i < n; i++ {
        h := x[i] - x[i-1]
        area += (y[i] + y[i-1]) * h / 2.0
    }

    return area
}

// TriangleMethod calculates the area under the curve using the triangle (midpoint) method.
// It takes slices of x and y values representing discrete data points and returns the calculated area.
func TriangleMethod(x, y []float64) float64 {
    n := len(x)
    if n != len(y) || n < 2 {
        return 0
    }

    area := 0.0
    for i := 1; i < n; i++ {
        h := x[i] - x[i-1]
        area += y[i-1] * h
    }

    return area
}
