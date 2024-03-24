package combinatorics

import "errors"

// Factorial calculates the factorial of n.
func Factorial(n int) (int) {
    if n < 0 {
        return 1
    }
    if n == 0 {
        return 1
    }
    result := 1
    for i := 1; i <= n; i++ {
        result *= i
    }
    return result
}

// Permutation calculates the number of permutations of k elements taken from a set of n elements.
func Permutation(n, k int) (int, error) {
    if n < 0 || k < 0 || k > n {
        return 0, errors.New("invalid input for permutation")
    }
    if k == 0 {
        return 1, nil
    }
    result := 1
    for i := n; i > n-k; i-- {
        result *= i
    }
    return result, nil
}

// Combination calculates the number of combinations of k elements taken from a set of n elements.
func Combination(n, k int) (int, error) {
    if n < 0 || k < 0 || k > n {
        return 0, errors.New("invalid input for combination")
    }
    if k == 0 || k == n {
        return 1, nil
    }
    if k > n-k {
        k = n - k
    }
    numerator := Factorial(n)
    denominator := Factorial(k)
    denominator *= Factorial(n - k)
    return numerator / denominator, nil
}
