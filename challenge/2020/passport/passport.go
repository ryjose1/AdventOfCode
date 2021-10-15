package passport

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func validateYear(year string, minYear, maxYear int) bool {
	if len(year) != 4 {
		return false
	}
	y, err := strconv.Atoi(year)
	if err != nil {
		return false
	}
	if y < minYear || y > maxYear {
		return false
	}
	return true
}

func validateHeight(height string) bool {
	var minHeight, maxHeight int

	if strings.HasSuffix(height, "cm") {
		minHeight = 150
		maxHeight = 193
		height = strings.TrimSuffix(height, "cm")
	} else if strings.HasSuffix(height, "in") {
		minHeight = 59
		maxHeight = 76
		height = strings.TrimSuffix(height, "in")
	} else {
		return false
	}

	h, err := strconv.Atoi(height)
	if err != nil {
		return false
	}
	if h < minHeight || h > maxHeight {
		return false
	}

	return true
}

func validateHair(color string) bool {
	if !strings.HasPrefix(color, "#") {
		return false
	}

	color = strings.TrimPrefix(color, "#")
	if len(color) != 6 {
		return false
	}
	for _, r := range color {
		if !(unicode.IsDigit(r) || strings.ContainsRune("abcdef", r)) {
			return false
		}
	}
	return true
}

func validateEyes(color string) bool {
	validColors := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
	for _, validColor := range validColors {
		if color == validColor {
			return true
		}
	}
	return false
}

func validateID(id string) bool {
	if len(id) != 9 {
		return false
	}
	for _, r := range id {
		if !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}

func isValid(passport map[string]string) bool {
	expectedKeys := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	for _, key := range expectedKeys {
		value, found := passport[key]
		if !found {
			return false
		}
		switch key {
		case "byr":
			if !validateYear(value, 1920, 2002) {
				return false
			}
		case "iyr":
			if !validateYear(value, 2010, 2020) {
				return false
			}
		case "eyr":
			if !validateYear(value, 2020, 2030) {
				return false
			}
		case "hgt":
			if !validateHeight(value) {
				return false
			}
		case "hcl":
			if !validateHair(value) {
				return false
			}
		case "ecl":
			if !validateEyes(value) {
				return false
			}
		case "pid":
			if !validateID(value) {
				return false
			}
		}
	}
	return true
}

func parsePassport(data []string) map[string]string {
	passport := make(map[string]string)
	for _, field := range data {
		fieldData := strings.Split(field, ":")
		key, value := fieldData[0], fieldData[1]
		passport[key] = value
	}

	return passport
}

func CountValid(input [][]string) (count int) {
	for _, passportData := range input {
		if isValid(parsePassport(passportData)) {
			count++
		}
	}
	return
}

func ScanFileLines(file *os.File) [][]string {
	scanner := bufio.NewScanner(file)
	input := [][]string{}
	passport := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		if line == "" {
			input = append(input, passport)
			fmt.Printf("Passport Data: %s\n\n", strings.Join(passport, " "))
			passport = []string{}
		} else {
			passport = append(passport, strings.Fields(line)...)
		}
	}
	input = append(input, passport)
	fmt.Printf("Passport Data: %s\n\n", strings.Join(passport, " "))
	return input
}
