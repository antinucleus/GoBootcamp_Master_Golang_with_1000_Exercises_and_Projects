package main

import (
	"fmt"
	"strings"
)

func main() {
	const N = 3

	reqs := []int{
		500, 600, 250,
		200, 400, 50,
		900, 800, 600,
		750, 250, 100,
		150, 654, 235,
		320, 534, 765,
		121, 876, 285,
		543, 642,
	}
	if l := len(reqs); l%3 != 0 && l > 3 {
		switch l % 3 {
		case 2:
			reqs = append(reqs, 0)
		case 1:
			reqs = append(reqs, 0, 0)
		}
	}
	size := len(reqs) / N

	daily := make([][]int, 0, size)
	for i := range reqs {
		if i%3 == 0 {
			daily = append(daily, reqs[i:i+N])
		}
	}
	fmt.Printf("%#v\n", daily)

	fmt.Printf("%-10s%-10s\n", "Day", "Requests")
	fmt.Println(strings.Repeat("=", 20))
	var gt int

	for _, d := range daily {
		var dt int
		for _, id := range d {
			dt += id
		}
		gt += dt
		fmt.Printf("%9s %-10d\n\n", "TOTAL:", dt)
	}
	fmt.Printf("%9s %-10d\n", "GRAND:", gt)

}
