package advent_day15

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"slices"
	"strconv"
	"strings"
	"testing"
)

const sampleCommands = "rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7"

func Test_encodingString(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{"one letter", "H", 200},
		{"one word", "HASH", 52},
		{"two words", "H,HASH", 252},
		{"sample", sampleCommands, 1320},
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

func Test_applyCommands(t *testing.T) {
	tests := []struct {
		commands string
		want     boxes
	}{
		{"rn-", boxes{
			nil, nil, nil, nil,
		}},
		{"rn=3", boxes{
			{{"rn", 3}},
			nil,
			nil,
			nil,
		}},
		{"rn=3,cm=4", boxes{
			{{"rn", 3}, {"cm", 4}},
			nil,
			nil,
			nil,
		}},
		{"rn=3,rn=4", boxes{
			{{"rn", 4}},
			nil,
			nil,
			nil,
		}},
		{"rn=3,rn-", boxes{
			{},
			nil,
			nil,
			nil,
		}},
		{"rn=1,cm-,qp=3", boxes{
			{{"rn", 1}},
			{{"qp", 3}},
			nil,
			nil,
		}},
		{"rn=1,cm-,qp=3,cm=2", boxes{
			{{"rn", 1}, {"cm", 2}},
			{{"qp", 3}},
			nil,
			nil,
		}},
		{"rn=1,cm-,qp=3,cm=2,qp-,pc=4", boxes{
			{{"rn", 1}, {"cm", 2}},
			{},
			nil,
			{{"pc", 4}},
		}},
		{sampleCommands, boxes{
			{{"rn", 1}, {"cm", 2}},
			{},
			nil,
			{{"ot", 7}, {"ab", 5}, {"pc", 6}},
		}},
	}
	for _, test := range tests {
		t.Run(test.commands, func(t *testing.T) {
			assert.Equal(t, test.want, execute(test.commands, 4))
		})
	}
}

func Test_focusingPower(t *testing.T) {
	tests := []struct {
		name  string
		input boxes
		want  int
	}{
		{
			name: "one box one lens",
			input: boxes{
				{{"rn", 4}},
				nil,
				nil,
				nil,
			},
			want: 4,
		},
		{
			name: "one box two lens",
			input: boxes{
				{{"", 4}, {"", 5}},
				nil,
			},
			want: 4 + 2*5,
		},
		{
			name: "two boxes",
			input: boxes{
				{{"", 4}, {"", 5}},
				{{"", 7}},
			},
			want: 4 + 2*5 + 2*7,
		},
		{
			name:  "sample",
			input: execute(sampleCommands, 4),
			want:  145,
		},
		{
			name:  "input.txt",
			input: execute(readFile("input.txt"), 256),
			want:  245223,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.want, focusingPower(test.input))
		})
	}
}

func focusingPower(input boxes) int {
	sum := 0
	for iBox, box := range input {
		for iLens, lens := range box {
			sum += (iBox + 1) * (iLens + 1) * lens.focalLength
		}
	}
	return sum
}

func execute(commands string, size int) boxes {
	result := make([]box, size)
	for _, token := range strings.Split(commands, ",") {
		cmd := parseCommand(token)
		box := result[cmd.index]
		boxIndex := slices.IndexFunc(box, func(elt Lens) bool {
			return elt.label == cmd.label
		})
		if cmd.kind == add {
			if boxIndex >= 0 {
				box[boxIndex].focalLength = cmd.focalLength
			} else {
				result[cmd.index] = append(result[cmd.index], Lens{cmd.label, cmd.focalLength})
			}
		} else {
			if boxIndex >= 0 {
				result[cmd.index] = append(box[:boxIndex], box[boxIndex+1:]...)
			}
		}
	}
	return result
}

func parseCommand(input string) *Command {
	if input[len(input)-1] == '-' {
		label := strings.Split(input, "-")[0]
		return &Command{
			label:       label,
			index:       hash1(label),
			kind:        remove,
			focalLength: 0,
		}
	} else if strings.Contains(input, "=") {
		split := strings.Split(input, "=")
		label, focalString := split[0], split[1]
		focal, err := strconv.Atoi(focalString)
		if err != nil {
			panic(err)
		}
		return &Command{
			label:       label,
			index:       hash1(label),
			kind:        add,
			focalLength: focal,
		}
	} else {
		panic("bad input commands: " + input)
	}
}

type CommandKind int

var add CommandKind = 1
var remove CommandKind = 2

type Lens struct {
	label       string
	focalLength int
}
type box []Lens
type boxes []box

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
	label       string
	index       int
	kind        CommandKind
	focalLength int
}

func (c *Command) String() string {
	return fmt.Sprintf("%s %d %s %d", c.label, c.index, c.kind, c.focalLength)
}

func readFile(fileName string) string {
	bytes, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	return strings.ReplaceAll(string(bytes), "\n", "")
}

func hash(input string) int {
	sum := 0
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
