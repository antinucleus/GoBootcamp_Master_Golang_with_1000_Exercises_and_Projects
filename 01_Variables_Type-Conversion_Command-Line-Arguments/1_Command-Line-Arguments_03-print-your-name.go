package main

import (
	"fmt"
	"os"
)

func main() {

	name := "Hi" + os.Args[1]
	fmt.Println(name)
	fmt.Println("How are you?")
}
