package main

import (
	"fmt"
	"os"
	"strings"
)

const corpus = "lazy cat jumps again and again and again since the beginning this was very important"

func main() {
	filtered := "and or was the since very"
	args := os.Args[1:]
	if l := len(args); l < 1 {
		fmt.Println("Please give me a word to search.")
		return
	}

	words := strings.Fields(corpus)
	for _, a := range args {
		for j, w := range words {
			if lowera := strings.ToLower(a); w == lowera && !strings.Contains(filtered, lowera) {
				fmt.Printf("#%d : %q\n", j+1, w)
			}
		}
	}

}
