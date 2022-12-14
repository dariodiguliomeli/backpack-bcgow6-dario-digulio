package main

import "fmt"

// Una empresa marinera necesita calcular el salario de sus empleados
// basándose en la cantidad de horas trabajadas por mes y la categoría.
// Si es categoría C, su salario es de $1.000 por hora
// Si es categoría B, su salario es de $1.500 por hora más un %20 de su salario mensual
// Si es de categoría A, su salario es de $3.000 por hora más un %50 de su salario mensual
// Se solicita generar una función que reciba por parámetro la cantidad de
// minutos trabajados por mes y la categoría, y que devuelva su salario.

func main() {
	var (
		category = "C"
		hours    = 136.0
		salary   = 0.0
	)

	if category == "C" {
		salary = hours * 1000.0
	} else if category == "B" {
		salary = hours * 1500 * 1.2
	} else if category == "A" {
		salary = hours * 3000 * 1.5
	}
	fmt.Println(salary)
}
