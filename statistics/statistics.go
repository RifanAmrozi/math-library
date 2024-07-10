package statistics

import (
	"math"
	"sort"
)

// CalculateIQR calculates the interquartile range (IQR) of a dataset.
func CalculateIQR(data []float64) float64 {
	sort.Float64s(data)
	q1 := percentile(data, 25)
	q3 := percentile(data, 75)
	return q3 - q1
}

// Percentile calculates the p-th percentile of a sorted dataset.
func percentile(data []float64, p float64) float64 {
	index := (p / 100.0) * float64(len(data)-1)
	lower := int(math.Floor(index))
	upper := int(math.Ceil(index))
	if lower == upper {
		return data[lower]
	}
	return data[lower] + (data[upper]-data[lower])*(index-math.Floor(index))
}

// FreedmanDiaconis calculates the number of bins for a histogram based on the Freedman-Diaconis rule.
func FreedmanDiaconis(data []float64) int {
	n := len(data)
	if n <= 1 {
		return 1
	}
	iqr := CalculateIQR(data)
	h := 2 * iqr / math.Pow(float64(n), 1.0/3.0)
	min, max := minMax(data)
	return int(math.Ceil((max - min) / h))
}

// MinMax calculates the minimum and maximum values of a dataset.
func minMax(data []float64) (float64, float64) {
	min := data[0]
	max := data[0]
	for _, v := range data {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	return min, max
}

// ConvertToGroupedData converts the input data into grouped data using Freedman-Diaconis rule.
func ConvertToGroupedData(data []float64) [][]float64 {
	groupedData := make([][]float64, 0)
	numData := len(data)
	if numData == 0 {
		return groupedData
	}

	// Step 1: Find the number of class interval
	dataSorted := make([]float64, numData)
	copy(dataSorted, data)
	sort.Float64s(dataSorted)

	// Step 2: Group the data into the class interval
	numBins := FreedmanDiaconis(dataSorted)
	min, max := minMax(dataSorted)
	binWidth := (max - min) / float64(numBins)

	groupedData = make([][]float64, numBins)
	for i := range groupedData {
		groupedData[i] = make([]float64, 2) // groupedData[i][0] for midpoint, groupedData[i][1] for count
	}

	for _, value := range dataSorted {
		binIndex := int((value - min) / binWidth)
		if binIndex == numBins { // Handle the edge case where value == max
			binIndex--
		}
		groupedData[binIndex][0] = min + float64(binIndex)*binWidth + binWidth/2 // Calculate midpoint
		groupedData[binIndex][1]++                                               // Increment count
	}
	return groupedData
}

// SturgessInterval calculates the class interval using Sturgess's rule.
func SturgessInterval(n int) float64 {
	return 1.253314*math.Pow(float64(n), -1/3.0) - 0.007785*math.Pow(float64(n), 1/3.0)
}
