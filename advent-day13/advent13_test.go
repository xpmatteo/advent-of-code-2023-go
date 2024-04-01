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

const pattern2 = `1 #...##..# 1
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

		{".#.#", 0},
		{".#..###", 0},

		//{"##\n##", 1},

		//{pattern1, 5},
		//{pattern2, 400},
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
			assert.Equal(t, test.isPalyndrome, Line(test.line).isPalyndromic())
		})
	}
}

func score(pattern Pattern) int {
	line := pattern[0]
	if isEven(len(line)) && line.isPalyndromic() {
		return len(line) / 2
	}
	if isEven(len(line)) {
		return 0
	}
	if line[:len(line)-1].isPalyndromic() {
		return len(line) / 2
	}
	if line[1:].isPalyndromic() {
		return len(line)/2 + 1
	}
	return 0
}

func isEven(n int) bool {
	return n%2 == 0
}
