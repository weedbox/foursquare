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

		if CheckWinningTiles(ts, hasEyes, rules) {
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
		candidates := FigureEyesCandidates(tiles, rules)
		eyes, _ := FigureEyesWithCandidates(tiles, candidates, rules)
		state.Eyes = append(state.Eyes, eyes)
	}

	// Check if the combination of tiles meets the conditions for winning
	if !CheckWinningTiles(tiles, hasEyes, rules) {

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

	groups := MakeSuitGroups(tiles)

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
