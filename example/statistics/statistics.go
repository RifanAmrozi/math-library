package main

import (
	"fmt"
    statistics "github.com/RifanAmrozi/math-library/statistics"
)

func main() {
	data := []float64{1, 2, 2, 3, 4, 5, 5, 5, 6, 7, 8, 9, 10, 10, 10} // Example dataset

	groupedData := statistics.ConvertToGroupedData(data)

	fmt.Println("Midpoint\tCount")
	for _, bin := range groupedData {
		fmt.Printf("%.2f\t\t%d\n", bin[0], int(bin[1]))
	}
}
