package main

import "fmt"

func calculateSalary(basicSalary float64) (tax float64) {
	const (
		MIN_TAX = 0.17
		ADD_TAX = 0.1
	)
	tax = MIN_TAX
	if basicSalary > 150.000 {
		tax += ADD_TAX
	}
	return basicSalary * tax
}

func main() {
	var user1Salary = 70000.00
	var user2Salary = 20000.00
	var user3Salary = 170000.00

	fmt.Printf("El salario final del primer empleado es %f", calculateSalary(user1Salary))
	fmt.Printf("El salario final del segundo empleado es %f", calculateSalary(user2Salary))
	fmt.Printf("El salario final del tercer empleado es %f", calculateSalary(user3Salary))
}
