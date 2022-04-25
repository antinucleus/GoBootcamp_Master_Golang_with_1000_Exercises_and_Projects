package main

import (
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

func main() {
	var (
		vowel = "aeiou"
	)
	args := os.Args
	l := len(args) - 1
	letter := args[1]

	if l < 1 || utf8.RuneCountInString(letter) > 1 {
		fmt.Println("Give me a letter")
		return
	}

	if strings.IndexAny(vowel, letter) != -1 {
		fmt.Printf("%q is vowel\n", args[1])
	} else if letter == "w" || letter == "y" {
		fmt.Printf("%q is sometimes a vowel, sometimes not.\n", args[1])
	} else {
		fmt.Printf("%q is consonant.\n", args[1])
	}
}
