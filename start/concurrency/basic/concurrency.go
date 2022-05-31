package main

import (
	"fmt"
	"time"
)

func printMessage(user, message string, times uint8) {
	for i := uint8(0); i < times; i++ {
		fmt.Println(user, ":", message, " - ", i)
		time.Sleep(time.Second)
	}
}

func printMessageWithChannel(user, message string, times uint, channel chan uint) {
	for i := uint(0); i < times; i++ {
		fmt.Println(user, ":", message, " - ", i)
		time.Sleep(time.Second)
		channel <- i
	}
}

func main() {
	// run sequential
	/*
		printMessage("user1", "hello", 1)
		printMessage("user2", "world", 1)
		fmt.Println("fim")
	*/

	// run concurrent but not wait, so not print enything
	/*
		go printMessage("user3", "hello", 2)
		go printMessage("user4", "world", 2)
		fmt.Println("fim")
	*/

	// run concurrent and wait using channel
	/*
		channel := make(chan uint)
		go printMessageWithChannel("user5", "hello", 3, channel)
		go printMessageWithChannel("user6", "world", 3, channel)
		for i := uint8(0); i < 6; i++ {
			fmt.Println(<-channel)
		}

		fmt.Println("fim")
	*/

	// run concurrent using channel with buffer
	channel := make(chan uint, 6)
	go printMessageWithChannel("user5", "hello", 3, channel)
	go printMessageWithChannel("user6", "world", 3, channel)

	time.Sleep(time.Second)
	fmt.Println("fim", <-channel)

}
