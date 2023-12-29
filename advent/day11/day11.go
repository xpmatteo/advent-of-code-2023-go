package day11

import (
	"fmt"
	"strings"
)

type Row int
type Col int

type Coordinate struct {
	row Row
	col Col
}

func (c Coordinate) String() string {
	return fmt.Sprintf("(%d, %d)", c.row, c.col)
}

func (c Coordinate) Distance(other Coordinate) int {
	return abs(int(c.row)-int(other.row)) + abs(int(c.col)-int(other.col))
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

type StarField struct {
	stars  []Coordinate
	maxRow Row
	maxCol Col
}

func (sf *StarField) expandEmptyRows(distance int) {
	for row := Row(0); row <= sf.maxRow; row++ {
		if sf.isEmptyRow(row) {
			sf.shiftStarsDown(distance, row)
			sf.maxRow += Row(distance)
			row += Row(distance)
		}
	}
}

func (sf *StarField) expandEmptyCols(distance int) {
	for col := Col(0); col <= sf.maxCol; col++ {
		if sf.isEmptyCol(col) {
			sf.shiftStarsRight(distance, col)
			sf.maxCol += Col(distance)
			col += Col(distance)
		}
	}
}

func (sf *StarField) isEmptyRow(row Row) bool {
	for _, star := range sf.stars {
		if star.row == row {
			return false
		}
	}
	return true
}

func (sf *StarField) isEmptyCol(col Col) bool {
	for _, star := range sf.stars {
		if star.col == col {
			return false
		}
	}
	return true
}

func (sf *StarField) shiftStarsDown(distance int, row Row) {
	for i, star := range sf.stars {
		if star.row > row {
			sf.stars[i].row += Row(distance)
		}
	}
}

func (sf *StarField) shiftStarsRight(distance int, col Col) {
	for i, star := range sf.stars {
		if star.col > col {
			sf.stars[i].col += Col(distance)
		}
	}
}

func NewStarField(data string) *StarField {
	sf := &StarField{}
	rows := strings.Split(data, "\n")
	for row, line := range rows {
		for col, char := range line {
			if char == '#' {
				sf.addStar(Row(row), Col(col))
			}
		}
	}
	return sf
}

func (sf *StarField) addStar(row Row, col Col) {
	sf.stars = append(sf.stars, Coordinate{row, col})
	if row > sf.maxRow {
		sf.maxRow = row
	}
	if col > sf.maxCol {
		sf.maxCol = col
	}
}

func (sf *StarField) String() string {
	var sb strings.Builder
	for row := Row(0); row <= sf.maxRow; row++ {
		for col := Col(0); col <= sf.maxCol; col++ {
			if sf.hasStar(row, col) {
				sb.WriteRune('#')
			} else {
				sb.WriteRune('.')
			}
		}
		sb.WriteRune('\n')
	}
	return sb.String()
}

func (sf *StarField) hasStar(row Row, col Col) bool {
	for _, star := range sf.stars {
		if star.row == row && star.col == col {
			return true
		}
	}
	return false
}

func (sf *StarField) Expand(replacement int) {
	sf.expandEmptyRows(replacement - 1)
	sf.expandEmptyCols(replacement - 1)
}

func (sf *StarField) SumDistances() int {
	sum := 0
	for _, starA := range sf.stars {
		for _, starB := range sf.stars {
			sum += starA.Distance(starB)
		}
	}
	return sum / 2
}
