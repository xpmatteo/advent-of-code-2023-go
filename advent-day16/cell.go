package advent_day16

type Direction string

func (d Direction) Opposite() Direction {
	switch d {
	case L:
		return R
	case R:
		return L
	case T:
		return B
	case B:
		return T
	default:
		panic("Bad direction " + d)
	}
}

const (
	L Direction = "L"
	R Direction = "R"
	T Direction = "T"
	B Direction = "B"
)

type CellFunc func(direction Direction) DirSet

type Cell struct {
	f         CellFunc
	energized bool
	neighbors map[Direction]*Cell
}

func newEmptyCell() *Cell {
	return &Cell{
		f:         nil,
		energized: false,
		neighbors: make(map[Direction]*Cell),
	}
}

func (c *Cell) Enter(d Direction) DirSet {
	c.energized = true
	if n, ok := c.neighbors[d.Opposite()]; ok {
		n.Enter(d)
	}
	return setOf(d.Opposite())
}
