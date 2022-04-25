package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	usage = `
	[command] [string]
	
	Available commands: lower, upper and title
	`
)

func main() {
	args := os.Args
	l := len(args) - 1

	if l < 2 {
		fmt.Printf("%s\n", usage)
		return
	}
	command := args[1]
	str := args[2]
	switch command {
	case "lower":
		str = strings.ToLower(str)
	case "upper":
		str = strings.ToUpper(str)
	case "title":
		str = strings.Title(str)
	default:
		fmt.Printf("Unknown command: %q\n", command)
	}
	fmt.Printf("%s\n", str)

}
