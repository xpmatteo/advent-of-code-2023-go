package advent_day16

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func Test_emptyCell_L(t *testing.T) {
	cell := newEmptyCell()

	result := cell.Enter(L)

	assert.Equal(t, setOf(R), result)
	assert.True(t, cell.energized, "energized")
}

func Test_rowEmptyCells(t *testing.T) {
	row := newRow1("...")
	assert.Equal(t, "...", row.String())

	row.cells[0].Enter(L)

	assert.Equal(t, "###", row.String())
}

func Test_rowEmptyCells_left(t *testing.T) {
	row := newRow1("...")
	assert.Equal(t, "...", row.String())

	row.cells[2].Enter(R)

	assert.Equal(t, "###", row.String())
}

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
			input:    `.v.`,
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
		input := removeWhiteSpace(test.input)
		name := input
		if test.name != "" {
			name = test.name
		}
		require.Equal(t, 1, len(test.enterDir), "no enterDir set")
		t.Run(name, func(t *testing.T) {
			wants := removeWhiteSpace(test.wants)
			row := newRow1(input)
			row.cells[test.col].Enter(test.enterDir)
			assert.Equal(t, wants, row.String())
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
