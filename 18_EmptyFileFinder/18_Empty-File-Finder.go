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

	var total int

	for _, file := range files {
		if file.Size() == 0 {
			total += len(file.Name()) + 1
		}
	}
	fmt.Printf("Total required spaee: %d bytes.\n", total)

	names := make([]byte, 0, total)
	for _, file := range files {
		if file.Size() == 0 {
			name := file.Name()
			names = append(names, name...)
			names = append(names, '\n')
			// fmt.Printf("File%d: %s\n", i, name)
		}
	}

	err2 := ioutil.WriteFile("out.txt", names, 0644)
	if err2 != nil {
		fmt.Println(err2)
		return
	}
	fmt.Printf("names: %s\n", names)

}
