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
		ok = false
		return
	}
	if len(record) == groupLength {
		ok = !strings.Contains(record, ".")
		remainder = ""
		return
	}
	if record[0:2] == "?#" {
		remainder = record[2:]
		ok = true
		return
	}

	remainder = record[groupLength:]
	ok = !strings.Contains(record[:groupLength], ".")

	if remainder[0:1] == "#" {
		ok = false
	}
	if remainder[0:1] == "?" {
		remainder = "." + remainder[1:]
	}
	return
}
