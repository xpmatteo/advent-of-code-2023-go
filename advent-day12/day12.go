package advent_day12

import (
	"strconv"
	"strings"
)

func singleMatch(record string, groupLength int) (remainder string, ok bool) {
	if len(record) < groupLength {
		return "", false
	}
	if record[0:1] == "." {
		return singleMatch(record[1:], groupLength)
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
	if len(record) == 0 {
		return []string{}
	}
	if record[0] == '.' {
		return waysToMatchASingleGroup(record[1:], groupLength)
	}
	if record[0] == '#' {
		remainder, ok := singleMatch(record, groupLength)
		if !ok {
			return []string{}
		}
		return []string{remainder}
	}
	if record[0] == '?' {
		remainder, ok := singleMatch(record, groupLength)
		result := []string{}
		if ok {
			result = append(result, remainder)
		}
		return append(result, waysToMatchASingleGroup(record[1:], groupLength)...)
	}
	panic("unknown first char: " + record)
}

func countMatches(record string, groups []int) int {
	return countMatchesAux(record, groups, 0)
}

func countMatchesAux(record string, groups []int, recursionLevel int) int {
	if len(groups) == 0 {
		if strings.Contains(record, "#") {
			// we did not consume all non-optional matches
			return 0
		}
		// we found a correct match!
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
		//if recursionLevel < 3 {
		//	fmt.Printf("countMatches: level %4d group %v way %s\n", recursionLevel, groups, way)
		//}
		result += countMatchesAux(way, groups[1:], recursionLevel+1)
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
		result += countMatches(record, groups)
	}
	return result
}

func part2(input string) int {
	result := 0
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		if len(line) < 1 {
			continue
		}
		record, groups := parse(unfold(line))
		result += countMatches(record, groups)
		//fmt.Printf("Done line %d\r", i)
	}
	return result
}

func repeat(s string, count int) []string {
	var result []string
	for i := 0; i < 5; i++ {
		result = append(result, s)
	}
	return result
}

func unfold(line string) string {
	tokens := strings.Split(line, " ")
	record5 := repeat(tokens[0], 5)
	unfoldedRecord := strings.Join(record5, "?")
	groups5 := repeat(tokens[1], 5)
	unfoldedGroups := strings.Join(groups5, ",")
	return unfoldedRecord + " " + unfoldedGroups
}
