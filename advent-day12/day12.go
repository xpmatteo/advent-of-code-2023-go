package advent_day12

type Record string

type Match string

func singleMatch(record string, groupLength int) (remainder string, ok bool) {
	if len(record) < groupLength {
		ok = false
		return
	}
	if len(record) == groupLength {
		remainder, ok = "", true
		return
	}
	remainder, ok = record[groupLength:], true
	if remainder[0:1] == "#" {
		ok = false
	}
	if remainder[0:1] == "?" {
		remainder = "." + remainder[1:]
	}
	return
}
