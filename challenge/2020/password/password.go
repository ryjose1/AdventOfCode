// Day Two challenge
package password

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isValidCount(minCount int, maxCount int, password string, character string) bool {
	actualCount := strings.Count(password, character)
	return minCount <= actualCount && actualCount <= maxCount
}

func isValidPos(startIndex int, endIndex int, password string, character string) bool {
	isAtFirstPos := string(password[startIndex]) == character
	isAtSecondPos := string(password[endIndex]) == character
	return (isAtFirstPos || isAtSecondPos) && !(isAtFirstPos && isAtSecondPos)
}

func isValid(entry string) bool {
	fields := strings.Fields(entry)
	validRange := strings.SplitN(fields[0], "-", 2)
	num1, err := strconv.Atoi(validRange[0])
	if err != nil {
		fmt.Println("Could not convert num1 to int")
		return false
	}
	num2, err := strconv.Atoi(validRange[1])
	if err != nil {
		fmt.Println("Could not convert num2 to int")
		return false
	}
	character := strings.TrimRight(fields[1], ":")
	password := fields[2]

	return isValidPos(num1-1, num2-1, password, character)
}

func FilterValidLines(input []string) []string {
	valid := []string{}
	for _, entry := range input {
		if isValid(entry) {
			valid = append(valid, entry)
		}
	}
	return valid
}

func ScanFileLines(file *os.File) []string {
	scanner := bufio.NewScanner(file)
	input := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		input = append(input, line)
	}
	return input
}
