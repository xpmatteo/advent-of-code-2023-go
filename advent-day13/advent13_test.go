package advent_day13

import (
	"github.com/stretchr/testify/assert"
	"os"
	"strings"
	"testing"
)

const pattern1 = `#.##..##.
..#.##.#.
##......#
##......#
..#.##.#.
..##..##.
#.#.##.#.`

const pattern2 = `#...##..#
#....#..#
..##..###
#####.##.
#####.##.
..##..###
#....#..#`

func Test_verticalSymmetry(t *testing.T) {
	var testCases = []struct {
		pattern string
		score   int
	}{
		{"##", 1},
		{"##.", 1},
		{".##", 2},
		{"####", 2},
		{"#..#", 2},
		{"####.", 2},
		{"#..#.", 2},
		{"#..##", 2},
		{".####", 3},

		{pattern1, 5},
		{pattern2, 400},
	}
	for _, test := range testCases {
		t.Run(test.pattern, func(t *testing.T) {
			assert.Equal(t, test.score, score(NewPattern(test.pattern)))
		})
	}
}

func Test_isPalyndrome(t *testing.T) {
	var testCases = []struct {
		line         Line
		isPalyndrome bool
	}{
		{"##", true},
		{"##.", false},
		{".##", false},
		{"####", true},
		{"#.#.", false},
		//{pattern1, 5},
		//{pattern2, 400},
	}
	for _, test := range testCases {
		t.Run(string(test.line), func(t *testing.T) {
			assert.Equal(t, test.isPalyndrome, test.line.isPalyndromic())
		})
	}
}

func Test_columns(t *testing.T) {
	var testCases = []struct {
		p       string
		columns []Column
	}{
		{"##\n.#\n..", []Column{"#..", "##."}},
	}
	for _, test := range testCases {
		t.Run(test.p, func(t *testing.T) {
			assert.Equal(t, test.columns, NewPattern(test.p).columns())
		})
	}
}

func Test_scoreMany(t *testing.T) {
	var testCases = []struct {
		patterns []string
		score    int
	}{
		{[]string{pattern1, pattern2}, 405},
		{parseFile("day13.txt"), 1407},
	}
	for _, test := range testCases {
		t.Run("", func(t *testing.T) {
			patterns := make([]Pattern, len(test.patterns))
			for i, p := range test.patterns {
				patterns[i] = NewPattern(p)
			}
			assert.Equal(t, test.score, scoreMany(patterns))
		})
	}
}

func parseFile(fileName string) []string {
	bytes, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	actual := strings.Split(string(bytes), "\n\n")

	return actual
}

func scoreMany(patterns []Pattern) int {
	result := 0
	for _, p := range patterns {
		result += score(p)
	}
	return result
}
