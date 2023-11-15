package foursquare

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

type DiscardCandidate struct {
	DiscardedTile string   `json:"discarded_tile"`
	TargetTiles   []string `json:"target_tiles"`
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

func (g *Game) figureReadyHandConditions(suit TileSuit, tiles []string, rules *ResolverRules) (bool, []string) {

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
		num = g.gs.Meta.TileSetDef.Wan.Numbers
	case TileSuitTong:
		num = g.gs.Meta.TileSetDef.Tong.Numbers
	case TileSuitBamboo:
		num = g.gs.Meta.TileSetDef.Bamboo.Numbers
	case TileSuitWind:
		num = g.gs.Meta.TileSetDef.Wind.Numbers
	case TileSuitDragon:
		num = g.gs.Meta.TileSetDef.Dragon.Numbers
	}

	tries := GenTiles(suit, num, 1)

	for _, t := range tries {

		// Add tile then check it
		ts := append(tiles, t)

		if CheckWinningTiles(ts, hasEyes, rules) {
			candidates = append(candidates, t)
		}
	}

	if len(candidates) == 0 {
		return false, candidates
	}

	return true, candidates
}

func (g *Game) resolveSuitTiles(suit TileSuit, tiles []string, hasEyes bool) *ResolvedState {

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
		candidates := FigureEyesCandidates(tiles, rules)
		eyes, _ := FigureEyesWithCandidates(tiles, candidates, rules)
		state.Eyes = append(state.Eyes, eyes)
	}

	// Check if the combination of tiles meets the conditions for winning
	if !CheckWinningTiles(tiles, hasEyes, rules) {

		state.IsWin = false

		if hasEyes {
			// Check if tiles meets the conditions for ready hand
			isReadyHand, candidates := g.figureReadyHandConditions(suit, tiles, rules)
			state.IsReadyHand = isReadyHand
			state.ReadyHandCandidates = candidates
		}
	} else {
		state.IsWin = true
	}

	return state
}

func (g *Game) Resolve(tiles []string) *ResolvedState {

	groups := MakeSuitGroups(tiles)

	var states []*ResolvedState
	for suit, gp := range groups {

		// This group has no eyes
		if len(gp)%3 == 0 {
			s := g.resolveSuitTiles(suit, gp, false)
			states = append(states, s)
			continue
		}

		// Find and remove eyes before resolve tiles
		s := g.resolveSuitTiles(suit, gp, true)
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

func (g *Game) FigureDiscardCandidatesForReadyHand(tiles []string) []*DiscardCandidate {

	candidates := make([]*DiscardCandidate, 0)

	for i, t := range tiles {

		newTiles := append(tiles[0:i], tiles[i+1:len(tiles)]...)

		// Check if it is ready hand condition
		state := g.Resolve(newTiles)
		if !state.IsReadyHand {
			continue
		}

		c := &DiscardCandidate{
			DiscardedTile: t,
			TargetTiles:   state.ReadyHandCandidates,
		}

		candidates = append(candidates, c)
	}

	return candidates
}
