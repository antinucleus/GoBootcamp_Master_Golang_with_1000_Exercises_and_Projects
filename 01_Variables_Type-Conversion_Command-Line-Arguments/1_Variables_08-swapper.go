package main

import "fmt"

func main() {

	color, color2 := "red", "blue"
	color, color2 = color2, color

	fmt.Printf("color :%s\ncolor2 :%s\n", color, color2)
}
