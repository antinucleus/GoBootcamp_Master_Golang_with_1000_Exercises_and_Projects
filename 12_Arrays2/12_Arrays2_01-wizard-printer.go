package main

import (
	"fmt"
	"strings"
)

func main() {

	scientist := [...][3]string{
		{"Albert", "Einstein", "time"},
		{"Isaac", "Newton", "apple"},
		{"Stephen", "Hawking", "blackhole"},
		{"Marie", "Curie", "radium"},
		{"Charles", "Darwin", "fittest"},
	}

	fmt.Printf("%-20s%-20s%-20s\n", "First Name", "Last Name", "Nickname")
	seperator := strings.Repeat("=", 50)
	fmt.Printf("%30s\n", seperator)

	for _, f := range scientist {
		for _, s := range f {
			fmt.Printf("%-20s", s)
		}
		fmt.Println()
	}
}
