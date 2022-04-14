package main

import "fmt"

func main() {

	var car interface{}
	car = "BMW"
	fmt.Println(car)
	car = 1.2
	fmt.Println(car)

	type generic interface{}

	var g generic
	g = "BMW"
	fmt.Println(g)
	g = 1.2
	fmt.Println(g)

}
