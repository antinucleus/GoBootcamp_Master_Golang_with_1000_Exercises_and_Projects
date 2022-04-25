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
		fmt.Println("Give me the magnitude of the earthquake.")
		return
	}
	n, err := strconv.ParseFloat(args[1], 64)
	if err != nil {
		fmt.Println("I couldn't get that, sorry.")
		return
	}

	switch {

	case n >= 10:
		fmt.Printf("%.2f is massive\n", n)
	case n >= 8:
		fmt.Printf("%.2f is great\n", n)
	case n >= 7:
		fmt.Printf("%.2f is major\n", n)
	case n >= 6:
		fmt.Printf("%.2f is strong\n", n)
	case n >= 5:
		fmt.Printf("%.2f is moderate\n", n)
	case n >= 4:
		fmt.Printf("%.2f is light\n", n)
	case n >= 3:
		fmt.Printf("%.2f is minor\n", n)
	case n >= 2:
		fmt.Printf("%.2f is very minor\n", n)
	default:
		fmt.Printf("%.2f is micro\n", n)
	}
}
