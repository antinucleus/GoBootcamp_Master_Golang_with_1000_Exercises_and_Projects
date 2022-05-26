package main

import "fmt"

func main() {

	words := []string{
		"gopher",
		"programmer",
		"go language",
		"go standard library",
	}

	var bwords [][]byte

	for _, s := range words {
		byteslice := []byte(s)
		fmt.Printf("%v\n", byteslice)

		bwords = append(bwords, byteslice)
	}

	for _, s := range bwords {
		fmt.Printf("%s\n", s)
	}
}
