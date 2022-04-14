package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {

	// prints()
	// constvar()
	types()
}

func prints() {
	fmt.Print("same")
	fmt.Print(" line.")

	fmt.Println(" new")
	fmt.Println("line.")

	x := 3.141516
	xs := fmt.Sprint(x)

	fmt.Println("x: " + xs)
	fmt.Println("x:", x)

	fmt.Printf("x: %.2f.", x)

	a := 1
	b := 1.9999
	c := false
	d := "hello"
	fmt.Printf("\n%d %f %.1f %t %s", a, b, b, c, d)
	fmt.Printf("\n%v %v %v %v", a, b, c, d)
}

func constvar() {

	const PI float64 = 3.1415
	var raio = 3.2

	area := PI * math.Pow(raio, 2) // := declare witou specify 'var'
	fmt.Println("area", area)

	const (
		a = 1
		b // same value that a
	)

	var (
		c = 3
		d = 4
	)

	fmt.Println("a:", a, "b:", b, "c:", c, "d:", d)

	var e, f bool = true, false
	fmt.Println(e, f)

	g, h, i := 2, false, "hello!"
	fmt.Println(g, h, i)
}

func types() {
	// only positive -  uint8 uint16 uint32 uint64
	// with signal - int8 int16 int32 int64
	// doesn't have char type

	fmt.Println("TypeOf 32000", reflect.TypeOf(32000))

	var b byte = 255
	fmt.Println("byte is:", reflect.TypeOf(b))

	i1 := math.MaxInt64
	fmt.Println("max int", i1)
	fmt.Println("TypeOf MaxInt64:", reflect.TypeOf(i1))

	var i2 rune = 'a' // Unicode (int32) map
	fmt.Println("rune: ", reflect.TypeOf(i2))
	fmt.Println(i2)

	var x float32 = 49.99 // explicit type 32bits
	fmt.Println("x: ", reflect.TypeOf(x))
	fmt.Println("type of 49.99:", reflect.TypeOf(49.99)) // float64

	// boolean
	bo := true
	fmt.Println("bo:", reflect.TypeOf(bo))

	// string
	s1 := "Hello World, size:"
	fmt.Println(s1, len(s1))

	// string com multiplas linhas
	s2 := `multi
	line
	size:`
	fmt.Println(s2, len(s2))

	fmt.Println(string([]rune{i2}))
}
