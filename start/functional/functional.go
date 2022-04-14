package main

import "fmt"


// functions can be attached to variables
var sum = func(a, b int) int {
	return a + b
}


func multiply(a, b int) int {
	return a * b
}

func exec(f func(int, int) int, a, b int) int {
	return f(a, b)
}

func main(){
	fmt.Println(sum(1, 2))

	// functions can be declared inside other functions
	sub := func(a, b int) int {
		return a - b
	}

	fmt.Println(sub(5, 3))

	r := exec(multiply, 3, 2)

	fmt.Println(r)
}