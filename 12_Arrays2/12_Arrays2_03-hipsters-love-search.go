package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	books := [...]string{
		"Kafka's Revenge",
		"Stay Golden",
		"Everythingship",
		"Kafka's Revenge 2nd Edition",
	}

	args := os.Args[1:]

	if len(args) != 1 {
		fmt.Println("Tell me a book title")
		return
	}

	search := strings.ToLower(args[0])
	nobook := true
	fmt.Printf("Search Results:\n")
	for i, n := range books {
		n := strings.ToLower(n)
		if strings.Contains(n, search) {
			nobook = false
			fmt.Printf("+ %s\n", books[i])
		}
	}
	if nobook {
		fmt.Printf("We don't have the book: %q\n", args[0])
	}

}
