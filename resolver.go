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
	IsWin               bool     `json:"is_win"`
	IsReadyHand         bool     `json:"is_ready_hand"`
	ReadyHandCandidates []string `json:"ready_hand_candidates"`
	Eyes                []string `json:"eyes"`
}

var (
	SuitedTileRule = &ResolverRules{
		Triplet:  true,
		Straight: true,
	}
	HonorTileRule = &ResolverRules{
		Triplet:  true,
		Straight: false,
	}
)

func NewResolvedState() *ResolvedState {
	return &ResolvedState{
		IsWin:               false,
		IsReadyHand:         false,
		ReadyHandCandidates: make([]string, 0),
		Eyes:                make([]string, 0),
	}
}

type Resolver struct {
	tilesOpts *TilesOptions
}

func NewResolver(tilesOpts *TilesOptions) *Resolver {
	return &Resolver{
		tilesOpts: tilesOpts,
	}
}

func (r *Resolver) makeSuitGroups(tiles []string) map[TileSuit][]string {

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

func (r *Resolver) aggregateTiles(tiles []string) []string {

	grouped := make([]string, 0)
	tileMap := make(map[string]int)

	for _, t := range tiles {
		m, ok := tileMap[t]
		if !ok {
			m = 0
		}

		m++

		tileMap[t] = m
	}

	for t, _ := range tileMap {
		grouped = append(grouped, t)
	}

	return grouped
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

func (r *Resolver) figureEyesWithCandidates(tiles []string, candidates []string, rules *ResolverRules) (string, []string) {

	for _, c := range candidates {

		if r.count(tiles, c) >= 2 {

			// Attempt to take off eyes than check the entire tiles
			assume := []string{c, c}
			t, _ := r.filterTiles(tiles, assume)

			// the remaining cards are insufficient for a winning hand
			if !r.checkTiles(t, false, rules) {
				continue
			}

			// Found
			return c, t
		}
	}

	return "", tiles
}

func (r *Resolver) figureEyesCandidates(tiles []string, rules *ResolverRules) []string {

	if len(tiles) == 0 {
		return []string{}
	}

	if len(tiles)%3 != 2 {
		return []string{}
	}

	if !rules.Straight {
		return r.aggregateTiles(tiles)
	}

	// Using first tile to figure out suit
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

	return candidates
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

func (r *Resolver) filterEyes(tiles []string, rules *ResolverRules) ([]string, string) {

	// Attempt to find eyes
	candidates := r.figureEyesCandidates(tiles, rules)
	eyes, t := r.figureEyesWithCandidates(tiles, candidates, rules)

	return t, eyes

}

func (r *Resolver) checkTiles(tiles []string, hasEyes bool, rules *ResolverRules) bool {

	if len(tiles) == 0 {
		return true
	}

	t := tiles

	if hasEyes {

		if len(t)%3 == 0 {
			return false
		}

		t, _ = r.filterEyes(tiles, rules)
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

func (r *Resolver) figureReadyHandConditions(suit TileSuit, tiles []string, rules *ResolverRules) (bool, []string) {

	var candidates []string

	mod := len(tiles) % 3
	if mod == 0 {
		return false, []string{}
	}

	hasEyes := false
	if mod == 1 {
		hasEyes = true
	}

	var num int

	switch suit {
	case TileSuitWan:
		num = r.tilesOpts.Wan.Numbers
	case TileSuitTong:
		num = r.tilesOpts.Tong.Numbers
	case TileSuitBamboo:
		num = r.tilesOpts.Bamboo.Numbers
	case TileSuitWind:
		num = r.tilesOpts.Wind.Numbers
	case TileSuitDragon:
		num = r.tilesOpts.Dragon.Numbers
	}

	tries := GenTiles(suit, num, 1)

	for _, t := range tries {

		// Add tile then check it
		ts := append(tiles, t)

		if r.checkTiles(ts, hasEyes, rules) {
			candidates = append(candidates, t)
		}
	}

	if len(candidates) == 0 {
		return false, candidates
	}

	return true, candidates
}

func (r *Resolver) resolveSuitTiles(suit TileSuit, tiles []string, hasEyes bool) *ResolvedState {

	// Determine rules for suit
	var rules *ResolverRules

	switch suit {
	case TileSuitWan, TileSuitTong, TileSuitBamboo:
		rules = SuitedTileRule
	case TileSuitWind, TileSuitDragon:
		rules = HonorTileRule
	}

	state := NewResolvedState()

	if hasEyes {

		// Attempt to find eyes
		candidates := r.figureEyesCandidates(tiles, rules)
		eyes, _ := r.figureEyesWithCandidates(tiles, candidates, rules)
		state.Eyes = append(state.Eyes, eyes)
	}

	// Check if the combination of tiles meets the conditions for winning
	if !r.checkTiles(tiles, hasEyes, rules) {

		state.IsWin = false

		if hasEyes {
			// Check if tiles meets the conditions for ready hand
			isReadyHand, candidates := r.figureReadyHandConditions(suit, tiles, rules)
			state.IsReadyHand = isReadyHand
			state.ReadyHandCandidates = candidates
		}
	} else {
		state.IsWin = true
	}

	return state
}

func (r *Resolver) Resolve(tiles []string) *ResolvedState {

	groups := r.makeSuitGroups(tiles)

	var states []*ResolvedState
	for suit, g := range groups {

		// This group has no eyes
		if len(g)%3 == 0 {
			s := r.resolveSuitTiles(suit, g, false)
			states = append(states, s)
			continue
		}

		// Find and remove eyes before resolve tiles
		s := r.resolveSuitTiles(suit, g, true)
		states = append(states, s)
	}

	state := NewResolvedState()
	state.IsWin = true

	for _, s := range states {

		if !s.IsWin {
			state.IsWin = false
		}

		if s.IsReadyHand {
			state.IsReadyHand = true
			state.ReadyHandCandidates = append(state.ReadyHandCandidates, s.ReadyHandCandidates...)
		}

		state.Eyes = append(state.Eyes, s.Eyes...)
	}

	// Shouldn't contains two set of eyes for winning, it is ready hand condition
	if state.IsWin && len(state.Eyes) > 1 {
		state.IsWin = false
		state.IsReadyHand = true
		state.ReadyHandCandidates = state.Eyes
		state.Eyes = []string{}
	}

	return state
}
