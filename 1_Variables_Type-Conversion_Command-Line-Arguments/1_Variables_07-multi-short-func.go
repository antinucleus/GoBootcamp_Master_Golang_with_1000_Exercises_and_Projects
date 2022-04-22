package main

import "fmt"

func main() {

	_, b := multi()
	fmt.Printf("%d\n", b)
}

func multi() (int, int) {
	return 5, 4
}
