package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	args := os.Args
	l := len(args) - 1
	if l < 1 {
		fmt.Println("Give me a month name")
		return
	}

	year := time.Now().Year()
	var leap bool
	if year%4 == 0 && (year%100 != 0 || year%400 == 0) {
		leap = true
	}

	days := 28

	month := os.Args[1]

	if m := strings.ToLower(month); m == "april" ||
		m == "june" ||
		m == "september" ||
		m == "november" {
		days = 30
	} else if m == "january" ||
		m == "march" ||
		m == "may" ||
		m == "july" ||
		m == "august" ||
		m == "october" ||
		m == "december" {
		days = 31
	} else if m == "february" {
		if leap {
			days = 29
		}
	} else {
		fmt.Printf("%q is not a month.\n", month)
		return
	}

	fmt.Printf("%q has %d days.\n", month, days)
}
