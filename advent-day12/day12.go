package advent_day12

type Record string

type Match string

func SimpleMatch(record string, length int) (remainder string, ok bool) {
	if len(record) >= length {
		remainder, ok = record[length:], true
	}
	return
}
