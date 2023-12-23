package day10

type Direction struct {
	direct string
}

var (
	N = Direction{"NORTH"}
	E = Direction{"EAST"}
	S = Direction{"SOUTH"}
	W = Direction{"WEST"}
)

type Map struct {
}

func (m *Map) Load(input string) {

}

func (m *Map) Go(row int, col int, dir Direction) (int, int, Direction) {
	return 0, 0, N
}

func NewMap(input string) *Map {
	m := &Map{}
	m.Load(input)
	return m
}
