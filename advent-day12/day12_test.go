package advent_day12

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_simpleMatch(t *testing.T) {
	pattern := Pattern("# 1")

	matches := SimpleMatch(pattern)

	assert.Equal(t, []Match{"#"}, matches)
}
