package main

import (
	"fmt"
	"os"
)

func main() {
	name := os.Args[1]

	greet := `hi ` + name + `!
how are you?`

	fmt.Println(greet)
}
