package advent_day16

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func Test_rayTracing(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		wants    string
		row, col int
		enterDir Direction
	}{
		{
			enterDir: L,
			input:    `...`,
			wants:    `###`,
		},
		{
			enterDir: R,
			col:      2,
			input:    `...`,
			wants:    `###`,
		},
		{
			enterDir: T,
			col:      1,
			input:    `...`,
			wants:    `.#.`,
		},
		{
			name:     "empty space, 3 rows",
			enterDir: L,
			row:      1,
			input: `...
		            ...
		            ...`,
			wants: `...
		            ###
		            ...`,
		},
		{
			name:     "empty space from top",
			col:      1,
			enterDir: T,
			input: `...
		            ...
		            ...`,
			wants: `.#.
		            .#.
		            .#.`,
		},
		{
			name:     "empty space from top",
			row:      2,
			col:      1,
			enterDir: B,
			input: `...
		            ...
		            ...`,
			wants: `.#.
		            .#.
		            .#.`,
		},
		// etc...
	}
	for _, test := range tests {
		input := removeWhiteSpace(test.input)
		name := input
		if test.name != "" {
			name = test.name
		}
		require.Equal(t, 1, len(test.enterDir), "no enterDir set")
		t.Run(name, func(t *testing.T) {
			wants := removeWhiteSpace(test.wants)
			grid := newGrid(input)
			grid.Enter(test.row, test.col, test.enterDir)
			assert.Equal(t, wants, grid.String())
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
