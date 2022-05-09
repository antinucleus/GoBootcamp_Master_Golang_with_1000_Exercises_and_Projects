package main

import (
	s "github.com/inancgumus/prettyslice"
)

func main() {
	s.PrintBacking = true // prints the backing array
	s.MaxPerLine = 10     // prints 10 slice elements per line
	s.Width = 60          // prints 60 character per line

	names := make([]string, 5)
	s.Show("1st step", names)

	names = append(names, "einstein", "tesla", "aristotle")
	s.Show("2nd step", names)

	names = make([]string, 0, 5)
	names = append(names, "einstein", "tesla", "aristotle")
	s.Show("3rd step", names)

	moreNames := [...]string{"plato", "khayyam", "ptolemy"}
	names = append(names, "", "")
	copy(names[3:], moreNames[:])
	s.Show("4th step", names)

	clone := make([]string, 3, 5)
	s.Show("5th step (before append)", clone)
	copy(clone, names[2:])
	clone = append(clone, names[:2]...)
	s.Show("5th step (after append)", clone)

	sliced := clone[2:]
	sliced = append(sliced, "hypatia")

	clone[2] = "elder"

	s.Show("6th step", clone, sliced)

}
