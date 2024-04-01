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
2 #....#..# 2
3 ..##..### 3
4v#####.##.v4
5^#####.##.^5
6 ..##..### 6
7 #....#..# 7`

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
		//{pattern1, 5},
		//{pattern2, 400},
	}
	for _, test := range testCases {
		t.Run(test.pattern, func(t *testing.T) {
			assert.Equal(t, test.score, score(test.pattern))
		})
	}
}

func Test_isPalyndrome(t *testing.T) {
	var testCases = []struct {
		pattern      string
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
		t.Run(test.pattern, func(t *testing.T) {
			assert.Equal(t, test.isPalyndrome, isPalyndrome(test.pattern))
		})
	}
}

func isPalyndrome(pattern string) bool {
	for i := 0; i < len(pattern)/2; i++ {
		if pattern[i] != pattern[len(pattern)-1-i] {
			return false
		}
	}
	return true
}
func score(pattern string) int {
	if isEven(len(pattern)) && isPalyndrome(pattern) {
		return len(pattern) / 2
	}
	if isEven(len(pattern)) {
		return 0
	}
	if isPalyndrome(pattern[:len(pattern)-1]) {
		return len(pattern) / 2
	}
	if isPalyndrome(pattern[1:]) {
		return len(pattern)/2 + 1
	}
	return 0
}

func isEven(n int) bool {
	return n%2 == 0
}
