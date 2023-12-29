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
