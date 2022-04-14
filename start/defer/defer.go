package main

import "fmt"


func exampleDefer() {
	// last defer will be executed first
 	defer fmt.Println("defer 1")
	defer fmt.Println("defer 2")
	fmt.Println("normal")
}

func deferWithError() {
	fmt.Println("\ndeferWithError")

	defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered", r)
        }
    }()
	fmt.Println("normal")
	panic("PANIC")
}



func panicOnDefer() {
	fmt.Println("\npanicOnDefer")
	defer fmt.Println("defer 3")
 	defer fmt.Println("defer 1")
	defer panic("PANIC")
	defer fmt.Println("defer 2")
	fmt.Println("normal")
}

func c() (i int) {
    defer func() { i++ }() // return 2
    return 1
}

func main() {
	fmt.Println(c())
	exampleDefer()
	deferWithError()
	panicOnDefer()
}