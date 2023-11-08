package foursquare

import (
	"fmt"
	"sort"
	"strconv"
)

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

func CountSpecificTile(tiles []string, tile string) int {

	count := 0
	for _, t := range tiles {
		if t == tile {
			count++
		}
	}

	return count
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

func CountTargetTiles(tiles []string, targets []string) int {

	count := 0
	for _, t := range tiles {
		for _, target := range targets {
			if t == target {
				count++
			}
		}
	}

	return count
}

func MakeTiles(suit string, numbers []int) []string {

	var tiles []string
	for _, n := range numbers {
		tiles = append(tiles, fmt.Sprintf("%s%d", suit, n))
	}

	return tiles
}

func MakeStraight(tile string) []string {

	suit := tile[0:1]
	num := tile[1:2]

	n, _ := strconv.Atoi(num)

	tiles := make([]string, 3)
	for i, _ := range tiles {
		tiles[i] = fmt.Sprintf("%s%d", suit, n+i)
	}

	return tiles
}

func RemoveTiles(tiles []string, targets []string) ([]string, int) {

	removed := 0

	var newTiles []string
	newTiles = append(newTiles, tiles...)

	for _, target := range targets {

		for i, t := range newTiles {
			if t == target {
				newTiles = append(newTiles[0:i], newTiles[i+1:len(newTiles)]...)
				removed++
				break
			}
		}
	}

	return newTiles, removed
}

func FigureEyesCandidates(tiles []string, rules *ResolverRules) []string {

	if len(tiles) == 0 {
		return []string{}
	}

	if len(tiles)%3 != 2 {
		return []string{}
	}

	if !rules.Straight {
		return AggregateTiles(tiles)
	}

	// Using first tile to figure out suit
	suit := tiles[0][0:1]

	parts := [][]string{
		MakeTiles(suit, []int{1, 4, 7}),
		MakeTiles(suit, []int{2, 5, 8}),
		MakeTiles(suit, []int{3, 6, 9}),
	}

	parts = append(parts)

	countByPart := []int{
		CountTargetTiles(tiles, parts[0]) % 3,
		CountTargetTiles(tiles, parts[1]) % 3,
		CountTargetTiles(tiles, parts[2]) % 3,
	}

	// Find the different one
	statistics := make(map[int][]int)
	for i, p := range countByPart {

		s, ok := statistics[p]
		if !ok {
			s = make([]int, 0)
		}

		s = append(s, i)
		statistics[p] = s
	}

	// figure the range for eyes
	found := 0
	for _, s := range statistics {
		if len(s) == 1 {
			found = s[0]
		}
	}

	candidates := parts[found]

	return candidates
}

func FigureEyesWithCandidates(tiles []string, candidates []string, rules *ResolverRules) (string, []string) {

	for _, c := range candidates {

		if CountSpecificTile(tiles, c) >= 2 {

			// Attempt to take off eyes than check the entire tiles
			assume := []string{c, c}
			t, _ := RemoveTiles(tiles, assume)

			// the remaining cards are insufficient for a winning hand
			if !CheckWinningTiles(t, false, rules) {
				continue
			}

			// Found
			return c, t
		}
	}

	return "", tiles
}

func RemoveEyes(tiles []string, rules *ResolverRules) ([]string, string) {

	// Attempt to find eyes
	candidates := FigureEyesCandidates(tiles, rules)
	eyes, t := FigureEyesWithCandidates(tiles, candidates, rules)

	return t, eyes

}

func CheckWinningTiles(tiles []string, hasEyes bool, rules *ResolverRules) bool {

	if len(tiles) == 0 {
		return true
	}

	t := tiles

	if hasEyes {

		if len(t)%3 == 0 {
			return false
		}

		t, _ = RemoveEyes(tiles, rules)
	}

	// Not win
	if len(t)%3 != 0 {
		return false
	}

	var leftTiles []string
	leftTiles = append(leftTiles, t...)

	sort.Strings(leftTiles)

	for len(leftTiles) > 0 {

		// Is triplet
		if rules.Triplet {
			if CountSpecificTile(leftTiles, leftTiles[0]) >= 3 {
				// Attempt to remove triplet
				leftTiles = leftTiles[3:len(leftTiles)]
				continue
			} else if !rules.Straight {
				// Triplet only
				return false
			}
		}

		if rules.Straight {

			// Attempt to remove staight
			straight := MakeStraight(leftTiles[0])
			newTiles, n := RemoveTiles(leftTiles, straight)

			// Not win
			if n != 3 {
				return false
			}

			leftTiles = newTiles
		}
	}

	return true
}

/*
func MakeTileSegmentations(tiles []string) [][]string {

	// Check if tiles contain eyes
	if len(tiles)%3 != 0 {
		return false
	}
}
*/
