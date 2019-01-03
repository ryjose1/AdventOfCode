package DayTwo

import (
	"fmt"
	"strconv"
	"strings"
)

type DayTwoA struct {
	Input string
}

func isTwoLetterId(id string) bool {
	return true
}

func isThreeLetterId(id string) bool {
	return true
}

func (d DayTwoA) DoChallenge() string {
	inputLines := strings.Split(d.Input, "\n")

	twoLetterIds := make([]string, 0, 250)
	threeLetterIds := make([]string, 0, 250)

	for i, inputLine := range inputLines {
		if isTwoLetterId(inputLine) {
			twoLetterIds = append(twoLetterIds, inputLine)
		}
		if isThreeLetterId(inputLine) {
			threeLetterIds = append(threeLetterIds, inputLine)
		}
		fmt.Printf("i:%d, len2: %d, len3: %d\n", i, len(twoLetterIds), len(threeLetterIds))
	}

	checksum := len(twoLetterIds) * len(threeLetterIds)
	return strconv.Itoa(checksum)
}
