package ringsandrods

func countPoints(rings string) int {
	m := make(map[rune]int)
	m['R'] = 0b001
	m['G'] = 0b010
	m['B'] = 0b100

	rods := [10]int{}
	for i := 0; i < len(rings); i += 2 {
		color := rings[i]
		rod := rings[i+1] - '0'
		rods[rod] = rods[rod] | m[rune(color)]
	}

	count := 0
	for _, v := range rods {
		if v == 0b111 {
			count++
		}
	}

	return count
}
