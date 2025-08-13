package main

import "fmt"

func main() {

	var car any
	car = "BMW"
	fmt.Println(car)
	car = 1.2
	fmt.Println(car)

	type generic any

	var g generic
	g = "BMW"
	fmt.Println(g)
	g = 1.2
	fmt.Println(g)

}
