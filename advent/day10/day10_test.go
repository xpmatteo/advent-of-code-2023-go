package day10

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

const simpleSample = `.....
.S-7|
||||L
.L-JJ
.....`

const simpleSampleClean = `.....
.S-7.
.|.|.
.L-J.
.....`

const lessSimpleSample = `7-F7-
.FJ|7
SJLL7
|F--J
LJ.LJ`

func TestMapAt(t *testing.T) {
	m := NewMap(simpleSample)

	assert.Equal(t, "L", m.At(3, 1))
}

func TestMapMark(t *testing.T) {
	m := NewMap(simpleSample)

	m.Mark(3, 1, 8)

	assert.Equal(t, "8", m.At(3, 1))
}

func TestFurthestPlaceSimpleSample(t *testing.T) {
	assert := assert.New(t)
	m := NewMap(simpleSample)

	row, col, distance, err := m.FurthestPlace(1, 1, S, E)

	if assert.NoError(err) {
		assert.Equal(3, row)
		assert.Equal(3, col)
		assert.Equal(4, distance)
	}
}

func TestFurthestPlaceLessSimpleSample(t *testing.T) {
	assert := assert.New(t)
	m := NewMap(lessSimpleSample)

	row, col, distance, err := m.FurthestPlace(2, 0, S, E)

	if assert.NoError(err) {
		assert.Equal(2, row)
		assert.Equal(4, col)
		assert.Equal(8, distance)
	}
}

func TestMap_Go(t *testing.T) {
	tests := []struct {
		startingRow           int
		startingColumn        int
		startingDirection     Direction
		expectedNewRow        int
		expectedNewColumn     int
		expectedNextDirection Direction
	}{
		// loop right
		{1, 1, E, 1, 2, E},
		{1, 2, E, 1, 3, S},
		{1, 3, S, 2, 3, S},
		{2, 3, S, 3, 3, W},
		{3, 3, W, 3, 2, W},
		{3, 2, W, 3, 1, N},
		{3, 1, N, 2, 1, N},

		// loop south
		{1, 1, S, 2, 1, S},
		{2, 1, S, 3, 1, E},
		{3, 1, E, 3, 2, E},
		{3, 2, E, 3, 3, N},
		{3, 3, N, 2, 3, N},
		{2, 3, N, 1, 3, W},
		{1, 3, W, 1, 2, W},
	}
	for _, test := range tests {
		name := fmt.Sprintf("After entering %d,%d from %v", test.startingRow, test.startingColumn, test.startingDirection)
		t.Run(name, func(t *testing.T) {
			assert := assert.New(t)
			aMap := NewMap(simpleSample)

			newRow, newCol, newDir := aMap.Go(test.startingRow, test.startingColumn, test.startingDirection)

			assert.Equal(test.expectedNewRow, newRow)
			assert.Equal(test.expectedNewColumn, newCol)
			assert.Equal(test.expectedNextDirection, newDir)
		})
	}
}

func TestAcceptancePart1(t *testing.T) {
	assert := assert.New(t)
	b, err := os.ReadFile("day10.txt") // just pass the file name
	require.NoError(t, err)
	m := NewMap(string(b))

	assert.Equal("S", m.At(57, 65))

	_, _, distance, err := m.FurthestPlace(57, 65, S, N)
	if assert.NoError(err) {
		assert.Equal(6951, distance)
	}
}

func TestMapCleanUp(t *testing.T) {
	assert := assert.New(t)
	m := NewMap(simpleSample)

	actual := m.CleanUp(1, 1, S, E)

	assert.Equal(string(simpleSampleClean), actual.String())
}

func xTestAcceptancePart2(t *testing.T) {
	assert := assert.New(t)
	b, err := os.ReadFile("day10.txt") // just pass the file name
	require.NoError(t, err)
	m := NewMap(string(b)).CleanUp(57, 65, S, N)

	area := m.Area(57, 65, "F", S, N)
	assert.Equal(1, area)
}

const convexElbowDown = `.......
.S-----7.
.|.....|.
.L-7.F-J.
...L-J...`

const concaveElbowDown = `.......
.S-----7.
.|.F-7.|.
.L-J.L-J.`

const convexElbowUp = `.......
...F-7...
.S-J.L-7.
.L-----J.`

const harderAreaSample = `.F----7F7F7F7F-7....
.|F--7||||||||FJ....
.||.FJ||||||||L7....
FJL7L7LJLJ||LJ.L-7..
L--J.L7...LJS7F-7L7.
....F-J..F7FJ|L7L7L7
....L7.F7||L7|.L7L7|
.....|FJLJ|FJ|F7|.LJ
....FJL-7.||.||||...
....L---J.LJ.LJLJ...`

func TestMap_AreaCases(t *testing.T) {
	tests := []struct {
		name           string
		data           string
		startingRow    int
		startingColumn int
		startingSymbol string
		dir0           Direction
		dir1           Direction
		expectedArea   int
	}{
		{"simple", simpleSample, 1, 1, "F", S, E, 1},
		{"convexElbowDown", convexElbowDown, 1, 1, "F", S, E, 6},
		{"concaveElbowDown", concaveElbowDown, 1, 1, "F", S, E, 2},
		{"convexElbowUp", convexElbowUp, 2, 1, "F", S, E, 1},
		//		{"harder", harderAreaSample, 4, 12, S, E, 8},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert := assert.New(t)
			m := NewMap(string(test.data))

			require.Equal(t, "S", m.At(test.startingRow, test.startingColumn), "starting point is wrong")

			area := m.Area(test.startingRow, test.startingColumn, test.startingSymbol, S, E)
			cleaned := m.CleanUp(test.startingRow, test.startingColumn, test.dir0, test.dir1).String()
			assert.Equal(test.expectedArea, area, cleaned)
		})
	}
}
