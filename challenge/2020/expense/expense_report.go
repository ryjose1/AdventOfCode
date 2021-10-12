// Day One challenge
// Package expense holds tools to help find info from an expense report
package expense

import (
	"bufio"
	"os"
	"sort"
	"strconv"
)

// FindTwoEntries returns two entries from input that sum up to target, returns false if not found
func FindTwoEntries(input []int, target int) (int, int, bool) {
	sort.Ints(input)
	for _, entry := range input {
		searchFor := target - entry
		index := sort.SearchInts(input, searchFor)
		if index < len(input) && input[index] == searchFor {
			return entry, input[index], true
		}
	}
	return 0, 0, false
}

// FindTwoEntries returns three entries from input that sum up to target, returns zeros if not found
func FindThreeEntries(input []int, target int) (int, int, int) {
	sort.Ints(input)

	for i, entry := range input {
		searchFor := target - entry
		entry2, entry3, found := FindTwoEntries(input[i+1:], searchFor)
		if found {
			return entry, entry2, entry3
		}
	}

	return 0, 0, 0
}

func ScanFileLines(file *os.File) []int {
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
