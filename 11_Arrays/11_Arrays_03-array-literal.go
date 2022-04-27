package main

import (
	"fmt"
	"strings"
)

func main() {

	var (
		names     = [3]string{"İbni Sina", "Mimar Sinan", "Gazali"}
		distances = [5]int{10, 20, 70, 35, 185}
		data      = [5]byte{'N', 'I', 'G', 'H', 'T'}
		ratios    = [1]float64{2.76}
		alives    = [4]bool{false, true, true, false}
		zero      = [0]byte{}
	)

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
