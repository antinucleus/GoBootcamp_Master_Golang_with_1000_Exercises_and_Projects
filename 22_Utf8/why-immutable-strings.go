package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// empty := ""
	// dump(empty)
	hello := "hello"
	dump(hello)
	dump("hello")
	dump("hello!")
	for i := range hello {
		dump(hello[i : i+1])

	}
	dump(string([]byte(hello)))
	dump(string([]byte(hello)))
	dump(string([]byte(hello)))

}

type StringHeader struct {
	pointer uintptr
	length  int
}

func dump(s string) {

	ptr := *(*StringHeader)(unsafe.Pointer(&s))
	fmt.Printf("%-12q: %+v\n", s, ptr)

}
