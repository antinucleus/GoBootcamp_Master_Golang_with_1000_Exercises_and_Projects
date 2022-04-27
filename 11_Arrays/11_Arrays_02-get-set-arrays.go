package main

import (
	"fmt"
	"strings"
)

func main() {
	var (
		names     [3]string
		distances [5]int
		data      [5]byte
		ratios    [1]float64
		alives    [4]bool
		zero      [0]byte
	)

	names[0] = "İbni Sina"
	names[1] = "Mimar Sinan"
	names[2] = "Gazali"

	distances[0] = 10
	distances[1] = 20
	distances[2] = 70
	distances[3] = 35
	distances[4] = 185

	data[0] = 'N'
	data[1] = 'I'
	data[2] = 'G'
	data[3] = 'H'
	data[4] = 'T'

	ratios[0] = 2.76

	alives[0] = false
	alives[1] = true
	alives[2] = true
	alives[3] = false

	_ = zero

	separator := "\n" + strings.Repeat("=", 20) + "\n"

	fmt.Print("names", separator)
	for i := 0; i < len(names); i++ {
		fmt.Printf("names[%d]: %q\n", i, names[i])
	}

	fmt.Print("\ndistances", separator)
	for i := 0; i < len(distances); i++ {
		fmt.Printf("distances[%d]: %d\n", i, distances[i])
	}

	fmt.Print("\ndata", separator)
	for i := 0; i < len(data); i++ {
		fmt.Printf("data[%d]: %d\n", i, data[i])
	}

	fmt.Print("\nratios", separator)
	for i := 0; i < len(ratios); i++ {
		fmt.Printf("ratios[%d]: %.2f\n", i, ratios[i])
	}

	fmt.Print("\nalives", separator)
	for i := 0; i < len(alives); i++ {
		fmt.Printf("alives[%d]: %t\n", i, alives[i])
	}

	fmt.Print("\nzero", separator)
	for i := 0; i < len(zero); i++ {
		fmt.Printf("zero[%d]: %d\n", i, zero[i])
	}
}
