package main

import "fmt"

func main() {

	nums := []int{56, 89, 15, 25, 30, 50}

	var newNums []int
	newNums = append(newNums, nums[:3]...)

	mine := newNums

	mine[0], mine[1], mine[2] = -50, -100, -150

	fmt.Println("Mine         :", mine)
	fmt.Println("Original nums:", nums[:3])
}
