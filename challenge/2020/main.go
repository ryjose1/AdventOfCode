package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ryjose1/advent2020/trees"
)

func main() {
	file, err := os.Open("inputs/day3.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	input := trees.ScanFileLines(file)
	slope1 := trees.EncounteredTrees(input, 1, 1)
	slope2 := trees.EncounteredTrees(input, 3, 1)
	slope3 := trees.EncounteredTrees(input, 5, 1)
	slope4 := trees.EncounteredTrees(input, 7, 1)
	slope5 := trees.EncounteredTrees(input, 1, 2)
	answer := slope1 * slope2 * slope3 * slope4 * slope5
	fmt.Printf("Hit %d trees", answer)
}
