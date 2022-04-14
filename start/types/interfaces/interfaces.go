package main

import "fmt"

type printable interface {
	print()
}

type person struct {
	name string
	age  int
}

type product struct {
	name  string
	price float64
}

func (p person) print() {
	fmt.Printf("%s is %d years old\n", p.name, p.age)
}

func (p product) print() {
	fmt.Printf("%s costs %.2f\n", p.name, p.price)
}

func main() {
	var p printable
	p = person{name: "John", age: 30}
	p.print()
	p = product{name: "iPhone", price: 1000}
	p.print()
}
