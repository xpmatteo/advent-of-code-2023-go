package advent_day13

import "strings"

type Line string
type Pattern []Line

func NewPattern(p string) Pattern {
	split := strings.Split(p, "\n")
	pattern := make(Pattern, len(split))
	for i, s := range split {
		pattern[i] = Line(s)
	}
	return pattern
}

func (line Line) isPalyndromic() bool {
	for i := 0; i < len(line)/2; i++ {
		if line[i] != line[len(line)-1-i] {
			return false
		}
	}
	return true
}
