package expense

import "testing"

var findTwoEntriesTests = []struct {
	input       []int
	expected1   int
	expected2   int
	description string
}{
	{[]int{1721, 979, 366, 299, 675, 1456}, 299, 1721, "empty string"},
}

var findThreeEntriesTests = []struct {
	input       []int
	expected1   int
	expected2   int
	expected3   int
	description string
}{
	{[]int{1721, 979, 366, 299, 675, 1456}, 366, 675, 979, "empty string"},
}

func TestFindTwoEntries(t *testing.T) {
	for _, test := range findTwoEntriesTests {
		if actual1, actual2, found := FindTwoEntries(test.input, 2020); !found {
			t.Errorf("FAIL %s - RunLengthEncode(input) = %d, %d, expected %d, %d.",
				test.description, actual1, actual2, test.expected1, test.expected2)
		}
		t.Logf("PASS RunLengthEncode - %s", test.description)
	}

}

func TestFindThreeEntries(t *testing.T) {
	for _, test := range findThreeEntriesTests {
		if actual1, actual2, actual3 := FindThreeEntries(test.input, 2020); actual1 != test.expected1 || actual2 != test.expected2 || actual3 != test.expected3 {
			t.Errorf("FAIL %s - RunLengthEncode(input) = %d, %d, %d, expected %d, %d, %d.",
				test.description, actual1, actual2, actual3, test.expected1, test.expected2, test.expected3)
		}
		t.Logf("PASS RunLengthEncode - %s", test.description)
	}

}
