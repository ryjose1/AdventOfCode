package DayOne

import (
	"fmt"
	"strings"
	"strconv"
	"advent-code/challenge/DayOne/utils"
)

type DayOneA struct {
	Input string
}

func (d DayOneA) DoChallenge() string {
	inputLines:= strings.Split(d.Input, "\n")

	var sum int
	for _, inputLine := range inputLines {
		frequencyChange := utils.ParseInputNumber(inputLine)
		sum += frequencyChange
	}
	
	fmt.Printf("The final sum is %d\n", sum)

	return strconv.Itoa(sum)
}

