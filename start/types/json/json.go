package main

import (
	"encoding/json"
	"fmt"
)

type product struct {
	Name           string  `json:"name"`
	Price          float64 `json:"price"`
	otheHidenField string
}

func main() {

	p1 := product{Name: "iPhone", Price: 1000, otheHidenField: "hidden"}
	pjson, err := json.Marshal(p1)
	if err != nil {
		fmt.Printf("Error marshaling JSON: %v\n", err)
		return
	}
	fmt.Println(string(pjson))

	var p2 product
	jsonString := `{"name":"Macbook","price":2000}`
	if err := json.Unmarshal([]byte(jsonString), &p2); err != nil {
		fmt.Printf("Error unmarshaling JSON: %v\n", err)
		return
	}
	fmt.Println(p2)

}
