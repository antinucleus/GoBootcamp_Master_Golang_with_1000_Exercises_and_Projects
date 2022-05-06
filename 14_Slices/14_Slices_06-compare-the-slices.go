package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {

	namesA := "Da Vinci, Wozniak, Carmack"
	namesB := []string{"Wozniak", "Da Vinci", "Carmack"}

	namesAS := strings.Split(namesA, ", ")

	sort.Strings(namesAS)
	sort.Strings(namesB)

	fmt.Printf("%q\n", namesAS)
	fmt.Printf("%q\n", namesB)

	equal := true
	if len(namesB) == len(namesAS) {
		for a := range namesAS {
			if namesAS[a] != namesB[a] {
				equal = false
				break
			}
		}
		if equal {
			fmt.Println("Slices are equal")
			return
		}
	}
	fmt.Println("Slices are not equal")
}
