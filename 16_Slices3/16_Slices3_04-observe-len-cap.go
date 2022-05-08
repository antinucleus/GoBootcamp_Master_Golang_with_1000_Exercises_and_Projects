package main

import "fmt"

func main() {
	var games = []string{"pacman", "mario", "tetris", "doom"}
	fmt.Printf("len:%d  cap:%d\n", len(games), cap(games))
	fmt.Println()

	for i := 0; i <= len(games); i++ {
		s := games[:i]
		fmt.Printf("len:%d  cap:%d\n", len(s), cap(s))
	}

	fmt.Println()

	zero := games[:0]
	fmt.Printf("games    len: %d cap: %d\n", len(games), cap(games))
	fmt.Printf("zero       len: %d cap: %d\n", len(zero), cap(zero))

	for _, v := range []string{"ultima", "dagger", "pong", "coldspot", "zetra"} {
		zero = append(zero, v)
		fmt.Printf("zero's      len: %d cap: %d\n", len(zero), cap(zero))
	}

	fmt.Println()

	for n := range zero {
		s := zero[:n]
		fmt.Printf("zero[:%d]'s  len: %d cap: %d\n", n, len(s), cap(s))
	}
	fmt.Println()

	zero = zero[:cap(zero)]
	for n := range zero {
		s := zero[:n]
		fmt.Printf("zero[:%d]'s  len: %d cap: %d - %q\n", n, len(s), cap(s), s)
	}

	fmt.Println()

	zero[0] = "command & conquer"
	games[0] = "red alert"
	fmt.Printf("zero  : %q\n", zero)
	fmt.Printf("games : %q\n", games)
}
