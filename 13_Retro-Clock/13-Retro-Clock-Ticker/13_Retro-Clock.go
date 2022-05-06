package main

import (
	"fmt"
	"time"

	"github.com/inancgumus/screen"
)

func main() {
	screen.Clear()
	ti := 0
	turn := false
	for {
		screen.Clear()
		screen.MoveTopLeft()

		now := time.Now()
		hour, min, sec := now.Hour(), now.Minute(), now.Second()

		clock := []placeholder{
			digits[hour/10], digits[hour%10],
			colon,
			digits[min/10], digits[min%10],
			colon,
			digits[sec/10], digits[sec%10],
		}

		if !turn {
			clock = clock[ti:]
		} else {
			for c := 0; c < ti; c++ {
				clock[c] = empty
			}
		}
		for line := range clock[0] {
			for index, digit := range clock {
				next := clock[index][line]
				if digit == colon && sec%2 == 0 {
					next = "   "
				}
				fmt.Print(next, "  ")
			}
			fmt.Println()
		}

		time.Sleep(time.Second)

		if !turn {
			ti++
			if ti == 8 {
				ti = 7
				turn = true
			}
		} else {
			ti--
			if ti == -1 {
				ti = 0
				turn = false
			}
		}
	}
}
