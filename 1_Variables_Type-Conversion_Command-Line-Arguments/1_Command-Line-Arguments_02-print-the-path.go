package main

import (
	"fmt"
	"os"
	"path"
)

func main() {

	_, file := path.Split(os.Args[0])

	fmt.Println(file)
}

// run this command   go build -o myprogram  1_Command-Line-Arguments_02-print-the-path.go
// this command create myprogram file then exeute it with this command ./myprogram   and see result
