package advent_day15

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"strconv"
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
		{"rn", "rn", 0},
		{"cm", "cm", 0},
		{"qp", "qp", 1},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.want, hash(test.input))
		})
	}
}

func Test_parseCommands(t *testing.T) {
	tests := []struct {
		input string
		want  *Command
	}{
		{"rn-", &Command{"rn", 0, remove, 0}},
		{"qp-", &Command{"qp", 1, remove, 0}},
		{"qp=3", &Command{"qp", 1, add, 3}},
	}
	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			assert.Equal(t, test.want, parseCommand(test.input))
		})
	}
}

func parseCommand(input string) *Command {
	if input[len(input)-1] == '-' {
		label := strings.Split(input, "-")[0]
		return &Command{
			label: label,
			index: hash1(label),
			kind:  remove,
			focal: 0,
		}
	} else if strings.Contains(input, "=") {
		split := strings.Split(input, "=")
		label, focalString := split[0], split[1]
		focal, err := strconv.Atoi(focalString)
		if err != nil {
			panic(err)
		}
		return &Command{
			label: label,
			index: hash1(label),
			kind:  add,
			focal: focal,
		}
	} else {
		panic("bad input command: " + input)
	}
}

type CommandKind int

var add CommandKind = 1
var remove CommandKind = 2

func (k CommandKind) String() string {
	if k == add {
		return "add"
	} else if k == remove {
		return "remove"
	} else {
		panic(fmt.Errorf("invalid CommandKind %d", k))
	}
}

type Command struct {
	label string
	index int
	kind  CommandKind
	focal int
}

func (c *Command) String() string {
	return fmt.Sprintf("%s %d %s %d", c.label, c.index, c.kind, c.focal)
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
