package main

import (
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Product struct {
	Id     string
	Price  float64
	Amount int
}

func main() {
	var products []Product
	products = append(products, Product{
		Id:     "1",
		Price:  10,
		Amount: 1,
	}, Product{
		Id:     "2",
		Price:  20,
		Amount: 2,
	}, Product{
		Id:     "3",
		Price:  30,
		Amount: 3,
	})
	var data []byte
	for _, product := range products {
		serializedProduct := fmt.Sprintf("%s;%f;%d\n", product.Id, product.Price, product.Amount)
		serializedProductInBytes := []byte(serializedProduct)
		data = append(data, serializedProductInBytes...)
	}
	err := os.WriteFile("./GoBases/packagesFmtOs/data.csv", data, 0644)
	check(err)
}
