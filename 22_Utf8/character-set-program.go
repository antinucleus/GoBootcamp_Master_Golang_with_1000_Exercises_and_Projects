package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var start, stop int

	if args := os.Args[1:]; len(args) == 2 {
		start, _ = strconv.Atoi(args[0])
		stop, _ = strconv.Atoi(args[1])
	}

	if start == 0 || stop == 0 {
		start, stop = 'A', 'Z'
	}

	fmt.Printf("%-10s %-10s %-10s %-10s\n%s\n", "literal", "dec ", "hex", "encoded", strings.Repeat("-", 45))
	for n := start; n <= stop; n++ {
		fmt.Printf("%-10c %-10[1]d %-10[1]x % -12x \n", n, string(n))
	}

}

//  [1 byte]

// 	asciiStart     = 32
// 	asciiStop      = 127
// 	upperCaseStart = 65
// 	upperCaseStop  = 90
// 	loverCaseStart = 97
// 	loverCaseStop  = 122

//  [2 bytes]

// 	latin1Start = 161
// 	latin1Stop  = 255

//  [3 bytes]

// 	dingbatStart = 9984
// 	dingbatStop  = 10175

//  [4 bytes]

// 	emojiStart = 128512
// 	emojiStop  = 128591
