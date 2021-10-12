package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ryjose1/advent2020/expense"
)

func main() {
	file, err := os.Open("input1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	input := expense.ScanFileLines(file)
	entry1, entry2, entry3 := expense.FindThreeEntries(input, 2020)
	fmt.Println(fmt.Sprintf("%d, %d, %d answer = %d", entry1, entry2, entry3, entry1*entry2*entry3))
}
