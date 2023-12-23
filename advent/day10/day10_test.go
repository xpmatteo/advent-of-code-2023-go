package day10

import "testing"

const simple_sample = `.....
.S-7.
.|.|.
.L-J.
.....`

func TestGoingEast(t *testing.T) {
	m := NewMap(simple_sample)
	newRow, newCol, newDir := m.Go(1, 1, E)
	if newRow != 1 || newCol != 2 || newDir != E {
		t.Errorf("Expected to go east, but was at %d, %d, %v", newRow, newCol, newDir)
	}
}
