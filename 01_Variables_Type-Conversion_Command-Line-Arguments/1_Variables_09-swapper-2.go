package main

import "fmt"

func main() {
	red, blue := "red", "blue"

	red, blue = blue, red

	fmt.Printf("red :%s\nblue :%s\n", red, blue)

}
