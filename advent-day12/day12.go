package advent_day12

import "strings"

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
		return record[groupLength+1:], true
	}
	if record[0:1] == "?" {
		return singleMatch(record[1:], groupLength)
	}
	return "", false
}

func multiMatch(record string, groups []int) int {
	if len(groups) == 0 {
		return 0
	}
	remainder, ok := singleMatch(record, groups[0])
	if !ok {
		return 0
	}
	if ok && len(groups) == 1 {
		return 1
	}
	return multiMatch(remainder, groups[1:])
}
