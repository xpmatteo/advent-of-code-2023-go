package day10

import (
	"errors"
	"fmt"
	"strings"
)

type Direction struct {
	direct string
}

var (
	N = Direction{"NORTH"}
	E = Direction{"EAST"}
	S = Direction{"SOUTH"}
	W = Direction{"WEST"}
)

// A Map contains an array of strings, each of which is a row of the map.
type Map struct {
	rows []string
}

func (m *Map) Load(input string) {
	m.rows = strings.Split(input, "\n")
}

func (m *Map) Go(row int, col int, dir Direction) (int, int, Direction, error) {
	type Cases struct {
		enteringDirection Direction
		symbol            string
	}
	symbol := m.At(row, col)
	fmt.Printf("At %d,%d, symbol is %v\n", row, col, symbol)
	x := Cases{dir, symbol}
	switch x {
	case Cases{W, "-"}:
		return row, col + 1, W, nil
	case Cases{W, "7"}:
		return row + 1, col, N, nil
	case Cases{N, "|"}:
		return row + 1, col, N, nil
	}
	return row, col, dir, errors.New("not implemented")
}

func (m *Map) At(row int, column int) string {
	return m.rows[row][column : column+1]
}

func NewMap(input string) *Map {
	m := &Map{}
	m.Load(input)
	return m
}
