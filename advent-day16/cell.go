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

type CellFunc func(direction Direction) Direction

type Cell struct {
	f         CellFunc
	energized bool
	neighbors map[Direction]*Cell
}

func (c *Cell) Enter(enterDir Direction) DirSet {
	c.energized = true
	exitDir := c.f(enterDir)
	if n, ok := c.neighbors[exitDir]; ok {
		n.Enter(exitDir.Opposite())
	}
	return setOf(enterDir.Opposite())
}

func newEmptyCell() *Cell {
	return &Cell{
		f:         func(dir Direction) Direction { return dir.Opposite() },
		energized: false,
		neighbors: make(map[Direction]*Cell),
	}
}

func newBackslashMirror() *Cell {
	f := func(dir Direction) Direction {
		switch dir {
		case L:
			return B
		case B:
			return L
		case R:
			return T
		case T:
			return R
		default:
			panic("bad direction " + dir)
		}
	}
	return &Cell{
		f:         f,
		energized: false,
		neighbors: make(map[Direction]*Cell),
	}
}
