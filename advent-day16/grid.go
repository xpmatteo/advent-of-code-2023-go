package advent_day16

import (
	"strings"
)

type Grid struct {
	rows []*Row
}

type Row struct {
	cells []*Cell
}

func newGrid(s string) *Grid {
	var rows = []*Row{}
	// create rows
	for _, s2 := range strings.Split(s, "\n") {
		rows = append(rows, newRow(s2))
	}
	// connect cells
	for iRow, row := range rows {
		if iRow > 0 {
			rowAbove := rows[iRow-1]
			rowBelow := rows[iRow]
			for iCol := range len(row.cells) {
				rowAbove.cells[iCol].neighbors[B] = rowBelow.cells[iCol]
				rowBelow.cells[iCol].neighbors[T] = rowAbove.cells[iCol]
			}
		}
	}
	return &Grid{rows: rows}
}

func (g *Grid) Enter(row, col int, dir Direction) {
	g.rows[row].cells[col].Enter(dir)
}

func (g *Grid) String() string {
	result := ""
	for _, row := range g.rows {
		result += row.String() + "\n"
	}
	return result[:len(result)-1]
}

func newRow(s string) *Row {
	cells := []*Cell{}
	for _, c := range s {
		cells = append(cells, newCell(int(c)))
	}
	for i := 0; i < len(cells)-1; i++ {
		cells[i].neighbors[R] = cells[i+1]
		cells[i+1].neighbors[L] = cells[i]
	}
	return &Row{cells: cells}
}

func (r *Row) String() string {
	result := ""
	for _, c := range r.cells {
		if c.energized {
			result += "#"
		} else {
			result += "."
		}
	}
	return result
}
