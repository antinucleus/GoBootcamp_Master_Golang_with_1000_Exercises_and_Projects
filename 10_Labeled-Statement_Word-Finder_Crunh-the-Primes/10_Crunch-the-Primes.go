package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var nInt []int
	args := os.Args
	n := args[1:]
	if len(n) < 1 {
		fmt.Println("Please enter the numbers that will be contolled if they are prime or not")
		return
	}

	for _, nm := range n {
		if ni, err := strconv.Atoi(nm); err == nil {
			nInt = append(nInt, ni)
		}
	}

	for _, v := range nInt {
		prime := true
		for i := 2; i < v; i++ {
			if v%i == 0 {
				prime = false
				break
			}
		}
		if prime && v != 1 {
			fmt.Printf("%-2d", v)
		}
	}
	fmt.Println()

}
