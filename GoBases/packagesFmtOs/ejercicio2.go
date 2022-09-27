package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	file, err := os.ReadFile("./GoBases/packagesFmtOs/data.csv")
	Check(err)
	s := string(file)
	fmt.Println("Id\t\tPrecio\t\t\t\tCantidad")
	var total int64 = 0
	for index, line := range strings.Split(s, "\n") {
		if index != len(strings.Split(s, "\n"))-1 {
			parts := strings.Split(line, ";")
			fmt.Printf("%s\t\t%s\t\t\t\t%s\n", parts[0], parts[1], parts[2])
			amount, err := strconv.ParseInt(parts[2], 10, 9)
			Check(err)
			total += amount
		}
	}
	fmt.Printf("Total\t\t\t\t\t\t\t%d", total)
}
