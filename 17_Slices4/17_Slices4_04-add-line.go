package main

import (
	"fmt"
	"strings"

	s "github.com/inancgumus/prettyslice"
)

func main() {
	lyric := strings.Fields(`yesterday all my troubles seemed so far away now it looks as though they are here to stay oh I believe in yesterday`)
	fix := make([]string, len(lyric))
	copy(fix, lyric)

	for i := range fix {
		if i == 7 || i == 17 || i == len(fix)-1 {
			fix[i] += "\n"
		}
	}

	s.Show("fix slice", fix)

	for _, w := range fix {
		fmt.Print(w)
		if w != "\n" {
			fmt.Print(" ")
		}
	}
}
