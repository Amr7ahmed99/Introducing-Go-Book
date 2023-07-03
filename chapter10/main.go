package main

import (
	"fmt"
	"math/rand"
	"time"

	// ch "./channel"
	ch "./channel"
)

func f(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
		amt := time.Duration(rand.Intn(250))
		time.Sleep(time.Millisecond * amt)
	}
}

func main() {
	// for i := 0; i < 10; i++ {
	// 	go f(i)
	// }

	// var c chan string = make(chan string)
	// go ch.Pinger(c)
	// go ch.Ponger(c)
	// go ch.Printer(c)
	ch.FetchWebPages()
	var input string
	fmt.Scanln(&input)
}
