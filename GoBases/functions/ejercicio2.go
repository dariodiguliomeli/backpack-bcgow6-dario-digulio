package main

import (
	"errors"
	"fmt"
)

func calculateAvg(notes ...int) (average float64, error error) {
	sum := 0
	for _, note := range notes {
		if note < 0 {
			return 0, errors.New("no se puede procesar una nota negativa")
		}
		sum += note
	}
	return float64(sum / len(notes)), nil
}

func showAverage(notes ...int) {
	average, err := calculateAvg(notes...)
	if err == nil {
		fmt.Printf("El promedio del curso es: %f", average)
		fmt.Println()
	} else {
		fmt.Println(err)
	}
}

func main() {
	showAverage(2, 4, 5, 10, 9, 8, 2, 1)
	showAverage(5, 10, -5, 1, 9, 3)
}
