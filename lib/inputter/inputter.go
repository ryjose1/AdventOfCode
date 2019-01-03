package inputter

import (
	"io/ioutil"
	"fmt"
	"os"
)

func Input() string {
	input, err := ioutil.ReadFile("input.txt")
	
	if err != nil {
		fmt.Println("Error reading input file")
		os.Exit(1)
	}
	
	return string(input)
}


