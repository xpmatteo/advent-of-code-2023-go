package day11

import (
	"fmt"
	"strings"
)

type Coordinate struct {
	row, col int
}

type StarField struct {
	stars  []Coordinate
	maxRow int
	maxCol int
}

func (sf *StarField) expandEmptyRows() {
	increment := 0
	for row := 0; row <= sf.maxRow; row++ {
		if sf.isEmptyRow(row) {
			increment++
		} else {
			sf.shiftStarsDown(increment, row)
		}
	}
}

func (sf *StarField) expandEmptyCols() {
	for col := 0; col <= sf.maxCol; col++ {
		if sf.isEmptyCol(col) {
			sf.shiftStarsRight(col)
			sf.maxCol++
			col++
		}
		fmt.Println("col ", col)
		fmt.Println(sf)
	}
}

func (sf *StarField) isEmptyRow(row int) bool {
	for _, star := range sf.stars {
		if star.row == row {
			return false
		}
	}
	return true
}

func (sf *StarField) isEmptyCol(col int) bool {
	for _, star := range sf.stars {
		if star.col == col {
			return false
		}
	}
	return true
}

func (sf *StarField) shiftStarsDown(increment int, row int) {
	for i, star := range sf.stars {
		if star.row == row {
			sf.stars[i].row += increment
		}
	}
}

func (sf *StarField) shiftStarsRight(col int) {
	for i, star := range sf.stars {
		if star.col > col {
			sf.stars[i].col++
		}
	}
}

func NewStarField(data string) *StarField {
	sf := &StarField{}
	rows := strings.Split(data, "\n")
	for row, line := range rows {
		for col, char := range line {
			if char == '#' {
				sf.addStar(row, col)
			}
		}
	}
	return sf
}

func (sf *StarField) addStar(row int, col int) {
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
	for row := 0; row <= sf.maxRow; row++ {
		for col := 0; col <= sf.maxCol; col++ {
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

func (sf *StarField) hasStar(row int, col int) bool {
	for _, star := range sf.stars {
		if star.row == row && star.col == col {
			return true
		}
	}
	return false
}
