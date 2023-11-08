package foursquare

func MakeSuitGroups(tiles []string) map[TileSuit][]string {

	groups := make(map[TileSuit][]string)

	for _, t := range tiles {
		suit := TileSuit(t[0:1])
		g, ok := groups[suit]
		if !ok {
			g = make([]string, 0)
		}

		g = append(g, t)
		groups[suit] = g
	}

	return groups
}

func AggregateTiles(tiles []string) []string {

	tileMap := CountByTiles(tiles)

	grouped := make([]string, 0)
	for t, _ := range tileMap {
		grouped = append(grouped, t)
	}

	return grouped
}

func CountByTiles(tiles []string) map[string]int {

	result := make(map[string]int)

	for _, t := range tiles {
		m, ok := result[t]
		if !ok {
			m = 0
		}

		m++

		result[t] = m
	}

	return result
}

func CountBySuits(tiles []string) map[TileSuit]int {

	result := make(map[TileSuit]int)

	for _, t := range tiles {
		suit := TileSuit(t[0:1])
		m, ok := result[suit]
		if !ok {
			m = 0
		}

		m++

		result[suit] = m
	}

	return result
}
