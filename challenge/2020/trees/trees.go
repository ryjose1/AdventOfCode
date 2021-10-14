package trees

import (
	"bufio"
	"os"
)

func isTree(line string, xPos int) bool {
	r := []rune(line)
	if xPos < len(r) && r[xPos] == '#' {
		return true
	}
	return false
}

func EncounteredTrees(input []string, xSlope, ySlope int) int {
	if len(input) <= 1 {
		return 0
	}

	xPos := 0
	trees := 0
	for i, line := range input {
		if i != 0 && i%ySlope == 0 {
			xPos = (xPos + xSlope) % len(line)
			if isTree(line, xPos) {
				trees++
			}
		}

	}
	return trees
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
