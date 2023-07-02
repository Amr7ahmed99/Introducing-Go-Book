package main

import "fmt"

func main() {

	fmt.Print("Hello")
	fizzOrBuzz()
}

func readNumberFromUser() {
	fmt.Print("Enter a number: ")
	var input float64
	fmt.Scanf("%f", &input)
	output := input * 2
	fmt.Println("You entered: ", output)
}

func convertFahrenheitIntoCelsius(fahrenValue float64) float64 {
	celsius := (fahrenValue - 32) * (5 / 9)
	return celsius
}

func convertFeetIntoMeters(feet float64) float64 {
	const feetInMeter = 0.3048
	return feet * feetInMeter
}

func oddsBetweenOneAndHundred() {
	for i := 3; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Print("\t", i)
		}
	}
}

func fizzOrBuzz() {
	for i := 3; i <= 100; i++ {
		if (i%3 == 0) && (i%5 == 0) {
			fmt.Print("\t", "FizzBuzz")
			continue
		} else if i%3 == 0 {
			fmt.Print("\t", "Fizz")
			continue
		} else if i%5 == 0 {
			fmt.Print("\t", "Buzz")
			continue
		}
		fmt.Print("\t", i)
	}
}
