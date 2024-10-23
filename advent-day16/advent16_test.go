package advent_day16

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"os"
	"strings"
	"testing"
)

const sample = `.|...\....
				|.-.\.....
				.....|-...
				........|.
				..........
				.........\
				..../.\\..
				.-.-/..|..
				.|....-|.\
				..//.|....`

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
			input:    sample,
			//      1234567890
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

func Test_countIlluminated_part1(t *testing.T) {
	//grid := newGrid(readFile("input.txt"))
	assert.Equal(t, 46, countEnergized(removeWhiteSpace(sample)))
}

func countEnergized(input string) int {
	grid := newGrid(input)
	grid.Enter(0, 1, T)
	illuminated := grid.CountIlluminated()
	return illuminated
}

func readFile(fileName string) string {
	bytes, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}
