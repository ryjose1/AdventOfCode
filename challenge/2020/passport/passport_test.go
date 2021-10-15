package passport

import (
	"testing"
)

var parsePassportTests = []struct {
	input       []string
	expected    map[string]string
	description string
}{
	{
		[]string{
			"iyr:2027",
			"pid:#37f002",
			"hgt:164cm",
			"ecl:#80df11",
			"hcl:#aeacee",
			"cid:320",
			"eyr:2039",
			"byr:1956",
		},
		map[string]string{
			"iyr": "2027",
			"pid": "#37f002",
			"hgt": "164cm",
			"ecl": "#80df11",
			"hcl": "#aeacee",
			"cid": "320",
			"eyr": "2039",
			"byr": "1956",
		}, "empty string",
	},
}

func TestParsePassport(t *testing.T) {
	for _, test := range parsePassportTests {
		actual := parsePassport(test.input)
		for key, value := range test.expected {
			if actual[key] != value {
				t.Errorf("FAIL %s - ParsePassport() key = %s, actual = %s, expected %s.",
					test.description, key, actual[key], value)
			}
		}
		t.Logf("PASS RunLengthEncode - %s", test.description)
	}
}

var isValidTests = []struct {
	input       map[string]string
	expected    bool
	description string
}{
	{
		map[string]string{
			"iyr": "2027",
			"pid": "#37f002",
			"hgt": "164cm",
			"cid": "350",
			"ecl": "#80df11",
			"hcl": "#aeacee",
			"eyr": "2039",
			"byr": "1956",
		},
		true,
		"empty string",
	},
	{
		map[string]string{
			"iyr": "2027",
			"pid": "#37f002",
			"hgt": "164cm",
			"ecl": "#80df11",
			"hcl": "#aeacee",
			"eyr": "2039",
			"byr": "1956",
		},
		true,
		"empty string",
	},
	{
		map[string]string{
			"iyr": "2027",
			"pid": "#37f002",
			"hgt": "164cm",
			"hcl": "#aeacee",
			"eyr": "2039",
			"byr": "1956",
		},
		false,
		"empty string",
	},
}

func TestIsValid(t *testing.T) {
	for _, test := range isValidTests {
		if actual := isValid(test.input); actual != test.expected {
			t.Errorf("FAIL %s - isValid() actual = %t, expected %t.",
				test.description, actual, test.expected)
		}
		t.Logf("PASS isValid - %s", test.description)
	}
}

func TestValidateYear(t *testing.T) {
	var validateYearTests = []struct {
		input       string
		expected    bool
		description string
	}{
		{"2000", true, ""},
		{"1980", false, ""},
		{"2020", false, ""},
	}
	for _, test := range validateYearTests {
		if actual := validateYear(test.input, 1990, 2010); actual != test.expected {
			t.Errorf("FAIL %s - validateYear(%s) actual = %t, expected %t.",
				test.description, test.input, actual, test.expected)
		}
		t.Logf("PASS validateYear - %s", test.description)
	}
}

func TestValidateHeight(t *testing.T) {
	var validateHeightTests = []struct {
		input       string
		expected    bool
		description string
	}{
		{"160cm", true, ""},
		{"140cm", false, ""},
		{"199cm", false, ""},
		{"70in", true, ""},
		{"50in", false, ""},
		{"80in", false, ""},
	}
	for _, test := range validateHeightTests {
		if actual := validateHeight(test.input); actual != test.expected {
			t.Errorf("FAIL %s - validateHeight(%s) actual = %t, expected %t.",
				test.description, test.input, actual, test.expected)
		}
		t.Logf("PASS validateHeight - %s", test.description)
	}
}

func TestValidateHair(t *testing.T) {
	var validateHairTests = []struct {
		input       string
		expected    bool
		description string
	}{
		{"#abcdef", true, ""},
		{"#012987", true, ""},
		{"#012abc", true, ""},
		{"#gh02fi", false, ""},
		{"#abcdef01234", false, ""},
		{"abcdef", false, ""},
		{"#bcdef", false, ""},
	}
	for _, test := range validateHairTests {
		if actual := validateHair(test.input); actual != test.expected {
			t.Errorf("FAIL %s - validateHair(%s) actual = %t, expected %t.",
				test.description, test.input, actual, test.expected)
		}
		t.Logf("PASS validateHair - %s", test.description)
	}
}

func TestValidateEyes(t *testing.T) {
	var validateEyesTests = []struct {
		input       string
		expected    bool
		description string
	}{
		{"amb", true, ""},
		{"blu", true, ""},
		{"brn", true, ""},
		{"gry", true, ""},
		{"grn", true, ""},
		{"hzl", true, ""},
		{"oth", true, ""},
		{"blk", false, ""},
	}
	for _, test := range validateEyesTests {
		if actual := validateEyes(test.input); actual != test.expected {
			t.Errorf("FAIL %s - validateEyes(%s) actual = %t, expected %t.",
				test.description, test.input, actual, test.expected)
		}
		t.Logf("PASS validateEyes - %s", test.description)
	}
}

func TestValidateID(t *testing.T) {
	var validateIDTests = []struct {
		input       string
		expected    bool
		description string
	}{
		{"000000000", true, ""},
		{"000000123", true, ""},
		{"123456789", true, ""},
		{"0010001000", false, ""},
		{"223423", false, ""},
	}
	for _, test := range validateIDTests {
		if actual := validateID(test.input); actual != test.expected {
			t.Errorf("FAIL %s - validateID(%s) actual = %t, expected %t.",
				test.description, test.input, actual, test.expected)
		}
		t.Logf("PASS validateID - %s", test.description)
	}
}
