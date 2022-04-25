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
		fmt.Println("Give me the size of the table")
		return
	}

	size, err := strconv.Atoi(args[1])

	if err != nil || size < 0 {
		fmt.Println("Wrong size")
		return
	}

	fmt.Printf("%6s", "X")
	for i := 0; i <= size; i++ {
		fmt.Printf("%6d", i)
	}
	fmt.Println()
	for i := 0; i <= size; i++ {
		fmt.Printf("%6d", i)

		for j := 0; j <= size; j++ {
			fmt.Printf("%6d", i*j)
		}
		fmt.Println()
	}

}
