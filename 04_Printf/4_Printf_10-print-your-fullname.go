package main

import (
	"fmt"
	"os"
)

func main() {
	arg := os.Args
	specifier := "Your first name is %s and your last name is %s\n"
	fmt.Printf(specifier, arg[1], arg[2])
}
