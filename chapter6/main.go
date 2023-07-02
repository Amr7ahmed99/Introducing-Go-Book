package main

import (
	"fmt"
	"math"
)

func main() {
	arr := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}
	_, ok := findElementInArray(arr, 19)
	fmt.Println(ok)

	fmt.Println("sum= ", add(arr...))
	fmt.Println("sum= ", add(1, 2, 3))

	sum := func(x, y int) int {
		return x + y
	}
	fmt.Println("sum= ", sum(1, 2))

	nextEven := makeEvenGenerator()
	fmt.Println("first even: ", nextEven())
	fmt.Println("second even: ", nextEven())
	fmt.Println("third even: ", nextEven())

	fact := factorial(5)
	fmt.Println("factorial 5 is =", fact)

	num := 1
	changeValue(num)
	fmt.Println("value of num=", num)

	changeValueByRef(&num)
	fmt.Println("value of num using & operator=", num)

	addressInMemoryFitInteger := new(int)
	changeValueByRef(addressInMemoryFitInteger)
	fmt.Println("value of num using new keyword=", *addressInMemoryFitInteger)

}

func findElementInArray(arr []int, element int) (int, string) {
	for idx, value := range arr {
		if element == value {
			return idx, fmt.Sprintf("%d is exist at index %d.", element, idx)
		}
	}
	return -1, fmt.Sprintf("%d does not exist", element)
}

func add(args ...int) int { //variadic function
	total := 0
	for _, value := range args {
		total += value
	}
	return total
}

func makeEvenGenerator() func() uint {
	i := uint(0)
	return func() uint {
		ret := i
		i += 2
		return ret
	}
}

func factorial(n uint) uint {
	if n == 0 {
		return 1
	} //base case

	return n * factorial(n-1)
}

func changeValue(num int) {
	num = 10
}

func changeValueByRef(num *int) {
	*num = 5
}

func evenOrOdd(x int) bool {
	if x%2 == 0 {
		return true
	}
	return false
}

func findGreatestNumber(args ...int) int {
	maxValue := math.MinInt
	for _, value := range args {
		if value > maxValue {
			maxValue = value
		}
	}
	return maxValue
}

func makeOddGenerator() func() uint {
	i := uint(1)
	return func() uint {
		ret := i
		i += 2
		return ret
	}
}

func fibonacci(n int) int {
	if n == 0 {
		return 0
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func swapValuesByRef(x, y *int) {
	*x, *y = *y, *x
}
