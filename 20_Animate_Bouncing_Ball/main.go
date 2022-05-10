package main

import (
	"fmt"
	"time"

	screen "github.com/inancgumus/screen"
)

// s "github.com/inancgumus/prettyslice"

const (
	height = 10
	width  = 50
)

func main() {
	var (
		x, y   int
		vx, vy = 1, 1
	)
	table := make([][]bool, height)

	for x := range table {
		table[x] = make([]bool, width)
	}

	table[0][0] = true

	for i := 0; i < 120; i++ {
		x += vx
		y += vy
		if x <= 0 || x >= width-1 {
			vx *= -1
		}
		if y <= 0 || y >= height-1 {
			vy *= -1
		}
		for h := range table {
			for w := range table[0] {
				table[h][w] = false
			}
		}
		table[y][x] = true

		buffer := make([]rune, 0, height*width)
		for h := range table {
			for w := range table[0] {
				c := ' '
				if table[h][w] {
					c = 'O'
				}
				buffer = append(buffer, c, ' ')
			}
			buffer = append(buffer, '\n')
		}
		screen.MoveTopLeft()
		fmt.Print(string(buffer))

		time.Sleep(time.Second / 20)
	}

}
