package advent_day12

import (
	"strconv"
	"strings"
)

type Record string

type Match string

func SimpleMatch(record Record) []Match {
	tokens := strings.Split(string(record), " ")
	//pattern := tokens[0]
	count, err := strconv.Atoi(tokens[1])
	if err != nil {
		panic(err)
	}

	return []Match{Match(strings.Repeat("#", count))}
}
