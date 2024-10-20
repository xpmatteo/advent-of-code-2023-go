package advent_day16

type Row struct {
	cells []*Cell
}

func newRow1(s string) *Row {
	cells := []*Cell{}
	for _, _ = range s {
		cells = append(cells, newEmptyCell())
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
