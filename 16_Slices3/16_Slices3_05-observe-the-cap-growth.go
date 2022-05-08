package main

import (
	"fmt"
)

func main() {

	var slice []int
	fmt.Printf("len:%-20dcap:%-20dgrowth:%.2f\n", len(slice), cap(slice), 0.)

	for i := 0; i < 1e7; i++ {
		cp := float64(cap(slice))
		slice = append(slice, 0)
		if cp == 0 || cp != float64(cap(slice)) {
			c := float64(cap(slice))
			growth := c / cp
			fmt.Printf("len:%-20dcap:%-20dgrowth:%.2f\n", len(slice), cap(slice), growth)
		}
	}

}
