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

func TestMap_Go(t *testing.T) {
	tests := []struct {
		fromDirection     Direction
		enteringRow       int
		enteringCol       int
		expectedDirection Direction
		expectedRow       int
		expectedCol       int
	}{
		{W, 1, 2, W, 1, 3},
		{W, 1, 3, N, 2, 3},
		{N, 2, 3, N, 3, 3},
	}
	for _, test := range tests {
		name := fmt.Sprintf("After entering %d,%d from %v", test.enteringRow, test.enteringCol, test.fromDirection)
		t.Run(name, func(t *testing.T) {
			aMap := NewMap(simpleSample)
			newRow, newCol, newDir, err := aMap.Go(test.enteringRow, test.enteringCol, test.fromDirection)
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			} else if newRow != test.expectedRow || newCol != test.expectedCol || newDir != test.expectedDirection {
				t.Errorf("Eexpected to enter %d,%d from %v, but got to %d,%d from %v", test.expectedRow, test.expectedCol, test.expectedDirection, newRow, newCol, newDir)
			}
		})
	}
}
