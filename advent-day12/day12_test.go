package advent_day12

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_singleMatch(t *testing.T) {
	tests := []struct {
		pattern           string
		length            int
		expectedRemainder string
		expectedOk        bool
	}{
		// record too short
		{"", 1, "", false},
		{"#", 2, "", false},
		{"?", 2, "", false},
		{"##", 3, "", false},
		{"?#", 3, "", false},
		{"#?", 3, "", false},

		// record is as long as group
		{"#", 1, "", true},
		{"?", 1, "", true},
		{".", 1, "", false},

		{"##", 2, "", true},
		{"??", 2, "", true},
		{"#?", 2, "", true},
		{"?#", 2, "", true},
		{"#.", 2, "", false},
		{"?.", 2, "", false},

		{"#.", 1, "", true},
		{"?.", 1, "", true},

		// skip ? prefix
		{"?#", 1, "", true},
		{"??#", 1, "#", true},
		{"???#", 1, "?#", true},
		{"?#?", 1, "", true},
		{"??#?", 1, "#?", true},
		{"???#?", 1, "?#?", true},

		{"..", 1, "", false},
		{"##", 1, "", false},

		{"#.#", 1, "#", true},
		{"?.#", 1, "#", true},
		{"??#", 1, "#", true},
		{"..#", 1, "", true},
		{"###", 1, "", false},
		{"?##", 1, "", false},
	}
	for _, test := range tests {
		t.Run(test.pattern, func(t *testing.T) {
			assert := assert.New(t)

			remainder, ok := singleMatch(test.pattern, test.length)

			if !assert.Equal(test.expectedOk, ok, "Expected ok") {
				return
			}
			if test.expectedOk {
				assert.Equal(test.expectedRemainder, remainder)
			}
		})
	}
}

func Test_waysToMatchASingleGroup(t *testing.T) {
	tests := []struct {
		record      string
		groupLength int
		expected    []string
	}{
		{"", 1, []string{}},
		{"?", 1, []string{""}},
		{"??", 1, []string{"", ""}},
		{"???", 1, []string{"?", "", ""}},
		{"??.##", 1, []string{".##", "##"}},
		{"???.##", 1, []string{"?.##", ".##", "##"}},
		{"?.?", 1, []string{"?", ""}},
		{"?....?", 1, []string{"...?", ""}},
	}
	for _, test := range tests {
		t.Run(test.record, func(t *testing.T) {
			assert := assert.New(t)

			assert.Equal(test.expected, waysToMatchASingleGroup(test.record, test.groupLength))
		})
	}
}

func Test_multiGroup(t *testing.T) {
	tests := []struct {
		record   string
		groups   []int
		expected int
	}{
		{"??", []int{1}, 2},
		{"???.###", []int{1, 1, 3}, 1},
		{"..??...?##.", []int{1, 3}, 2},
		{".??..??...?##.", []int{1, 1, 3}, 4},
		{"#?#?#?", []int{6}, 1},
		{"?#?#?#?", []int{6}, 2},
		{"#?#?#?#?", []int{1, 6}, 1},
		{"?#?#?#?#?", []int{1, 6}, 1},
		{"?#?#?#?#?#?#?", []int{3, 1, 6}, 1},
		{"?#?#?#?#?#?#?#?", []int{1, 3, 1, 6}, 1},
	}
	for _, test := range tests {
		t.Run(test.record, func(t *testing.T) {
			assert := assert.New(t)

			assert.Equal(test.expected, multiGroup(test.record, test.groups))
		})
	}
}
