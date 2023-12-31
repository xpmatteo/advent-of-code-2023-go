package advent_day12

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_simpleMatch(t *testing.T) {
	pattern := Record("# 1")

	matches := SimpleMatch(pattern)

	assert.Equal(t, []Match{"#"}, matches)
}

func Test(t *testing.T) {
	tests := []struct {
		pattern  Record
		expected []Match
	}{
		{"# 1", []Match{"#"}},
		{"## 2", []Match{"##"}},
		{"### 3", []Match{"###"}},
		{"? 1", []Match{"#"}},
	}
	for _, test := range tests {
		t.Run(string(test.pattern), func(t *testing.T) {
			assert.Equal(t, test.expected, SimpleMatch(test.pattern))
		})
	}
}
