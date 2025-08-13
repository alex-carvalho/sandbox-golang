package main

import "fmt"

func simple() {
	fmt.Println("f1")
}

func withParameters(a int, b string) {
	fmt.Println(a, b)
}

func withReturn() (int, string) {
	return 2, "world"
}

func namedReturn() (first int, second string) {
	first = 3 
	second = "go"
	return
}

func varargs(args ...float64) float64 {
	if len(args) == 0 {
		return 0
	}
	total := 0.0
	for _, v := range args {
		total += v
	}
	return total / float64(len(args))
}

func main() {
	simple()
	withParameters(1, "hello")
	a, b := withReturn()
	fmt.Println(a, b)
	first, second := namedReturn()
	fmt.Println(first, second)

	average := varargs(1.0, 2.0, 3.0, 4.0)
	fmt.Println(average)

	// only slice can be passed as varargs
	values := []float64{1.0, 2.0, 3.0, 4.0}
	average = varargs(values...)
	fmt.Println(average)
}


func init() {
	fmt.Println("init function is executed before main function")
}