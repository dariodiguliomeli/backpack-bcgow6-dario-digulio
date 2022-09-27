package main

import "fmt"

func main() {
	const (
		MIN_EDAD       = 22
		ES_EMPLEADO    = true
		MIN_ANTIGUEDAD = 1
		MIN_SUELDO     = 100000.00
	)

	var (
		edadEmpleado       = 23
		situacionLaboral   = true
		antiguedadEmpleado = 1
		sueldoEmpleado     = 200000.00
	)

	switch {
	case edadEmpleado <= MIN_EDAD:
		fmt.Println("Debes ser mayor de 22 años de edad.")
	case situacionLaboral != ES_EMPLEADO:
		fmt.Println("Debes estar empleado.")
	case antiguedadEmpleado < MIN_ANTIGUEDAD:
		fmt.Println("Debes tener una antiguedad mayor a 1 año")
	case sueldoEmpleado < MIN_SUELDO:
		fmt.Println("Otorgar credito sin interés ✅")
	default:
		fmt.Println("Otorgar credito con interés ✔️")
	}
}
