package main

import (
	"fmt"
	"math"
)

func main() {
	//Array
	arr := [10]int{100, 200, 300}
	fmt.Println("array: ", arr)
	//slice
	slice := make([]int, 5, 10)
	fmt.Println("slice1: ", slice, len(slice), cap(slice))
	slice = arr[0:5]
	fmt.Println("slice2: ", slice, len(slice), cap(slice))
	slice = append(slice, 400, 500, 600)
	fmt.Println("slice3: ", slice, len(slice), cap(slice))
	copiedSlice := make([]int, 2)
	copy(copiedSlice, slice)
	fmt.Println("copiedSlice: ", copiedSlice, len(copiedSlice), cap(copiedSlice))
	for _, value := range slice {
		fmt.Print("\t", value)
	}
	//map
	person := map[string]string{
		"name": "Amr",
		"age":  "24",
		"id":   "1771999",
	}
	fmt.Println("\nmap: ", person)
	if value, ok := person["name"]; ok {
		fmt.Println("is present")
		fmt.Println(value, ok)
	}
	fmt.Println(person)
	delete(person, "name")
	fmt.Println(person)
	//slice of maps
	employees := make([]map[string]string, 2)
	employees[0] = map[string]string{"name": "amr"}
	employees[1] = map[string]string{"name": "omar"}
	employees = append(employees, map[string]string{"name": "saleh"})
	fmt.Println(employees)
	for _, value := range employees {
		fmt.Print("\t", value["name"])
	}
	fmt.Println()
	//exercise-1
	exerciseArray := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}
	fmt.Println("the smallest number is ", findTheSmallestNumberInArray(exerciseArray))

}

func findTheSmallestNumberInArray(arr []int) int {
	min := math.MaxInt
	for _, value := range arr {
		if value < min {
			min = value
		}
	}
	return min
}
