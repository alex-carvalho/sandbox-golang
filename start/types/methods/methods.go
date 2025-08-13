package main

import (
	"fmt"
	"strings"
)

type person struct {
	firstName string
	lastName  string
}

func (p person) FullName() string {
	return p.firstName + " " + p.lastName
}

func setFullName(p *person, name string) {
	names := strings.Split(name, " ")
	if len(names) >= 2 {
		p.firstName = names[0]
		p.lastName = names[1]
	} else if len(names) == 1 {
		p.firstName = names[0]
		p.lastName = ""
	}
}

func main() {
	p := person{
		firstName: "John",
		lastName:  "Doe",
	}
	fmt.Println(p.FullName())
	setFullName(&p, "Jane Doe")
	fmt.Println(p.FullName())
}
