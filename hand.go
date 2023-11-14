package foursquare

import (
	"fmt"
	"sort"
	"strconv"
)

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

func (h *Hand) DiscardDrawTile() {

	for i, t := range h.Tiles {
		if t == h.Draw[0] {
			h.Tiles = append(h.Tiles[:i], h.Tiles[i+1:]...)
			h.Draw = []string{}
			break
		}
	}
}

func (h *Hand) DoChow(tile string, selectedTiles []string) error {

	if tile == "" || len(selectedTiles) != 2 {
		return ErrInvalidAction
	}

	newTiles, n := RemoveTiles(h.Tiles, selectedTiles)
	if n != len(selectedTiles) {
		return ErrInvalidAction
	}

	h.Tiles = newTiles

	s := append(selectedTiles, tile)
	sort.Strings(s)

	h.Straight = append(h.Straight, s)

	return nil
}

func (h *Hand) DoPung(tile string) error {

	if tile == "" {
		return ErrInvalidAction
	}

	newTiles, n := RemoveTiles(h.Tiles, []string{tile, tile})
	if n != 2 {
		return ErrInvalidAction
	}

	h.Tiles = newTiles

	h.Triplet = append(h.Triplet, tile)

	return nil
}

func (h *Hand) DoKong(tile string, isConcealed bool) error {

	if tile == "" {
		return ErrInvalidAction
	}

	newTiles, n := RemoveTiles(h.Tiles, []string{tile, tile, tile})
	if n != 3 {
		return ErrInvalidAction
	}

	h.Tiles = newTiles

	if isConcealed {
		if h.Draw[0] != tile {
			return ErrInvalidAction
		}

		h.Kong.Concealed = append(h.Kong.Concealed, tile)
	} else {
		h.Kong.Open = append(h.Kong.Open, tile)
	}

	h.Draw = []string{}

	return nil
}

func (h *Hand) FigureStraightCandidate(tile string) [][]string {

	var candidates [][]string

	suit := tile[0:1]

	tileNumber, err := strconv.Atoi(tile[1:])
	if err != nil {
		return candidates
	}

	possibleCombos := [][]int{
		{tileNumber - 2, tileNumber - 1},
		{tileNumber - 1, tileNumber + 1},
		{tileNumber + 1, tileNumber + 2},
	}

	for _, combo := range possibleCombos {

		if combo[0] < 1 || combo[1] > 9 {
			continue
		}

		t1 := fmt.Sprintf("%s%d", suit, combo[0])
		t2 := fmt.Sprintf("%s%d", suit, combo[1])

		if ContainsTile(h.Tiles, t1) && ContainsTile(h.Tiles, t2) {
			candidates = append(candidates, []string{t1, t2})
		}
	}

	return candidates

}

func (h *Hand) FigureActions() []*Action {

	var actions []*Action

	// Win by self draw
	isWin := CheckWinningTiles(h.Tiles, true, SuitedTileRule)

	if isWin {
		actions = append(actions, &Action{Name: "win"})
	}

	// Concealed kong
	if CountSpecificTile(h.Tiles, h.Draw[0]) == 4 {
		actions = append(actions, &Action{Name: "kong"})
	}

	if len(actions) > 0 {
		actions = append(actions, &Action{Name: "discard"})
	}

	return actions
}

func (h *Hand) FigureReactions(tile string, relativeSeatIdx int) []*Action {

	var actions []*Action

	// Win
	tiles := append(h.Tiles, tile)
	isWin := CheckWinningTiles(tiles, true, &ResolverRules{
		Triplet:  true,
		Straight: true,
	})

	if isWin {
		actions = append(actions, &Action{Name: "win"})
	}

	// Kong
	if CountSpecificTile(h.Tiles, tile) == 3 {
		actions = append(actions, &Action{Name: "kong"})
	}

	// Pung
	if CountSpecificTile(h.Tiles, tile) == 2 {
		actions = append(actions, &Action{Name: "pung"})
	}

	if relativeSeatIdx == 1 {
		// Chow
		candidates := h.FigureStraightCandidate(tile)
		if len(candidates) != 0 {
			actions = append(actions, &Action{
				Name:       "chow",
				Candidates: candidates,
			})
		}
	}

	return actions
}
