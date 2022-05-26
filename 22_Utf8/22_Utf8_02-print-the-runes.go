package main

import "fmt"

func main() {
	const word = "console"

	fmt.Printf("%-12s %-12s %-12s\n", "dec", "hex", "bin")
	for _, r := range word {
		fmt.Printf("%-12d %-12[1]x %-12[1]b\n", r)
	}

	wordByte := make([]byte, 7)

	wordByte[0] = word[0]
	wordByte[1] = word[1]
	wordByte[2] = word[2]
	wordByte[3] = word[3]
	wordByte[4] = word[4]
	wordByte[5] = word[5]
	wordByte[6] = word[6]

	fmt.Printf("with decimals\t\t:%12s\n", wordByte)

	wordByte[0] = 'c'
	wordByte[1] = 'o'
	wordByte[2] = 'n'
	wordByte[3] = 's'
	wordByte[4] = 'o'
	wordByte[5] = 'l'
	wordByte[6] = 'e'

	fmt.Printf("with runes\t\t:%12s\n", wordByte)

	wordByte[0] = 0x63
	wordByte[1] = 0x6f
	wordByte[2] = 0x6e
	wordByte[3] = 0x73
	wordByte[4] = 0x6f
	wordByte[5] = 0x6c
	wordByte[6] = 0x65

	fmt.Printf("with hexadecimals\t:%12s\n", wordByte)

}
