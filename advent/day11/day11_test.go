package day11

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

const sampleData = `...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....`

func TestStarField_New(t *testing.T) {
	sf := NewStarField(".#.\n# #")

	expected := []Coordinate{{0, 1}, {1, 0}, {1, 2}}
	assert.Equal(t, expected, sf.stars)
}

func TestStarField_maxCoords(t *testing.T) {
	assert := assert.New(t)
	sf := NewStarField(`.#.
...
#..`)

	assert.Equal(Row(2), sf.maxRow)
	assert.Equal(Col(1), sf.maxCol)
}

func Test_IsEmptyRow(t *testing.T) {
	assert := assert.New(t)
	sf := NewStarField(sampleData)

	assert.False(sf.isEmptyRow(2), "row 2 is not empty")
	assert.True(sf.isEmptyRow(3), "row 3 is empty")
}

func TestStarField_ExpandEmptyRows(t *testing.T) {
	sf := NewStarField(`.
#
.
#`)

	sf.expandEmptyRows(1)

	expected := []Coordinate{{2, 0}, {5, 0}}
	assert.Equal(t, expected, sf.stars)
}

func TestStarField_ExpandEmptyCols(t *testing.T) {
	sf := NewStarField(`#.##.#`)

	sf.expandEmptyCols(1)

	expected := []Coordinate{{0, 0}, {0, 3}, {0, 4}, {0, 7}}
	assert.Equal(t, expected, sf.stars)
}

func TestStarField_Expand(t *testing.T) {
	sf := NewStarField(`.#.
...
#..#`)

	sf.Expand(1)

	expected := []Coordinate{{0, 1}, {3, 0}, {3, 4}}
	assert.Equal(t, expected, sf.stars, sf.String())
}

func Test_StarDistance(t *testing.T) {
	assert := assert.New(t)

	s0, s1 := Coordinate{0, 0}, Coordinate{3, 4}

	assert.Equal(7, s0.Distance(s1))
}

func Test_SamplePart1(t *testing.T) {
	assert := assert.New(t)
	sf := NewStarField(sampleData)
	sf.Expand(1)

	assert.Equal(374, sf.SumDistances())
}

func Test_Acceptance_Part1(t *testing.T) {
	assert := assert.New(t)
	bytes, err := os.ReadFile("day11.txt")
	require.NoError(t, err)
	sf := NewStarField(string(bytes))
	sf.Expand(1)

	assert.Equal(9521550, sf.SumDistances())
}
