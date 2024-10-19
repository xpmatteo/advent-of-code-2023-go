package advent_day16

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

type Direction string

const (
	E Direction = "E"
	W Direction = "W"
	U Direction = "U"
	D Direction = "D"
)

type Set[T comparable] map[T]struct{}

type Cell func(direction Direction) Set[Direction]

func empty(direction Direction) Set[Direction] {
	switch direction {
	case E:
		return setOf(W)
	}
}

func setOf(d ...Direction) Set[Direction] {
	result := make(Set[Direction])
	for _, direction := range d {
		result[direction] = struct{}{}
	}
	return result
}

func Test_rayTracing(t *testing.T) {
	tests := []struct {
		name  string
		input string
		wants string
	}{
		{
			name:  "empty space one cell",
			input: `>`,
			wants: `#`,
		},
		{
			name:  "empty space one row",
			input: `>..`,
			wants: `###`,
		},
		{
			name: "empty space from left",
			input: `...
		           >..
		           ...`,
			wants: `...
		           ###
		           ...`,
		},
		//{
		//	name: "empty space from top",
		//	input: `.v.
		//            ...
		//            ...`,
		//	wants: `.#.
		//            .#.
		//            .#.`,
		//},
		// etc...
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			input := removeWhiteSpace(test.input)
			wants := removeWhiteSpace(test.wants)
			assert.Equal(t, wants, rayTrace(input))
		})
	}
}

func removeWhiteSpace(s string) string {
	s = strings.ReplaceAll(s, " ", "")
	s = strings.ReplaceAll(s, "\t", "")
	return s
}

func rayTrace(input string) string {
	var result []string
	for _, row := range strings.Split(input, "\n") {
		if row[0] == '>' {
			result = append(result, illuminate(row))
		} else {
			result = append(result, row)
		}
	}
	return strings.Join(result, "\n")
}

func illuminate(row string) string {
	result := ""
	for _, _ = range row {
		result += "#"
	}
	return result
}
