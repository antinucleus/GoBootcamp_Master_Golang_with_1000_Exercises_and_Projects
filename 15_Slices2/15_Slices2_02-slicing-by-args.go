package main

import (
	"fmt"
	"os"
	"strconv"
)

const (
	error  = "Provide only the [starting] and [stopping] positions\n"
	error2 = "Wrong positions\n"
	error3 = "%q is not a number\n"
)

func main() {
	ships := []string{"Normandy", "Verrikan", "Nexus", "Warsaw"}
	fmt.Printf("%v\n", ships)

	args := os.Args[1:]

	if l := len(args); l != 2 {
		fmt.Printf(error)
		return
	}

	start, errstart := strconv.Atoi(args[0])
	stop, errstop := strconv.Atoi(args[1])

	if errstart != nil {
		fmt.Printf(error3, start)
		return
	}
	if errstop != nil {
		fmt.Printf(error3, stop)
	}

	if start > stop {
		fmt.Printf(error2)
		return
	} else if start < 0 {
		fmt.Printf(error2)
		return
	} else if stop > len(ships) {
		fmt.Printf(error2)
		return
	} else {
		fmt.Printf("%v\n", ships[start:stop])
	}
}
