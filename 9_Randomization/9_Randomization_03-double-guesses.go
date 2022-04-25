package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const (
	usage = `
Welcome to Lucky Number Game

The program will pick %d random numbers.
Your mission is to guess one those numbers.

The greater your number is, the harder it gets.

Wanna play?
`
	maxTurns = 5
)

func main() {

	rand.Seed(time.Now().UnixNano())

	args := os.Args[1:]
	if len(args) != 2 {
		fmt.Printf(usage, maxTurns)
		return
	}
	guess, err := strconv.Atoi(args[0])
	guess2, err2 := strconv.Atoi(args[1])
	if err != nil || err2 != nil {
		fmt.Println("Not a number.")
		return
	}

	if guess <= 0 || guess2 <= 0 {
		fmt.Println("Please pick a positive number.")
		return
	}

	var max = guess
	if guess2 > guess {
		max = guess2
	}
	for turn := 0; turn < maxTurns; turn++ {
		n := rand.Intn(max + 1)

		if n == guess || n == guess2 {
			if turn == 0 {
				fmt.Println("YOU WIN IN FIRST TIME!")
			} else {
				fmt.Println("YOU WIN!")
			}
			return
		}
	}
	fmt.Println("YOU LOST... Try again?")

}
