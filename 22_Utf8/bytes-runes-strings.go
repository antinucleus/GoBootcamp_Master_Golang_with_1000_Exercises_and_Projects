package main

import (
	"fmt"
	"unicode/utf8"
	"unsafe"
)

func main() {

	str := "Yūgen ☯ 💀"
	bytes := []byte(str)
	str = string(bytes)

	fmt.Printf("%s\n", str)
	fmt.Printf("\t%d bytes \n", len(str))
	fmt.Printf("\t%d runes \n", utf8.RuneCountInString(str))

	fmt.Printf("% x\n", bytes)
	fmt.Printf("\t%d bytes \n", len(bytes))
	fmt.Printf("\t%d runes \n", utf8.RuneCount(bytes))

	fmt.Println()

	for i, r := range str {
		fmt.Printf("str[%2d] = % -12x = %q\n", i, string(r), r)
	}

	fmt.Println()
	fmt.Printf("1st byte\t:%c\n", str[0])
	fmt.Printf("2nd byte\t:%c\n", str[1])
	fmt.Printf("2nd rune\t:%s\n", str[1:3])
	fmt.Printf("Last rune\t:%s\n", str[len(str)-4:])

	fmt.Println()
	runes := []rune(str)
	fmt.Printf("%s\n", string(runes))
	fmt.Printf("\t%d bytes \n", int(unsafe.Sizeof(runes[0]))*len(runes))
	fmt.Printf("\t%d runes \n", len(runes))

	fmt.Printf("1st rune\t:%c\n", runes[0])
	fmt.Printf("2nd rune\t:%c\n", runes[1])
	fmt.Printf("first five runes\t:%c\n", runes[:5])

}
