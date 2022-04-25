package main

import (
	"fmt"
	"os"
)

func main() {

	names := os.Args
	peopleCount := len(names) - 1
	if peopleCount < 3 {
		fmt.Printf("You must enter 3 names\n")
		return
	}

	fmt.Printf("Ther are %d people!\n", peopleCount)
	fmt.Println("Hello great", names[1], "!")
	fmt.Println("Hello grea", names[2], "!")
	fmt.Println("Hello great", names[3], "!")
	fmt.Println("Nice to meet you all.")

}
