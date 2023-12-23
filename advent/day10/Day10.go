package day10

import "strings"

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

func (m *Map) Go(row int, col int, dir Direction) (int, int, Direction) {
	return 1, 2, E
}

func (m *Map) At(row int, column int) string {
	return m.rows[row][column : column+1]
}

func NewMap(input string) *Map {
	m := &Map{}
	m.Load(input)
	return m
}
