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
		fmt.Println("Pick a number")
		return
	}
	number, error := strconv.Atoi(args[1])

	if error != nil {
		fmt.Printf("%q is not a number\n", args[1])
	} else if number%8 == 0 {
		fmt.Printf("%d is an even number and it's divisible by 8\n", number)
	} else if number%2 == 0 {
		fmt.Printf("%d is an even number\n", number)
	} else {
		fmt.Printf("%d is a odd number\n", number)
	}

}
