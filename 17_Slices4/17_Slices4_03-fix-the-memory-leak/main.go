package main

import (
	"api"
	"fmt"
	"io/ioutil"
)

func main() {
	api.Report()

	millions := api.Read()

	newMillions := make([]int, 10)

	copy(newMillions, millions[len(millions)-10:])
	millions = newMillions

	fmt.Printf("\nLast 10 elements: %d\n\n", millions)

	api.Report()

	fmt.Fprintln(ioutil.Discard, millions[0])
}
