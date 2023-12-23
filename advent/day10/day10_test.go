package day10

import (
	"fmt"
	"testing"
)

const simpleSample = `.....
.S-7.
.|.|.
.L-J.
.....`

func TestMapAt(t *testing.T) {
	m := NewMap(simpleSample)
	actual := m.At(3, 1)
	expected := "L"
	if actual != expected {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestFurthestPlaceSimpleSample(t *testing.T) {
	m := NewMap(simpleSample)
	row, col, distance, err := m.FurthestPlace(1, 1, S, E)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	expectedRow, expectedCol, expectedDistance := 3, 3, 4
	if row != expectedRow || col != expectedCol || distance != expectedDistance {
		t.Errorf("Expected %d,%d at distance %d, got %d,%d at distance %d", expectedRow, expectedCol, expectedDistance, row, col, distance)
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
			aMap := NewMap(simpleSample)
			newRow, newCol, newDir, err := aMap.Go(test.startingRow, test.startingColumn, test.startingDirection)
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			} else if newRow != test.expectedNewRow || newCol != test.expectedNewColumn || newDir != test.expectedNextDirection {
				t.Errorf("Eexpected to enter %d,%d from %v, but got to %d,%d from %v", test.expectedNewRow, test.expectedNewColumn, test.expectedNextDirection, newRow, newCol, newDir)
			}
		})
	}
}
