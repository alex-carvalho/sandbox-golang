package main

import "fmt"

type sport interface {
	isSport() bool
}

type lux interface {
	isLux() bool
}

type car interface {
	sport
	lux
	getBrand() string
}

type bmw struct {
	brand string
}

func (b bmw) isSport() bool {
	return true
}

func (b bmw) isLux() bool {
	return true
}

func (b bmw) getBrand() string {
	return b.brand
}

func main() {
	fmt.Println("Hello, playground")

	var c car
	c = bmw{brand: "BMW"}

	fmt.Println(c.isSport())
	fmt.Println(c.isLux())
	fmt.Println(c.getBrand())
}
