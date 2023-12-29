package day11

import (
	//"fmt"
	"github.com/stretchr/testify/assert"
	//"github.com/stretchr/testify/require"
	//"testing"
	"testing"
)

func TestStarField_New(t *testing.T) {
	sf := NewStarField(".#.\n# #")

	expected := []Coordinate{{0, 1}, {1, 0}, {1, 2}}
	assert.Equal(t, expected, sf.stars)
}

func TestStarField_enumCoords(t *testing.T) {
	assert := assert.New(t)
	sf := NewStarField(`.#.
...
#..`)

	assert.Equal(2, sf.maxRow)
	assert.Equal(1, sf.maxCol)
}

func TestStarField_ExpandEmptyRows(t *testing.T) {
	t.Skip("for now")
	sf := NewStarField(`.#.
...
#..
...`)

	sf.ExpandEmptyRows()

	expected := []Coordinate{{0, 1}, {3, 0}}
	assert.Equal(t, expected, sf.stars)
}
