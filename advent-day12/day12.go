package advent_day12

import "strings"

type Record string

type Match string

func singleMatch(record string, groupLength int) (remainder string, ok bool) {
	if len(record) < groupLength {
		ok = false
		return
	}
	if len(record) == groupLength {
		ok = !strings.Contains(record, ".")
		remainder = ""
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
