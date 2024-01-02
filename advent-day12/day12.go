package advent_day12

import (
	"strconv"
	"strings"
)

type Record string

type Match string

// return the input string without any initial "."
func skipDotsPrefix(s string) string {
	for len(s) > 0 && s[0:1] == "." {
		s = s[1:]
	}
	return s
}

func singleMatch(record string, groupLength int) (remainder string, ok bool) {
	record = skipDotsPrefix(record)

	if len(record) < groupLength {
		return "", false
	}
	if len(record) == groupLength {
		return "", !strings.Contains(record, ".")
	}
	prefix := record[:groupLength]
	nextChar := record[groupLength : groupLength+1]
	if !strings.Contains(prefix, ".") && nextChar != "#" {
		return "." + record[groupLength+1:], true
	}
	if record[0:1] == "?" {
		return singleMatch(record[1:], groupLength)
	}
	return "", false
}

func waysToMatchASingleGroup(record string, groupLength int) []string {
	result := []string{}
	for i := 0; i < len(record); i++ {
		if record[i:i+1] == "." {
			continue
		}
		remainder, ok := singleMatch(record[i:], groupLength)
		if !ok {
			break
		}
		result = append(result, remainder)
		if record[i:i+1] == "#" {
			break
		}
	}
	return result
}

func multiGroup(record string, groups []int) int {
	if len(groups) == 0 {
		return 1
	}
	result := 0
	ways := waysToMatchASingleGroup(record, groups[0])

	// remove adjacent duplicates from ways
	for i := 1; i < len(ways); i++ {
		if ways[i] == ways[i-1] {
			ways = append(ways[:i], ways[i+1:]...)
			i--
		}
	}
	for _, way := range ways {
		result += multiGroup(way, groups[1:])
	}
	return result
}

func parse(line string) (string, []int) {
	tokens := strings.Split(line, " ")
	record := tokens[0]

	numbers := strings.Split(tokens[1], ",")
	groups := []int{}
	for _, s := range numbers {
		num, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		groups = append(groups, num)
	}
	return record, groups
}

func part1(input string) int {
	result := 0
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		if len(line) < 1 {
			continue
		}
		record, groups := parse(line)
		result += multiGroup(record, groups)
	}
	return result
}
