package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

// <-chan - ready-only channel
// generator is a function that returns a channel, hidden goroutine
func title(urls ...string) <-chan string {
	c := make(chan string)
	for _, url := range urls {
		go func(url string) {
			resp, _ := http.Get(url)
			html, _ := ioutil.ReadAll(resp.Body)

			r, _ := regexp.Compile("<title>(.*?)<\\/title>")
			c <- r.FindStringSubmatch(string(html))[1]
		}(url)
	}
	return c
}

func main() {
	t1 := title("https://www.cod3r.com.br", "https://www.google.com")
	t2 := title("https://www.amazon.com", "https://www.youtube.com")
	fmt.Println("First:", <-t1, "|", <-t2)
	fmt.Println("Second:", <-t1, "|", <-t2)
}
