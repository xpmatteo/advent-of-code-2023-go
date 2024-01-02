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

func numberOfInitialQuestionMarks(s string) int {
	count := 0
	for len(s) > 0 && s[0:1] == "?" {
		count++
	}
	return count
}

func singleMatch(record string, groupLength int) (remainder string, ok bool) {
	record = skipDotsPrefix(record)
	// partition the string in three parts
	// [questionmarks][hashes][remainder]
	// ???####...   can match between 4 and 7
	// ???####?..   can match between 4 and 8
	// ???####?#..  can match between 4 and 7 or 9
	// ???####??#.. can match between 4 and 8 or 10

	if len(record) < groupLength {
		ok = false
		return
	}
	if len(record) == groupLength {
		ok = !strings.Contains(record, ".")
		remainder = ""
		return
	}
	prefix := record[:groupLength]
	nextChar := record[groupLength : groupLength+1]
	if !strings.Contains(prefix, ".") && nextChar != "#" {
		ok = true
		remainder = record[groupLength+1:]
		return
	}
	if record[0:1] == "?" {
		return singleMatch(record[1:], groupLength)
	}
	return "", false
}
