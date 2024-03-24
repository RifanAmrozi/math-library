package linear_algebra

import (
    "errors"
    "math"
)

// NormalizeGramSchmidt performs normalization of vectors using the Gram-Schmidt process.
// It takes a slice of vectors represented as slices of float64 values, and returns the normalized vectors.
// If any vector has zero length, it returns an error.
func NormalizeGramSchmidt(vectors [][]float64) ([][]float64, error) {
    if len(vectors) == 0 {
        return nil, errors.New("no vectors provided")
    }

    n := len(vectors[0])
    normalized := make([][]float64, len(vectors))

    for i, v := range vectors {
        if len(v) != n {
            return nil, errors.New("vectors must have the same length")
        }
        if isZeroVector(v) {
            return nil, errors.New("zero-length vector found")
        }

        normalized[i] = make([]float64, n)
        copy(normalized[i], v)

        // Orthogonalize against previous vectors
        for j := 0; j < i; j++ {
            proj := projection(normalized[j], normalized[i])
            for k := 0; k < n; k++ {
                normalized[i][k] -= proj[k]
            }
        }

        // Normalize the resulting vector
        magnitude := math.Sqrt(dotProduct(normalized[i], normalized[i]))
        for j := 0; j < n; j++ {
            normalized[i][j] /= magnitude
        }
    }

    return normalized, nil
}

// isZeroVector checks if a vector has zero length.
func isZeroVector(vector []float64) bool {
    for _, v := range vector {
        if v != 0 {
            return false
        }
    }
    return true
}

// dotProduct calculates the dot product of two vectors.
func dotProduct(a, b []float64) float64 {
    result := 0.0
    for i := 0; i < len(a); i++ {
        result += a[i] * b[i]
    }
    return result
}

// projection calculates the projection of vector 'a' onto vector 'b'.
func projection(a, b []float64) []float64 {
    dot := dotProduct(a, b)
    dotOverDot := dot / dotProduct(b, b)
    projection := make([]float64, len(b))
    for i := range projection {
        projection[i] = dotOverDot * b[i]
    }
    return projection
}
