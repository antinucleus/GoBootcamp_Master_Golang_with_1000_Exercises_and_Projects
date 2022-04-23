package main

import "fmt"

func main() {
	var sum int
	for i := 1; i <= 10; i++ {
		sum += i
		fmt.Print(i)
		if i != 10 {
			fmt.Print(" + ")
		} else {
			fmt.Print(" = ")
		}
	}
	fmt.Printf("%d\n", sum)
}
