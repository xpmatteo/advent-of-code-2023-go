package day11

import "strings"

type Coordinate struct {
	row, col int
}

type StarField struct {
	stars  []Coordinate
	maxRow int
	maxCol int
}

func (sf *StarField) ExpandEmptyRows() {
	increment := 0
	for row := 0; row < 4; row++ {
		if sf.emptyRow(row) {
			increment++
		} else {
			sf.shiftStarsDown(increment, row)
		}
	}
}

func (sf *StarField) emptyRow(row int) bool {
	return row == 1
}

func (sf *StarField) shiftStarsDown(increment int, row int) {
	for i, star := range sf.stars {
		if star.row == row {
			sf.stars[i].row += increment
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
