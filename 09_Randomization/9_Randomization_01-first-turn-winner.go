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
	if len(args) != 1 {
		fmt.Printf(usage, maxTurns)
		return
	}
	guess, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("Not a number.")
		return
	}

	if guess <= 0 {
		fmt.Println("Please pick a positive number.")
		return
	}

	for turn := 0; turn < maxTurns; turn++ {
		n := rand.Intn(guess + 1)

		if n == guess {
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
