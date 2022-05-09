package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"unicode/utf8"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("Send me some items and I will sort them")
		return
	}
	var total int
	for _, name := range args {
		total += utf8.RuneCountInString(name) + 1
	}
	fmt.Println("Total :", total)
	names := make([]byte, 0, total)

	sort.Strings(args)
	for _, name := range args {
		names = append(names, name...)
		names = append(names, '\n')
	}

	err := ioutil.WriteFile("sorted.txt", names, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}

}
