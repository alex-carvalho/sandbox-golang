package main

import (
	"fmt"
	"reflect"
)

func main() {

	// array_example()
	// slice_example()
	map_example()

}

func array_example() {
	var firstArray [3]string
	firstArray[0], firstArray[1], firstArray[2] = "test1", "test2", "test3"
	fmt.Println(firstArray)

	for i := 0; i < len(firstArray); i++ {
		fmt.Println(firstArray[i])
	}

	secondArray := [...]float32{23.555, 56.9, 89.0}

	for index, value := range secondArray {
		fmt.Printf("index: %d | value: %.2f \n", index, value)
	}
}

func slice_example() {
	array := [3]int{4, 3, 2}
	slice := []int{4, 3, 2}
	fmt.Println(array, slice)
	fmt.Println(reflect.TypeOf(array), reflect.TypeOf(slice))

	nSlice := array[0:1]

	fmt.Println(nSlice)

	// using make
	// slice use a internal array
	newSlice := make([]int, 5, 10) // type, size, capacity
	fmt.Println(newSlice, len(newSlice), cap(newSlice))

	// when append more than the capacity is created a new internal array
	newSlice = append(newSlice, 1, 2, 3, 4, 5, 6)
	fmt.Println(newSlice, len(newSlice), cap(newSlice))

	// copy values from one slice to another
	// copy do not expande de array
	as1 := []int{3, 4, 5, 6, 7, 8}
	as2 := make([]int, 3)

	copy(as2, as1)

	fmt.Println(as1, as2)
}

func map_example() {

	m1 := make(map[string]int)

	// add values
	m1["A"] = 1
	m1["B"] = 2
	m1["C"] = 5

	fmt.Println(m1)

	// get value by key
	fmt.Println(m1["A"])

	// delete key
	delete(m1, "B")
	fmt.Println(m1)

	m2 := map[string]float32{
		"K": 1.1,
		"J": 3.3,
	}

	fmt.Println(m2)

	multiMap := map[string]map[string]int{
		"A": {
			"Alex": 1,
			"Ana":  2,
		},
		"J": {
			"John": 5,
		},
	}

	fmt.Println(multiMap)

}
