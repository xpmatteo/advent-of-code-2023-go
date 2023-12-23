package day10

import "testing"

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

func TestGoingEast(t *testing.T) {
	m := NewMap(simpleSample)
	newRow, newCol, newDir := m.Go(1, 1, E)
	if newRow != 1 || newCol != 2 || newDir != E {
		t.Errorf("Expected to go east, but was at %d, %d, %v", newRow, newCol, newDir)
	}
}

func TestGoingSouth(t *testing.T) {
	m := NewMap(simpleSample)
	newRow, newCol, newDir := m.Go(1, 2, E)
	if newRow != 1 || newCol != 2 || newDir != S {
		t.Errorf("Expected to go south, but was at %d, %d, %v", newRow, newCol, newDir)
	}
}
