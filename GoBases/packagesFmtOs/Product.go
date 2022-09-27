package main

import "fmt"

type Product struct {
	Id     string
	Price  float64
	Amount int
}

func (product Product) serialize() string {
	return fmt.Sprintf("%s;%f;%d\n", product.Id, product.Price, product.Amount)
}
