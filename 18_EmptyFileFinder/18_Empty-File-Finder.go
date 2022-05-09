package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	args := os.Args[1:]

	if l := len(args); l != 1 {
		fmt.Println("Please provide a directory")
		return
	}
	dir := args[0]
	files, err := ioutil.ReadDir(dir)

	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}

	for i, file := range files {
		if file.Size() == 0 {
			fmt.Printf("File%d: %s\n", i, file.Name())
		}
	}
}
