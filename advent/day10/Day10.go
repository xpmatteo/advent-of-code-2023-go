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
	switch dir {
	case E:
		col++
	case S:
		row++
	case W:
		col--
	case N:
		row--
	default:
		return row, col, dir, errors.New(fmt.Sprintf("Unknown direction %v", dir))
	}

	symbol := m.At(row, col)
	x := Cases{dir, symbol}
	switch x {
	case Cases{E, "-"}:
		return row, col, E, nil
	case Cases{W, "-"}:
		return row, col, W, nil
	case Cases{E, "7"}:
		return row, col, S, nil
	case Cases{N, "7"}:
		return row, col, W, nil
	case Cases{W, "F"}:
		return row, col, S, nil
	case Cases{N, "F"}:
		return row, col, E, nil
	case Cases{S, "|"}:
		return row, col, S, nil
	case Cases{N, "|"}:
		return row, col, N, nil
	case Cases{S, "J"}:
		return row, col, W, nil
	case Cases{E, "J"}:
		return row, col, N, nil
	case Cases{W, "L"}:
		return row, col, N, nil
	case Cases{S, "L"}:
		return row, col, E, nil
	}
	return row, col, dir, errors.New(fmt.Sprintf("Unknown symbol/direction %v/%v", symbol, dir))
}

func (m *Map) At(row int, column int) string {
	return m.rows[row][column : column+1]
}

func (m *Map) String() string {
	return strings.Join(m.rows, "\n")
}

func (m *Map) FurthestPlace(startingRow int, startingColumn int, dir0 Direction, dir1 Direction) (int, int, int, error) {
	distance := 1
	row0, col0, dir0, err0 := m.Go(startingRow, startingColumn, dir0)
	row1, col1, dir1, err1 := m.Go(startingRow, startingColumn, dir1)
	for err0 == nil && err1 == nil && (row0 != row1 || col0 != col1) {
		row0, col0, dir0, err0 = m.Go(row0, col0, dir0)
		row1, col1, dir1, err1 = m.Go(row1, col1, dir1)
		distance++
	}
	if err0 != nil || err1 != nil {
		return 0, 0, 0, errors.New(fmt.Sprintf("Error: %v, %v", err0, err1))
	}
	return row0, col0, distance, nil
}

func (m *Map) set(row int, column int, s string) {
	m.rows[row] = m.rows[row][:column] + s + m.rows[row][column+1:]
}

func (m *Map) Mark(row int, column int, mark int) {
	m.set(row, column, fmt.Sprintf("%d", mark))
}

func (m *Map) CleanUp(startingRow int, startingColumn int, dir0 Direction, dir1 Direction) *Map {
	result := NewMap(m.String())
	for row := 0; row < len(m.rows); row++ {
		runes := "-|JFL7"
		for i := 0; i < len(runes); i++ {
			result.rows[row] = strings.ReplaceAll(result.rows[row], runes[i:i+1], ".")
		}
	}
	copyLoop(result, m, startingRow, startingColumn, dir0, dir1)
	return result
}

func (m *Map) Area(startingRow int, startingColumn int, dir0 Direction, dir1 Direction) (int, error) {
	return 100, nil
}

func copyLoop(result *Map, m *Map, startingRow int, startingColumn int, dir0 Direction, dir1 Direction) {
	row0, col0, dir0, err0 := m.Go(startingRow, startingColumn, dir0)
	row1, col1, dir1, err1 := m.Go(startingRow, startingColumn, dir1)
	result.set(row0, col0, m.At(row0, col0))
	result.set(row1, col1, m.At(row1, col1))
	for err0 == nil && err1 == nil && (row0 != row1 || col0 != col1) {
		row0, col0, dir0, err0 = m.Go(row0, col0, dir0)
		row1, col1, dir1, err1 = m.Go(row1, col1, dir1)
		result.set(row0, col0, m.At(row0, col0))
		result.set(row1, col1, m.At(row1, col1))
	}
}

func NewMap(input string) *Map {
	m := &Map{}
	m.Load(input)
	return m
}
