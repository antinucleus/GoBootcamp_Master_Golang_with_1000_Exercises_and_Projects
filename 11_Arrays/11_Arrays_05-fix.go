package main

import "fmt"

func main() {
	var names = [...]string{
		"Einstein",
		"Shepard",
		"Tesla",
	}

	var books = [...]string{
		"Kafka's Revenge",
		"Stay Golden",
	}

	fmt.Printf("%q\n", names)
	fmt.Printf("%q\n", books)
}
