package DayOne

import (
	"fmt"
	"strings"
	"strconv"
	"advent-code/challenge/DayOne/utils"
)

type DayOneB struct {
	Input string
}

var frequencies map[int]bool

func checkRepeatFrequency(frequency int) bool{
	if frequencies[frequency] {
		return true
	} else {
		frequencies[frequency] = true
		//fmt.Printf("First time seeing frequency %d\n", frequency)
		return false
	}
}

func (d DayOneB) DoChallenge() string {
	d.Input = strings.TrimSpace(d.Input)
	inputLines:= strings.Split(d.Input, "\n")

	frequencies = make(map[int]bool)
	var frequency int	
	isRepeatFrequency := false
	

	for i := 0; !isRepeatFrequency; i++ {
		if i == len(inputLines) {
			i = 0
		}
		frequencyChange := utils.ParseInputNumber(inputLines[i])
		frequency += frequencyChange
		
		isRepeatFrequency = checkRepeatFrequency(frequency)
	}
	
	fmt.Printf("The first repeat frequency is %d\n", frequency)

	return strconv.Itoa(frequency)
}

