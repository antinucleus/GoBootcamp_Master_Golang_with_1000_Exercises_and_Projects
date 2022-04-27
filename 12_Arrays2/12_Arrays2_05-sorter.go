package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	args := os.Args[1:]

	if l := len(args); l > 5 {
		fmt.Println("Sorry. Go arrays are fixed. So, for now, I'm only supporting sorting 5 numbers...")
		return
	} else if l != 5 {
		fmt.Println("Please give me up to 5 numbers.")
		return
	}

	var (
		numbers [5]int
	)

	for i, a := range args {
		n, err := strconv.Atoi(a)
		if err == nil {
			numbers[i] = n
		}
	}

	for range numbers {
		for i := range numbers {
			if i != len(numbers)-1 && numbers[i] > numbers[i+1] {
				numbers[i], numbers[i+1] = numbers[i+1], numbers[i]
			}
		}
	}
	fmt.Printf("%v\n", numbers)

}
