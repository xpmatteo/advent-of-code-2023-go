package advent_day15

import (
	"github.com/stretchr/testify/assert"
	"os"
	"strings"
	"testing"
)

func Test_encodingString(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{"one letter", "H", 200},
		{"one word", "HASH", 52},
		{"two words", "H,HASH", 252},
		{"ignores newlines", "H,\nHASH", 252},
		{"sample", "rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7", 1320},
		{"part1", readFile("input.txt"), 495972},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.want, hash(test.input))
		})
	}
}

func readFile(fileName string) string {
	bytes, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}

func hash(input string) int {
	sum := 0
	input = strings.ReplaceAll(input, "\n", "")
	for _, token := range strings.Split(input, ",") {
		sum += hash1(token)
	}
	return sum
}

func hash1(input string) int {
	sum := 0
	for _, r := range input {
		sum += int(r)
		sum *= 17
		sum %= 256
	}
	return sum
}
