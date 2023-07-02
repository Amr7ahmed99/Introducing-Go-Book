package main

import "fmt"

func main() {
	fmt.Print("Hello")
	fizzOrBuzz()
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
