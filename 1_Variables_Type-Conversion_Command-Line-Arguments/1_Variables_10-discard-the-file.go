package main

import (
	"fmt"
	"path"
)

func main() {
	directory, _ := path.Split("secret/file.txt")
	fmt.Println(directory)
}
