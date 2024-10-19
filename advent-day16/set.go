package advent_day16

type DirSet map[Direction]struct{}

func setOf(d ...Direction) DirSet {
	result := make(DirSet)
	for _, direction := range d {
		result[direction] = struct{}{}
	}
	return result
}

func (set DirSet) String() string {
	result := ""
	for dir := range set {
		result += " " + string(dir)
	}
	return result[1:]
}
