package main

import (
	"fmt"
	"sort"
)

func main() {
	items := []string{
		"pacman", "mario", "tetris", "doom", "galaga", "frogger",
		"asteroids", "simcity", "metroid", "defender", "rayman",
		"tempest", "ultima",
	}

	fmt.Println("Original:", items)

	mid := len(items)/2 - 1
	sort.Strings(items[mid : mid+3])

	fmt.Println()
	fmt.Println("Sorted  :", items)
}
