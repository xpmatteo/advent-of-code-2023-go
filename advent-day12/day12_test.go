package advent_day12

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	tests := []struct {
		pattern           string
		length            int
		expectedRemainder string
		expectedOk        bool
	}{
		{"", 1, "", false},
		{"#", 1, "", true},
		//{"#.", 1, "", true},
		//{"#..", 1, ".", true},
		//{"#.?", 1, "?", true},
		//{"#.#", 1, "#", true},
		//{"#...", 1, "..", true},
		//{"#.##", 1, "##", true},
		//{"##", 1, "", false},
		//{"?", 1, "", true},
		//{"?.", 1, "", true},
	}
	for _, test := range tests {
		t.Run(string(test.pattern), func(t *testing.T) {
			assert := assert.New(t)

			remainder, ok := SimpleMatch(test.pattern, test.length)

			assert.Equal(test.expectedOk, ok)
			assert.Equal(test.expectedRemainder, remainder)
		})
	}
}
