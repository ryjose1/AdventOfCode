package password

import (
	"strings"
	"testing"
)

var filterValidLinesTests = []struct {
	input       []string
	expected    []string
	description string
}{
	{
		[]string{
			"1-3 a: abcde",
			"1-3 b: cdefg",
			"2-9 c: ccccccccc",
		},
		[]string{
			"1-3 a: abcde",
		}, "empty string",
	},
}

func TestFilterValidLines(t *testing.T) {
	for _, test := range filterValidLinesTests {
		if actual := FilterValidLines(test.input); len(actual) != len(test.expected) {
			t.Errorf("FAIL %s - FilterValidLines(input) = %s, expected %s.",
				test.description, strings.Join(actual, "\n"), strings.Join(test.expected, "\n"))
		}
		t.Logf("PASS RunLengthEncode - %s", test.description)
	}

}
