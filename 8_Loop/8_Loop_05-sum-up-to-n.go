package main

import (
	"fmt"
	"os"
	"strconv"
)

const (
	usage             = "[min] [max]"
	isNotaNumber      = "%q or %q are not a number\n"
	minGreaterThanMax = "Min value not be greater or equal than max value"
)

func main() {
	args := os.Args
	l := len(args) - 1
	if l < 2 {
		fmt.Printf("%s\n", usage)
		return
	}
	min, err := strconv.Atoi(args[1])
	max, err2 := strconv.Atoi(args[2])
	if err != nil || err2 != nil {
		fmt.Printf(isNotaNumber, args[1], args[2])
		return
	}
	if min >= max {
		fmt.Println(minGreaterThanMax)
		return
	}

	var sum int
	for i := min; i <= max; i++ {
		sum += i
		fmt.Print(i)
		if i != max {
			fmt.Print(" + ")
		} else {
			fmt.Print(" = ")
		}
	}
	fmt.Printf("%d\n", sum)
}
