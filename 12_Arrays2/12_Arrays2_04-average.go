package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	args := os.Args[1:]
	if l := len(args); l != 5 || l > 5 {
		fmt.Println("Please tell me numbers (maximum 5 numbers).")
		return
	}

	var (
		sum     int
		numbers [5]int
		total   int
	)

	for i, arg := range args {
		n, err := strconv.Atoi(arg)
		if err == nil {
			numbers[i] = n
			sum += n
			total++
		}

	}
	average := sum / total
	fmt.Printf("Your numbers :%v\n", numbers)
	fmt.Printf("Average :%v\n", average)

}
