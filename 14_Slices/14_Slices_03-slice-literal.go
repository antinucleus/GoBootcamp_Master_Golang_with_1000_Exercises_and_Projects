package main

import "fmt"

func main() {
	var (
		names     []string
		distances []int
		data      []byte
		ratios    []float64
		alives    []bool
	)

	names = []string{"Mustaaf", "Faruk", "MF"}
	distances = []int{10, 20, 40, 50, 60}
	data = []byte{'N', 'I', 'G', 'H', 'T'}
	ratios = []float64{2.1, 5.5}
	alives = []bool{false}

	fmt.Printf("names    : %T %d %t\n", names, len(names), names == nil)
	fmt.Printf("distances: %T %d %t\n", distances, len(distances), distances == nil)
	fmt.Printf("data     : %T %d %t\n", data, len(data), data == nil)
	fmt.Printf("ratios   : %T %d %t\n", ratios, len(ratios), ratios == nil)
	fmt.Printf("alives   : %T %d %t\n", alives, len(alives), alives == nil)
	if len(distances) == len(data) {
		fmt.Println("The length of the distances and the data slices are the same.")
	}
}
