package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("Please provide directory paths")
		return
	}

	var dirnames []byte

	for _, dir := range args {
		files, err := ioutil.ReadDir(dir)
		if err != nil {
			fmt.Println(err)
			return
		}
		dirnames = append(dirnames, dir...)
		dirnames = append(dirnames, '\n')
		for _, d := range files {
			if d.IsDir() {
				dirnames = append(dirnames, '\t', '\t')
				dirnames = append(dirnames, d.Name()...)
			}
		}
		dirnames = append(dirnames, '\n')
	}

	err2 := ioutil.WriteFile("dirs.txt", dirnames, 0644)
	if err2 != nil {
		fmt.Println(err2)
		return
	}
}
