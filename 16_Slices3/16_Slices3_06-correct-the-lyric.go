package main

import (
	"fmt"
	"strings"
)

func main() {
	lyric := strings.Fields(`all my troubles seemed so far away oh I believe in yesterday now it looks as though they are here to stay`)
	lyric = append([]string{"yesterday"}, lyric...)
	lyric = append(lyric, lyric[8:13]...)
	lyric = append(lyric[:8], lyric[13:]...)
	fmt.Printf("%s\n", lyric)
}
