package foursquare

import (
	"fmt"
	"sort"
	"strconv"
)

type ResolverRules struct {
	Triplet  bool
	Straight bool
}

type ResolvedState struct {
	IsWin       bool     `json:"is_win"`
	IsReadyHand bool     `json:"is_ready_hand"`
	Eyes        []string `json:"eyes"`
}

func NewResolvedState() *ResolvedState {
	return &ResolvedState{
		IsWin:       false,
		IsReadyHand: false,
		Eyes:        make([]string, 0),
	}
}

type Resolver struct {
}

func NewResolver() *Resolver {
	return &Resolver{}
}

func (r *Resolver) makeGroups(tiles []string) map[string][]string {

	groups := make(map[string][]string)

	for _, t := range tiles {
		suit := t[0:1]
		g, ok := groups[suit]
		if !ok {
			g = make([]string, 0)
		}

		g = append(g, t)
		groups[suit] = g
	}

	return groups
}

func (r *Resolver) count(tiles []string, tile string) int {

	count := 0
	for _, t := range tiles {
		if t == tile {
			count++
		}
	}

	return count
}

func (r *Resolver) countByTiles(tiles []string, targets []string) int {

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

func (r *Resolver) genTiles(suit string, numbers []int) []string {

	var tiles []string
	for _, n := range numbers {
		tiles = append(tiles, fmt.Sprintf("%s%d", suit, n))
	}

	return tiles
}

func (r *Resolver) figureEyesOfGroupWithCandidates(tiles []string, candidates []string, rules *ResolverRules) string {

	for _, c := range candidates {

		if r.count(tiles, c) >= 2 {

			// Attempt to take off eyes than check the entire tiles
			assume := []string{c, c}
			t, _ := r.filterTiles(tiles, assume)

			// the remaining cards are insufficient for a winning hand
			if !r.resolveGroup(t, rules) {
				continue
			}

			return c
		}
	}

	return ""
}

func (r *Resolver) figureEyesOfGroup(tiles []string, rules *ResolverRules) string {

	if len(tiles) == 0 {
		return ""
	}

	if len(tiles)%3 != 2 {
		return ""
	}

	if !rules.Straight {
		return r.figureEyesOfGroupWithCandidates(tiles, tiles, rules)
	}

	suit := tiles[0][0:1]

	parts := [][]string{
		r.genTiles(suit, []int{1, 4, 7}),
		r.genTiles(suit, []int{2, 5, 8}),
		r.genTiles(suit, []int{3, 6, 9}),
	}

	parts = append(parts)

	countByPart := []int{
		r.countByTiles(tiles, parts[0]) % 3,
		r.countByTiles(tiles, parts[1]) % 3,
		r.countByTiles(tiles, parts[2]) % 3,
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

	return r.figureEyesOfGroupWithCandidates(tiles, candidates, rules)
}

func (r *Resolver) isTiplet(tiles []string) bool {

	for _, t := range tiles {
		if t != tiles[0] {
			return false
		}
	}

	return true
}

func (r *Resolver) makeStraight(tile string) []string {

	suit := tile[0:1]
	num := tile[1:2]

	n, _ := strconv.Atoi(num)

	tiles := make([]string, 3)
	for i, _ := range tiles {
		tiles[i] = fmt.Sprintf("%s%d", suit, n+i)
	}

	return tiles
}

func (r *Resolver) filterTiles(tiles []string, targets []string) ([]string, int) {

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

func (r *Resolver) resolveGroup(tiles []string, rules *ResolverRules) bool {

	// Not win
	if len(tiles)%3 != 0 {
		return false
	}

	var leftTiles []string
	leftTiles = append(leftTiles, tiles...)

	sort.Strings(leftTiles)

	for len(leftTiles) > 0 {

		// Is triplet
		if rules.Triplet {
			if r.count(leftTiles, leftTiles[0]) >= 3 {
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
			straight := r.makeStraight(leftTiles[0])
			newTiles, n := r.filterTiles(leftTiles, straight)

			// Not win
			if n != 3 {
				return false
			}

			leftTiles = newTiles
		}
	}

	return true
}

func (r *Resolver) Resolve(tiles []string) *ResolvedState {

	state := NewResolvedState()

	groups := r.makeGroups(tiles)

	for s, g := range groups {

		switch s {
		case TileSuitWan:
			fallthrough
		case TileSuitTong:
			fallthrough
		case TileSuitBamboo:
			isWin := r.resolveGroup(g, &ResolverRules{
				Triplet:  true,
				Straight: true,
			})

			if !isWin {
				state.IsWin = false
			}
		case TileSuitWind:
			fallthrough
		case TileSuitDragon:
			isWin := r.resolveGroup(g, &ResolverRules{
				Triplet:  true,
				Straight: false,
			})

			if !isWin {
				state.IsWin = false
			}
		}
	}

	return state
}
