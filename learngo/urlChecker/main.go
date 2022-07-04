package main

import (
	"fmt"
	"net/http"
)

type requestResult struct {
	url    string
	status string
}

func main() {
	c := make(chan requestResult)
	var results = make(map[string]string)
	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://www.google.com/",
		"https://soundcloud.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
		"https://academy.nomadcoders.co/",
	}

	for _, url := range urls {

		go hitUrl(url, c)

	}
	for i := 0; i < len(urls); i++ {
		printUrl(&results, c)
	}
	for url2, status2 := range results {
		fmt.Println(url2, status2)
	}

}

// <- send only
func hitUrl(url string, c chan<- requestResult) {
	fmt.Println("checking url", url)
	resp, err := http.Get(url)
	status := "OK"
	if err != nil || resp.StatusCode >= 400 {
		status = "FAILED"
	}
	c <- requestResult{url: url, status: status}
}

// reciver olny
func printUrl(results *map[string]string, c <-chan requestResult) {
	result := <-c
	(*results)[result.url] = result.status
}
