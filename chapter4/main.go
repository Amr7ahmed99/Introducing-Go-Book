package main

import "fmt"

func main() {

	// var grades [5]float64
	// grades[0]= 100

	grades := [5]float64{
		98,
		97,
		81,
		73,
		88,
	}
	fmt.Print(getAvg(grades[:]))
}

func getAvg(arr []float64) float64 {
	var total float64 = 0
	for _, value := range arr {
		total += value
	}
	return total / float64(len(arr))
}
