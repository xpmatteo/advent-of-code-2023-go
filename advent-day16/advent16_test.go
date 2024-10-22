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
		{
			name:     "diag \\ mirror from left",
			row:      1,
			col:      0,
			enterDir: L,
			input: `...
		           .\.
		           ...`,
			wants: `...
		           ##.
		           .#.`,
		},
		{
			name:     "diag \\ mirror from right",
			row:      1,
			col:      2,
			enterDir: R,
			input: `...
		           .\.
		           ...`,
			wants: `.#.
		           .##
		           ...`,
		},
		{
			name:     "diag \\ mirror from top",
			row:      0,
			col:      1,
			enterDir: T,
			input: `...
		           .\.
		           ...`,
			wants: `.#.
		           .##
		           ...`,
		},
		{
			name:     "diag / mirror from left",
			row:      1,
			col:      0,
			enterDir: L,
			input: `...
		           ./.
		           ...`,
			wants: `.#.
		           ##.
		           ...`,
		},
		{
			name:     "vert mirror from left",
			row:      1,
			col:      0,
			enterDir: L,
			input: `...
		           .|.
		           ...`,
			wants: `.#.
		           ##.
		           .#.`,
		},
		{
			name:     "vert mirror from right",
			row:      1,
			col:      2,
			enterDir: R,
			input: `...
		           .|.
		           ...`,
			wants: `.#.
		           .##
		           .#.`,
		},
		{
			name:     "vert mirror from top",
			row:      0,
			col:      1,
			enterDir: T,
			input: `...
		           .|.
		           ...`,
			wants: `.#.
		           .#.
		           .#.`,
		},
		{
			name:     "hor mirror from top",
			row:      0,
			col:      1,
			enterDir: T,
			input: `...
		            .-.
		            ...`,
			wants: `.#.
		            ###
		            ...`,
		},
		{
			name:     "hor mirror from bottom",
			row:      2,
			col:      1,
			enterDir: B,
			input: `...
		            .-.
		            ...`,
			wants: `...
		            ###
		            .#.`,
		},
		{
			name:     "loops are ok",
			row:      4,
			col:      3,
			enterDir: B,
			input: `.......
		            ./...\.
					.......
					.\.-./.
		            .......`,
			wants: `.......
		            .#####.
					.#...#.
		            .#####.
		            ...#...`,
		},
		{
			name:     "sample",
			enterDir: L,
			input: `.|...\....
					|.-.\.....
					.....|-...
					........|.
					..........
					.........\
					..../.\\..
					.-.-/..|..
					.|....-|.\
					..//.|....`,
			wants: `######....
					.#...#....
					.#...#####
					.#...##...
					.#...##...
					.#...##...
					.#..####..
					########..
					.#######..
					.#...#.#..`,
		},
	}
	for _, test := range tests {
		input := removeWhiteSpace(test.input)
		require.Equal(t, 1, len(test.enterDir), "no enterDir set")
		t.Run(test.name, func(t *testing.T) {
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

func Test_countIlluminated(t *testing.T) {
	input := `...
			  .-.
		      ...`
	grid := newGrid(removeWhiteSpace(input))
	grid.Enter(0, 1, T)
	assert.Equal(t, 4, grid.CountIlluminated())
}
