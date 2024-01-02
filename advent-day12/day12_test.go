package advent_day12

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"os"
	"reflect"
	"testing"
)

func Test_multiGroup(t *testing.T) {
	tests := []struct {
		record   string
		groups   []int
		expected int
	}{
		{"??", []int{1}, 2},
		{"???.###", []int{1, 1, 3}, 1},
		{"..??...?##.", []int{1, 3}, 2},
		{".??..??...?##.", []int{1, 1, 3}, 4},
		{"#?#?#?", []int{6}, 1},
		{"?#?#?#?", []int{6}, 2},
		{"#?#?#?#?", []int{1, 6}, 1},
		{"?#?#?#?#?", []int{1, 6}, 1},
		{"?#?#?#?#?#?#?", []int{3, 1, 6}, 1},
		{"?#?#?#?#?#?#?#?", []int{1, 3, 1, 6}, 1},
		{"????.#...#...", []int{4, 1, 1}, 1},
		{"????.######..#####.", []int{1, 6, 5}, 4},
		{"?###????????", []int{3, 2, 1}, 10},
		{".???????", []int{2, 1}, 10},
	}
	for _, test := range tests {
		t.Run(test.record, func(t *testing.T) {
			variants := generateVariants(test.record)
			count := 0
			for _, variant := range variants {
				if reflect.DeepEqual(variantToGroups(variant), test.groups) {
					count++
				}
			}
			actual := count

			assert.Equal(t, test.expected, actual)
		})
	}
}

func Test_parseLine(t *testing.T) {
	assert := assert.New(t)

	record, groups := parse("#?. 1,2,3")

	assert.Equal("#?.", record)
	assert.Equal([]int{1, 2, 3}, groups)
}

func Test_parseLine_oneGroup(t *testing.T) {
	assert := assert.New(t)

	record, groups := parse("#?. 2")

	assert.Equal("#?.", record)
	assert.Equal([]int{2}, groups)
}

const sample = `???.### 1,1,3
.??..??...?##. 1,1,3
?#?#?#?#?#?#?#? 1,3,1,6
????.#...#... 4,1,1
????.######..#####. 1,6,5
?###???????? 3,2,1`

func Test_samplePart_I(t *testing.T) {
	actual := part1(sample)

	assert.Equal(t, 21, actual)
}

func Test_acceptancePart_I(t *testing.T) {
	t.Skip("too slow")
	bytes, err := os.ReadFile("day12.txt")
	require.NoError(t, err)

	actual := part1(string(bytes))

	assert.Equal(t, 7674, actual)
}

func Test_generateVariants(t *testing.T) {
	tests := []struct {
		record   string
		expected []string
	}{
		{"#", []string{"#"}},
		{".", []string{"."}},
		{"?", []string{".", "#"}},
		{".??.", []string{"....", "..#.", ".#..", ".##."}},
	}
	for _, test := range tests {
		t.Run(test.record, func(t *testing.T) {
			assert := assert.New(t)

			variants := generateVariants(test.record)

			assert.Equal(test.expected, variants)
		})
	}
}

func Test_variantToGroups(t *testing.T) {
	assert.Equal(t, []int{1, 2, 3}, variantToGroups("...#..##.###"))
}
