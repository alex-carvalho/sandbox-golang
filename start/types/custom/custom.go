package main

import "fmt"

type myNumber int

func (n myNumber) add(other int) myNumber {
	return n + myNumber(other)
}

func (n myNumber) isEven() bool {
	return n%2 == 0
}

func main() {
	var n myNumber = 1
	fmt.Println(n)
	fmt.Println("isEven", n.isEven())
	n = n.add(1)
	fmt.Println(n)
	fmt.Println("isEven", n.isEven())

}
