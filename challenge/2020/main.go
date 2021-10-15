package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ryjose1/advent2020/passport"
)

func main() {
	file, err := os.Open("inputs/day4.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	input := passport.ScanFileLines(file)

	answer := passport.CountValid(input)
	fmt.Printf("Found %d valid passports", answer)
}
