package main

import "fmt"

func main() {

	var (
		perimeter     int
		width, heigth = 5, 6
	)

	perimeter = 2 * (width + heigth)
	fmt.Printf("perimeter :%d\n", perimeter)
}
