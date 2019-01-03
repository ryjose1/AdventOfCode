package inputter

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func Input() string {
	input, err := ioutil.ReadFile("input.txt")

	if err != nil {
		fmt.Println("Error reading input file")
		os.Exit(1)
	}
	return strings.TrimSpace(string(input))
}
