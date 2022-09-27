package main

import "os"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	var products []Product
	products = append(products, Product{
		Id:     "1",
		Price:  10,
		Amount: 15,
	}, Product{
		Id:     "2",
		Price:  20,
		Amount: 20,
	}, Product{
		Id:     "3",
		Price:  30,
		Amount: 33,
	})
	var data []byte
	for _, product := range products {
		data = append(data, []byte(product.serialize())...)
	}
	err := os.WriteFile("./GoBases/packagesFmtOs/data.csv", data, 0644)
	check(err)
}
