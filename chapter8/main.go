package main

import (
	"container/list"
	"fmt"
	"io"
	"net/http"
	"sort"

	m "./math"
)

type Person struct {
	name string
	age  int
}

type ByName []Person

func (ps ByName) Len() int {
	return len(ps)
}

func (ps ByName) Less(i, j int) bool {
	return ps[i].name < ps[j].name
}

func (ps ByName) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}

type ByAge []Person

func (ps ByAge) Len() int {
	return len(ps)
}
func (ps ByAge) Less(i, j int) bool {
	return ps[i].age < ps[j].age
}
func (ps ByAge) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}
func main() {
	var linkedList list.List
	linkedList.PushBack(2)
	linkedList.PushFront(1)
	linkedList.PushBack(3)
	linkedList.PushBack(4)
	for el := linkedList.Front(); el != nil; el = el.Next() {
		fmt.Println(el.Value.(int))
	}

	kids := []Person{
		{name: "mohamed", age: 10},
		{name: "ahmed", age: 11},
		{name: "ali", age: 8},
	}
	fmt.Println("before sorting by name", kids)
	sort.Sort(ByName(kids))
	fmt.Println("after sorting by name", kids)

	sort.Sort(ByAge(kids))
	fmt.Println("after sorting by age", kids)

	xs := []float64{1, 2, 3, 4}
	avg := m.Average(xs)
	fmt.Println("Average=", avg)

	http.HandleFunc("/hello", helloHttp)
	http.ListenAndServe(":5000", nil)
}

func helloHttp(res http.ResponseWriter, req *http.Request) {
	res.Header().Set(
		"Content-Type",
		"text/html",
	)
	io.WriteString(res,
		`<DOCTYPE html>
			<html>
				<head>
					<title>Hello, World</title>
				</head>
				<body>
					Hello, World!
				</body>
			</html>`,
	)
}
