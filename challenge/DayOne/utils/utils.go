package utils

import (
	"fmt"
	"strconv"
)

func ParseInputNumber(inputNum string) int {
	parsedNumber, err := strconv.ParseInt(inputNum, 10, 64)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return 0
	}
	
	return int(parsedNumber)
}

