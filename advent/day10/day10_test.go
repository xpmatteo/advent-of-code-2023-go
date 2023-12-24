package day10

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

const simpleSample = `.....
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

			newRow, newCol, newDir, err := aMap.Go(test.startingRow, test.startingColumn, test.startingDirection)

			if assert.NoError(err) {
				assert.Equal(test.expectedNewRow, newRow)
				assert.Equal(test.expectedNewColumn, newCol)
				assert.Equal(test.expectedNextDirection, newDir)
			}
		})
	}
}

func TestAcceptancePart1(t *testing.T) {
	assert := assert.New(t)
	b, err := os.ReadFile("day10.txt") // just pass the file name
	if !assert.NoError(err) {
		return
	}
	m := NewMap(string(b))

	assert.Equal("S", m.At(57, 65))

	_, _, distance, err := m.FurthestPlace(57, 65, S, N)
	if assert.NoError(err) {
		assert.Equal(6951, distance)
	}
}
