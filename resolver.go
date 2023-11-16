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

func ResolveSuitTiles(tileSetDef *TileSetDef, suit TileSuit, tiles []string, hasEyes bool) *ResolvedState {

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
			isReadyHand, candidates := FigureReadyHandConditions(tileSetDef, suit, tiles, rules)
			state.IsReadyHand = isReadyHand
			state.ReadyHandCandidates = candidates
		}
	} else {
		state.IsWin = true
	}

	return state
}

func Resolve(tileSetDef *TileSetDef, tiles []string) *ResolvedState {

	groups := MakeSuitGroups(tiles)

	var states []*ResolvedState
	for suit, gp := range groups {

		// This group has no eyes
		if len(gp)%3 == 0 {
			s := ResolveSuitTiles(tileSetDef, suit, gp, false)
			states = append(states, s)
			continue
		}

		// Find and remove eyes before resolve tiles
		s := ResolveSuitTiles(tileSetDef, suit, gp, true)
		states = append(states, s)
	}

	state := NewResolvedState()
	state.IsWin = true

	notWinCount := 0
	for _, s := range states {
		if !s.IsWin {
			notWinCount++
		}
	}

	// impossible to win
	if notWinCount > 1 {
		state.IsWin = false
		state.IsReadyHand = false
		state.Eyes = []string{}
		return state
	}

	for _, s := range states {

		//fmt.Println(s.IsWin, s.IsReadyHand, s.ReadyHandCandidates, s.Eyes)

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

func FigureDiscardCandidatesForReadyHand(tileSetDef *TileSetDef, tiles []string) []*DiscardCandidate {

	candidates := make([]*DiscardCandidate, 0)

	var checked []string

	for i, t := range tiles {

		// Check already
		if ContainsTile(checked, t) {
			continue
		}

		checked = append(checked, t)

		var newTiles []string
		newTiles = append(newTiles, tiles[:i]...)
		newTiles = append(newTiles, tiles[i+1:]...)

		// Check if it is ready hand condition
		state := Resolve(tileSetDef, newTiles)
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
