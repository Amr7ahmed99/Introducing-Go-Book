package channel

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func Pinger(c chan<- string) {
	for i := 0; ; i++ {
		c <- "ping"
	}
}

func Ponger(c chan<- string) {
	for i := 0; ; i++ {
		c <- "pong"
	}
}

func Printer(c <-chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

type HomePage struct {
	URL  string
	size int
}

func FetchWebPages() {
	urls := []string{
		"http://www.apple.com",
		"http://www.amazon.com",
		"http://www.google.com",
		"http://www.microsoft.com",
	}

	results := make(chan HomePage)
	for _, url := range urls {
		go func(url string) {
			res, err := http.Get(url)
			if err != nil {
				panic(err)
			}

			defer res.Body.Close()

			bs, err := ioutil.ReadAll(res.Body)
			if err != nil {
				panic(err)
			}
			results <- HomePage{URL: url, size: len(bs)}
		}(url)
	}

	var biggestPage HomePage

	for range urls {
		page := <-results
		fmt.Println(page.URL, page.size)
		if page.size > biggestPage.size {
			biggestPage = page
		}
	}
	fmt.Println("The biggest home page:", biggestPage.URL)
}
