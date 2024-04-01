package advent_day13

import "strings"

type Pattern []Line
type Line string
type Column string

func (c Column) score() int {
	return Line(c).score()
}

func NewPattern(p string) Pattern {
	split := strings.Split(p, "\n")
	pattern := make(Pattern, len(split))
	for i, s := range split {
		pattern[i] = Line(s)
	}
	return pattern
}

func (pattern Pattern) lines() []Line {
	return pattern
}

func (pattern Pattern) columns() []Column {
	columns := make([]Column, len(pattern[0]))
	for i := 0; i < len(pattern[0]); i++ {
		column := ""
		for _, line := range pattern {
			column += string(line[i])
		}
		columns[i] = Column(column)
	}
	return columns
}

func (line Line) isPalyndromic() bool {
	for i := 0; i < len(line)/2; i++ {
		if line[i] != line[len(line)-1-i] {
			return false
		}
	}
	return true
}

func score(pattern Pattern) int {
	v := verticalScore(pattern)
	h := horizontalScore(pattern)
	if v == 0 && h == 0 {
		panic("both hor and vert are 0")
	}
	return v + h*100
}

func verticalScore(pattern Pattern) int {
	candidateScore := pattern.lines()[0].score()
	for _, line := range pattern.lines() {
		if line.score() != candidateScore {
			return 0
		}
	}
	return candidateScore
}

func horizontalScore(pattern Pattern) int {
	columns := pattern.columns()
	candidateScore := columns[0].score()
	for _, column := range columns {
		if column.score() != candidateScore {
			return 0
		}
	}
	return candidateScore
}

func (line Line) score() int {
	if isEven(len(line)) && line.isPalyndromic() {
		return len(line) / 2
	}
	if isEven(len(line)) {
		return 0
	}
	if line[:len(line)-1].isPalyndromic() {
		return len(line) / 2
	}
	if line[1:].isPalyndromic() {
		return len(line)/2 + 1
	}
	return 0
}

func isEven(n int) bool {
	return n%2 == 0
}
