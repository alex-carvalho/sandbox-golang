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
	pjson, _ := json.Marshal(p1)
	fmt.Println(string(pjson))

	var p2 product
	jsonString := `{"name":"Macbook","price":2000}`
	json.Unmarshal([]byte(jsonString), &p2)
	fmt.Println(p2)

}
