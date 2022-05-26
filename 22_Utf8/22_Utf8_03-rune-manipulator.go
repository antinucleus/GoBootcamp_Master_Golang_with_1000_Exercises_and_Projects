package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	words := []string{
		"cool",
		"güzel",
		"jīntiān",
		"今天",
		"read 🤓",
	}

	for _, s := range words {
		fmt.Printf("%s: %d bytes and its length %d. Bytes : % [1]x. Runes : ", s, len(s), utf8.RuneCountInString(s))
		for _, r := range s {
			fmt.Printf("%-3c", r)
		}

		fmt.Println()
		fmt.Println()

		r, size := utf8.DecodeRuneInString(s)
		fmt.Printf("\tfirst   : %q (%d bytes)\n", r, size)

		fmt.Println()
		fmt.Println()

		r, size = utf8.DecodeLastRuneInString(s)
		fmt.Printf("\tlast    : %q (%d bytes)\n", r, size)

		fmt.Println()
		fmt.Println()

		_, first := utf8.DecodeRuneInString(s)
		_, second := utf8.DecodeRuneInString(s[first:])
		fmt.Printf("\tfirst 2 : %q\n", s[:first+second])

		fmt.Println()
		fmt.Println()

		_, last1 := utf8.DecodeLastRuneInString(s)
		_, last2 := utf8.DecodeLastRuneInString(s[:len(s)-last1])
		fmt.Printf("\tlast 2  : %q\n", s[len(s)-last2-last1:])

		fmt.Println()
		fmt.Println()

		rs := []rune(s)
		fmt.Printf("\tfirst 2 : %q\n", string(rs[:2]))
		fmt.Printf("\tlast 2  : %q\n", string(rs[len(rs)-2:]))

		fmt.Println()
		fmt.Println()
	}

}
