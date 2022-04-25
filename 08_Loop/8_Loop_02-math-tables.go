package main

import (
	"fmt"
	"os"
	"strconv"
)

const (
	usage        = "Usage: [op=*/+-] [size]"
	invalidOpMsg = "Invalid operator."
)

func main() {
	args := os.Args
	l := len(args) - 1
	if l < 2 {
		fmt.Println(usage)
		return
	}

	op, s := args[1], args[2]
	size, err := strconv.Atoi(s)

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

		switch op {
		case "*":
			for j := 0; j <= size; j++ {
				fmt.Printf("%6d", i*j)
			}
		case "-":
			for j := 0; j <= size; j++ {
				fmt.Printf("%6d", i-j)
			}
		case "/":
			for j := 0; j <= size; j++ {
				if j != 0 {
					fmt.Printf("%6d", i/j)
				} else {
					fmt.Printf("%6d", 0)
				}
			}
		case "+":
			for j := 0; j <= size; j++ {
				fmt.Printf("%6d", i+j)
			}
		default:
			fmt.Printf("%s\n", invalidOpMsg)
			return
		}

		fmt.Println()
	}

}
