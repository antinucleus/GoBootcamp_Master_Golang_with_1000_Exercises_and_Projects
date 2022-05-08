package main

import (
	"fmt"
	"runtime"
	"runtime/debug"
	"unsafe"
)

const size = 1e7

func main() {
	debug.SetGCPercent(-1)
	report("initial memory usage")
	var arr [size]int
	report("after declaring an array")
	newArr := arr
	report("after copying the array")
	passArray(newArr)
	slc := arr[:]
	slc2 := arr[:1000]
	slc3 := arr[1000:10000]
	report("after slicings")

	passSlice(slc3)

	fmt.Println()
	fmt.Printf("Arr size : %d bytes.\n", unsafe.Sizeof(arr))
	fmt.Printf("Newarr size: %d bytes.\n", unsafe.Sizeof(newArr))
	fmt.Printf("Slc size: %d bytes.\n", unsafe.Sizeof(slc))
	fmt.Printf("Slc2 size: %d bytes.\n", unsafe.Sizeof(slc2))
	fmt.Printf("Slc3 size: %d bytes.\n", unsafe.Sizeof(slc3))

}

func passArray(items [size]int) {
	items[0] = 100
	report("inside passArray")
}

// only passes 24-bytes of slice header
//
// observe that passing a slice doesn't affect the memory usage
func passSlice(items []int) {
	items[0] = 100
	report("inside passSlice")
}

// reports the current memory usage
// don't worry about this code
func report(msg string) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("[%s]\n", msg)
	fmt.Printf("\t> Memory Usage: %v KB\n", m.Alloc/1024)
}
