package main

import "fmt"

type dimension struct {
	length, width, height float64
}

type props struct {
	dimension
	color string
}

type product struct {
	name     string
	price    float64
	discount float64
	props    props
}

func (p product) calculateDiscount() float64 {
	return p.price * (1 - p.discount)
}

func main() {
	p := product{
		name:     "iPhone",
		price:    699.00,
		discount: 0.1,
		props: props{
			dimension: dimension{
				length: 5.5,
			},
			color: "black",
		},
	}
	fmt.Println(p, p.calculateDiscount())
	fmt.Println(p.props.length)

	p2 := product{"Macbook", 1399.00, 0.2, props{dimension{12.5, 13.5, 14.5}, "silver"}}
	fmt.Println(p2, p2.calculateDiscount())
}
