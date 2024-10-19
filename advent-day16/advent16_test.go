package advent_day16

import (
	"github.com/stretchr/testify/assert"
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
	empty0 := newEmptyCell()
	empty1 := newEmptyCell()
	empty2 := newEmptyCell()
	row := newRow(&empty0, &empty1, &empty2)
	assert.Equal(t, "...", row.String())

	empty0.Enter(L)

	assert.True(t, empty1.energized, "propagate ray of light")
	assert.Equal(t, "###", row.String())
}

func xTest_rayTracing(t *testing.T) {
	tests := []struct {
		input string
		wants string
	}{
		{
			input: `>`,
			wants: `#`,
		},
		{
			input: `>..`,
			wants: `###`,
		},
		{
			input: `.v.`,
			wants: `.#.`,
		},
		{
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
		input := removeWhiteSpace(test.input)
		t.Run(input, func(t *testing.T) {
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
