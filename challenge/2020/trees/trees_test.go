package trees

import (
	"testing"
)

var encounteredTreesTests = []struct {
	input       []string
	expected    int
	description string
}{
	{
		[]string{
			"..##.......",
			"#...#...#..",
			".#....#..#.",
			"..#.#...#.#",
			".#...##..#.",
			"..#.##.....",
			".#.#.#....#",
			".#........#",
			"#.##...#...",
			"#...##....#",
			".#..#...#.#",
		},
		7, "",
	},
}

func TestEncounteredTrees(t *testing.T) {
	for _, test := range encounteredTreesTests {
		if actual := EncounteredTrees(test.input, 3, 1); actual != test.expected {
			t.Errorf("FAIL %s - Encounteredtrees(input) = %d, expected %d.",
				test.description, actual, test.expected)
		}
		t.Logf("PASS EncounterddTrees - %s", test.description)
	}

}

var isTreeTests = []struct {
	input       string
	expected    bool
	description string
}{
	{"..##.......", true, ""},
	{"#...#...#..", false, ""},
	{".#....#..#.", false, ""},
	{"..#.#...#.#", true, ""},
	{".#...##..#.", false, ""},
	{"..#.##.....", true, ""},
	{".#.#.#....#", false, ""},
	{".#........#", false, ""},
	{"#.##...#...", true, ""},
	{"#...##....#", false, ""},
	{".#..#...#.#", false, ""},
}

func TestIsTree(t *testing.T) {
	for _, test := range isTreeTests {
		if actual := isTree(test.input, 2); actual != test.expected {
			t.Errorf("FAIL %s - IsTree(input) = %t, expected %t.",
				test.description, actual, test.expected)
		}
		t.Logf("PASS IsTree - %s", test.description)
	}

}
