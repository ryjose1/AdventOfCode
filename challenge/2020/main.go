package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ryjose1/advent2020/password"
)

func main() {
	file, err := os.Open("inputs/day2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	input := password.ScanFileLines(file)
	answer := password.FilterValidLines(input)
	for _, line := range answer {
		fmt.Println(line)
	}
	fmt.Printf("Valid Lines: %d", len(answer))
}
