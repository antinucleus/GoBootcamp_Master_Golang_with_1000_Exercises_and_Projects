package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var (
		even []int
		odd  []int
	)
	data := "2 4 6 1 3 5"
	dataArray := strings.Fields(data)

	for _, el := range dataArray {
		n, _ := strconv.Atoi(el)
		if n%2 == 0 {
			even = append(even, n)
		} else {
			odd = append(odd, n)
		}
	}

	fmt.Printf("nums    :%v\n", dataArray)
	fmt.Printf("even	:%v\n", even)
	fmt.Printf("odd 	:%v\n", odd)
	fmt.Printf("middle 	:%v\n", dataArray[2:4])
	fmt.Printf("first 2 :%v\n", dataArray[:2])
	fmt.Printf("last 2 :%v\n", dataArray[len(dataArray)-2:])
	fmt.Printf("even last 1 :%v\n", even[len(even)-1:])
	fmt.Printf("odd last 2 :%v\n", odd[len(odd)-2:])

}
