package main

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
	"time"
)

// <-chan - ready-only channel
var titleRegex = regexp.MustCompile(`<title>(.*?)</title>`)

func title(urls ...string) <-chan string {
	c := make(chan string)
	for _, url := range urls {
		go func(url string) {
			resp, err := http.Get(url)
			if err != nil {
				c <- "Error: " + err.Error()
				return
			}
			defer resp.Body.Close()
			
			html, err := io.ReadAll(resp.Body)
			if err != nil {
				c <- "Error reading body"
				return
			}

			matches := titleRegex.FindStringSubmatch(string(html))
			if len(matches) > 1 {
				c <- matches[1]
			} else {
				c <- "No title found"
			}
		}(url)
	}
	return c
}

func fastest(url1, url2, url3 string) string {
	c1 := title(url1)
	c2 := title(url2)
	c3 := title(url3)

	// concurrency control
	select {
	case t1 := <-c1:
		return t1
	case t2 := <-c2:
		return t2
	case t3 := <-c3:
		return t3
	case <-time.After(1000 * time.Millisecond):
		return "everyone lost!"
		// default:
		// 	return "Without answer!"
	}
}

func main() {
	winner := fastest(
		"https://www.youtube.com",
		"https://www.amazon.com",
		"https://www.google.com",
	)
	fmt.Println(winner)
}
