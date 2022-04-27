package main

import (
	"fmt"
	"strings"
)

func main() {
	book := [...]string{"Kafka's Revenge", "Stay Golden", "Everythingship"}
	upper := book
	lower := book

	for i, b := range book {
		upper[i] = strings.ToUpper(b)
		lower[i] = strings.ToLower(b)
	}

	fmt.Printf("books:%q\n", book)
	fmt.Printf("upper:%q\n", upper)
	fmt.Printf("lower:%q\n", lower)

}
