package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/ryjose1/advent2020/expense"
)

func readFileLines(filename string) []int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	input := []int{}
	for scanner.Scan() {
		line := scanner.Text()
		entry, err := strconv.Atoi(string(line))
		if err != nil {
			panic("Couldn't parse line in file")
		}
		input = append(input, entry)
	}

	return input
}

func main() {
	input := readFileLines("input1.txt")
	entry1, entry2, entry3 := expense.FindThreeEntries(input, 2020)
	fmt.Println(fmt.Sprintf("%d, %d, %d answer = %d", entry1, entry2, entry3, entry1*entry2*entry3))
}
