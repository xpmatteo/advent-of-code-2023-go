package day11

import "strings"

type Coordinate struct {
	row, col int
}

type StarField struct {
	stars []Coordinate
}

func NewStarField(data string) *StarField {
	sf := &StarField{}
	rows := strings.Split(data, "\n")
	for row, line := range rows {
		for col, char := range line {
			if char == '#' {
				sf.stars = append(sf.stars, Coordinate{row, col})
			}
		}
	}
	return sf
}
