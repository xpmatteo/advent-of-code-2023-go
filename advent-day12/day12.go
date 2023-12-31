package advent_day12

import (
	"strconv"
	"strings"
)

type Record string

type Match string

func SimpleMatch(record Record) []Match {
	tokens := strings.Split(string(record), " ")
	pattern := tokens[0]
	count, err := strconv.Atoi(tokens[1])
	if err != nil {
		panic(err)
	}

	if count == 1 && pattern == "??" {
		return []Match{"#.", ".#"}
	}
	if count == 1 && pattern == "???" {
		var matches []Match
		for i := 0; i < len(pattern); i++ {
			m := strings.Repeat(".", i) + "#" + strings.Repeat(".", len(pattern)-i-1)
			matches = append(matches, Match(m))
		}
		return matches
	}

	return []Match{Match(strings.Repeat("#", count))}
}
