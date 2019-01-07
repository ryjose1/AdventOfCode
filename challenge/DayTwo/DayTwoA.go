package DayTwo

import (
	"fmt"
	"strconv"
	"strings"
)

type DayTwoA struct {
	Input string
}

func countLetters(input string) map[string]int {
	letterMap := make(map[string]int)

	for _, letter := range input {
		letterString := string(letter)
		letterMap[letterString] = letterMap[letterString] + 1
	}

	return letterMap
}

func doesAnyLetterRepeat(repeatAmount int, letterCountMap map[string]int) bool {
	for _, letterCount := range letterCountMap {
		if letterCount == repeatAmount {
			return true
		}
	}
	return false
}

func (d DayTwoA) DoChallenge() string { 
	inputLines := strings.Split(d.Input, "\n")

	var twoSameLetterIds, threeSameLetterIds int

	for _, inputLine := range inputLines {
		letterCountMap := countLetters(inputLine)

		if doesAnyLetterRepeat(2, letterCountMap) {
			twoSameLetterIds++
		}

		if doesAnyLetterRepeat(3, letterCountMap) {
			threeSameLetterIds++
		}
	}

	checksum := twoSameLetterIds * threeSameLetterIds

	fmt.Printf("There were %d ids with two of the same letter and %d ids with three of the same letter. \n" + 
		"The checksum is %d \n", twoSameLetterIds * threeSameLetterIds, checksum)
	return strconv.Itoa(checksum)
}
