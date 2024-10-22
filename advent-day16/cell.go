package advent_day16

import "fmt"

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

type Cell struct {
	rayTrace  CellFunc
	energized map[Direction]bool
	neighbors map[Direction]*Cell
}

type CellFunc func(direction Direction) DirSet

var cellFuncs = map[int]CellFunc{
	'.':  emptySpaceFunc,
	'/':  slashMirror,
	'\\': backslashMirror,
	'|':  vertMirror,
	'-':  horMirror,
}

func emptySpaceFunc(direction Direction) DirSet {
	return setOf(direction.Opposite())
}

func slashMirror(dir Direction) DirSet {
	switch dir {
	case L:
		return setOf(T)
	case B:
		return setOf(R)
	case R:
		return setOf(B)
	case T:
		return setOf(L)
	default:
		panic("bad direction " + dir)
	}
}

func backslashMirror(dir Direction) DirSet {
	switch dir {
	case L:
		return setOf(B)
	case B:
		return setOf(L)
	case R:
		return setOf(T)
	case T:
		return setOf(R)
	default:
		panic("bad direction " + dir)
	}
}

func vertMirror(dir Direction) DirSet {
	switch dir {
	case L:
		return setOf(T, B)
	case R:
		return setOf(T, B)
	default:
		return setOf(dir.Opposite())
	}
}

func horMirror(dir Direction) DirSet {
	switch dir {
	case T:
		return setOf(L, R)
	case B:
		return setOf(L, R)
	default:
		return setOf(dir.Opposite())
	}
}

func newCell(c int) *Cell {
	cellFunc, ok := cellFuncs[c]
	if !ok {
		panic(fmt.Sprintf("Bad character '%c'", c))
	}
	return &Cell{
		rayTrace:  cellFunc,
		energized: make(map[Direction]bool),
		neighbors: make(map[Direction]*Cell),
	}
}

func (c *Cell) Enter(enterDir Direction) {
	if _, ok := c.energized[enterDir]; ok {
		return
	}
	c.energized[enterDir] = true
	exitDirs := c.rayTrace(enterDir)
	for exitDir, _ := range exitDirs {
		if n, ok := c.neighbors[exitDir]; ok {
			n.Enter(exitDir.Opposite())
		}
	}
}

func (c *Cell) isEnergized() bool {
	return len(c.energized) > 0
}
