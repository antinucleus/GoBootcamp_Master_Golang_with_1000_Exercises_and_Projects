package main

import "fmt"

func main() {

	var (
		names    = []string{}
		distance = []int{}
		data     = []uint8{}
		ratios   = []float64{}
		alives   = []bool{}
	)

	fmt.Printf("names   :%T %d %t\n", names, len(names), names == nil)
	fmt.Printf("distance:%T %d %t\n", distance, len(distance), distance == nil)
	fmt.Printf("data	:%T %d %t\n", data, len(data), data == nil)
	fmt.Printf("ratios	:%T %d %t\n", ratios, len(ratios), ratios == nil)
	fmt.Printf("alives	:%T %d %t\n", alives, len(alives), alives == nil)

}
