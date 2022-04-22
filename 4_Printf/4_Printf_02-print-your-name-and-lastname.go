package main

import "fmt"

func main() {
	specifier := "My name is %s and my last name is %s\n"
	name := "Faruk"

	fmt.Printf(specifier, name, "Bağcı")
}
