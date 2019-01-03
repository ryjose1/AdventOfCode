package main

import (
	"advent-code/challenge/DayOne"
	"advent-code/challenge/DayTwo"
	"advent-code/lib/inputter"
	"fmt"
)

type Challenger interface {
	DoChallenge() string
}

func main() {
	input := inputter.Input()

	var challenge Challenger

	challenge = DayOne.DayTwoA{input}
	output := challenge.DoChallenge()
	fmt.Printf("The answer is '%s'\n", output)
}
