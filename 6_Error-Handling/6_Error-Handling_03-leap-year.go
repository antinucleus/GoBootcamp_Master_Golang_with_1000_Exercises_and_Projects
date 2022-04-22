package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args
	l := len(args) - 1
	if l < 1 {
		fmt.Println("Give me a year number")
		return
	}

	year, error := strconv.Atoi(os.Args[1])
	if error != nil {
		fmt.Printf("%q is not a valid year.\n", os.Args[1])
		return
	}

	var leap bool
	if year%400 == 0 {
		leap = true
	} else if year%100 == 0 {
		leap = false
	} else if year%4 == 0 {
		leap = true
	}

	if leap {
		fmt.Printf("%d is a leap year.\n", year)
	} else {
		fmt.Printf("%d is not a leap year.\n", year)
	}
}
