package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	args := os.Args
	l := len(args)

	if l < 1 {
		fmt.Println("Requires age")
		return
	}
	age, error := strconv.Atoi(args[1])
	if error != nil || age < 0 {
		fmt.Println("Check your input")
	} else if age > 17 {
		fmt.Println("R-rated")
	} else if age > 13 {
		fmt.Println("PG-13")
	} else {
		fmt.Println("PG-rated")
	}
}
