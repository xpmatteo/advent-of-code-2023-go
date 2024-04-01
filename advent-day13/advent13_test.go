package advent_day13

import (
	"github.com/stretchr/testify/assert"
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
