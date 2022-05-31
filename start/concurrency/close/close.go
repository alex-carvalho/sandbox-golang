package main

import (
	"fmt"
	"time"
)

func isOdd(i int) bool {
	return i%2 != 0
}

func oddGeretor(n int, c chan int) {
	for i := 0; i < n; i++ {
		if isOdd(i) {
			c <- i
			time.Sleep(time.Millisecond * 180)
		}
	}
	close(c)
}

func main() {
	c := make(chan int, 30)
	go oddGeretor(30, c)

	// infinit range until channel is closed
	for odd := range c {
		fmt.Printf("%d ", odd)
	}
	fmt.Println("Fim!")
}
