package main

import "fmt"



func increment(x int) {
	x++
}

func increment2(x *int) {
	*x++
}

func main() {
	x := 1
	increment(x)
	fmt.Println(x)

	increment2(&x)
	fmt.Println(x)
}