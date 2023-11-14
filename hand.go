package foursquare

type Kong struct {
	Open      []string `json:"open"`
	Concealed []string `json:"concealed"`
}

type Hand struct {
	Flowers  []string   `json:"flowers"`
	Triplet  []string   `json:"triplets"`
	Straight [][]string `json:"straight"`
	Kong     Kong       `json:"kong"`
	Tiles    []string   `json:"tiles"`
	Draw     []string   `json:"draw"`
}

func NewHand() *Hand {
	return &Hand{
		Flowers:  make([]string, 0),
		Triplet:  make([]string, 0),
		Straight: make([][]string, 0),
		Kong: Kong{
			Open:      make([]string, 0),
			Concealed: make([]string, 0),
		},
		Tiles: make([]string, 0),
		Draw:  make([]string, 0),
	}
}

func (h *Hand) Deal(tiles []string) {
	h.Draw = tiles
	h.Tiles = append(h.Tiles, tiles...)
}

func (h *Hand) Exists(tile string) bool {

	for _, t := range h.Tiles {
		if tile == t {
			return true
		}
	}

	return false
}

func (h *Hand) Discard(tile string) bool {

	for i, t := range h.Tiles {
		if t == tile {
			h.Tiles = append(h.Tiles[:i], h.Tiles[i+1:]...)
			return true
		}
	}

	return false
}

func (h *Hand) DiscardDrawTile() bool {

	for i, t := range h.Tiles {
		if t == h.Draw[0] {
			h.Tiles = append(h.Tiles[:i], h.Tiles[i+1:]...)
			return true
		}
	}

	return false
}

func (h *Hand) FigureActions() []string {

	var actions []string

	// Win by self draw
	isWin := CheckWinningTiles(h.Tiles, true, &ResolverRules{
		Triplet:  true,
		Straight: true,
	})

	if isWin {
		actions = append(actions, "win")
	}

	// Concealed kong
	if CountSpecificTile(h.Tiles, h.Draw[0]) == 4 {
		actions = append(actions, "kong")
	}

	if len(actions) > 0 {
		actions = append(actions, "discard")
	}

	return actions
}
