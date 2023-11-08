package foursquare

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_TileParser_CheckWinningTiles_NoEyes_Win(t *testing.T) {

	tileSets := [][]string{
		{"W1", "W2", "W3", "W4", "W5", "W6", "W7", "W8", "W9"},
		{"W1", "W1", "W1", "W4", "W5", "W6", "W7", "W8", "W9"},
		{"W1", "W1", "W1", "W5", "W5", "W5", "W7", "W8", "W9"},
		{"W1", "W1", "W1", "W5", "W5", "W5", "W9", "W9", "W9"},
		{"W1", "W1", "W1", "W1", "W2", "W3", "W9", "W9", "W9"},
		{"W1", "W1", "W1", "W1", "W2", "W3", "W3", "W3", "W3"},
		{"W1", "W1", "W1", "W2", "W2", "W3", "W3", "W4", "W4"},
		{"W1", "W1", "W1", "W2", "W2", "W3", "W3", "W4", "W4"},
		{"W1", "W1", "W2", "W2", "W3", "W3", "W4", "W4", "W4"},
	}

	for _, tiles := range tileSets {
		isWin := CheckWinningTiles(tiles, false, &ResolverRules{
			Triplet:  true,
			Straight: true,
		})
		assert.True(t, isWin, tiles)
	}

}

func Test_TileParser_CheckWinningTiles_HasEyes_Win(t *testing.T) {

	tileSets := [][]string{
		{"W1", "W2", "W3", "W4", "W5", "W6", "W7", "W8", "W9", "W3", "W3"},
		{"W1", "W1", "W1", "W4", "W5", "W6", "W7", "W8", "W9", "W5", "W5"},
		{"W1", "W1", "W1", "W5", "W5", "W5", "W7", "W8", "W9", "W7", "W7"},
		{"W1", "W1", "W1", "W5", "W5", "W5", "W9", "W9", "W9", "W2", "W2"},
		{"W1", "W1", "W1", "W1", "W2", "W3", "W9", "W9", "W9", "W8", "W8"},
		{"W1", "W1", "W1", "W1", "W2", "W3", "W3", "W3", "W3", "W2", "W2"},
		{"W1", "W1", "W1", "W2", "W2", "W3", "W3", "W4", "W4", "W3", "W3"},
		{"W1", "W1", "W1", "W2", "W2", "W3", "W3", "W4", "W4", "W9", "W9"},
		{"W1", "W1", "W2", "W2", "W3", "W3", "W4", "W4", "W4", "W3", "W3"},
	}

	for _, tiles := range tileSets {
		isWin := CheckWinningTiles(tiles, true, &ResolverRules{
			Triplet:  true,
			Straight: true,
		})
		assert.True(t, isWin, tiles)
	}

}

func Test_TileParser_CheckWinningTiles_NoEyes_Win_OnlyTriplet(t *testing.T) {

	tileSets := [][]string{
		{"I1", "I1", "I1", "I2", "I2", "I2", "I3", "I3", "I3"},
	}

	for _, tiles := range tileSets {
		isWin := CheckWinningTiles(tiles, false, &ResolverRules{
			Triplet:  true,
			Straight: false,
		})
		assert.True(t, isWin, tiles)
	}

}

func Test_TileParser_CheckWinningTiles_HasEyes_Win_OnlyTriplet(t *testing.T) {

	tileSets := [][]string{
		{"I1", "I1", "I1", "I2", "I2", "I2", "I3", "I3", "I3", "I4", "I4"},
	}

	for _, tiles := range tileSets {
		isWin := CheckWinningTiles(tiles, true, &ResolverRules{
			Triplet:  true,
			Straight: false,
		})
		assert.True(t, isWin, tiles)
	}

}

func Test_TileParser_CheckWinningTiles_NoEyes_NotWin(t *testing.T) {

	tileSets := [][]string{
		{"W1", "W1", "W3", "W4", "W5", "W6", "W7", "W8", "W9"},
		{"W1", "W1", "W1", "W1", "W5", "W6", "W7", "W8", "W9"},
		{"W1", "W1", "W1", "W5", "W5", "W9", "W9", "W9", "W9"},
		{"W1", "W2", "W2", "W5", "W5", "W5", "W9", "W9", "W9"},
		{"W1", "W2", "W2", "W2", "W3", "W3", "W3", "W4", "W5"},
		{"W1", "W1", "W1", "W1", "W2", "W3", "W3", "W3", "W6"},
		{"W1", "W1", "W1", "W2", "W2", "W3", "W3", "W4", "W8"},
		{"W1", "W1", "W1", "W2", "W2", "W3", "W4", "W4", "W4"},
	}

	for _, tiles := range tileSets {
		isWin := CheckWinningTiles(tiles, false, &ResolverRules{
			Triplet:  true,
			Straight: true,
		})
		assert.False(t, isWin, tiles)
	}

}

func Test_TileParser_CheckWinningTiles_HasEyes_NotWin(t *testing.T) {

	tileSets := [][]string{
		{"W1", "W1", "W3", "W4", "W5", "W6", "W7", "W8", "W9", "W5", "W5"},
		{"W1", "W1", "W1", "W1", "W5", "W6", "W7", "W8", "W9", "W5", "W5"},
		{"W1", "W1", "W1", "W5", "W5", "W9", "W9", "W9", "W9", "W8", "W8"},
		{"W1", "W2", "W2", "W5", "W5", "W5", "W9", "W9", "W9", "W3", "W3"},
		{"W1", "W2", "W2", "W2", "W3", "W3", "W3", "W4", "W5", "W5", "W5"},
		{"W1", "W1", "W1", "W1", "W2", "W3", "W3", "W3", "W6", "W7", "W7"},
		{"W1", "W1", "W1", "W2", "W2", "W3", "W3", "W4", "W8", "W9", "W9"},
		{"W1", "W1", "W1", "W2", "W2", "W3", "W4", "W4", "W4", "W5", "W5"},
	}

	for _, tiles := range tileSets {
		isWin := CheckWinningTiles(tiles, true, &ResolverRules{
			Triplet:  true,
			Straight: true,
		})
		assert.False(t, isWin, tiles)
	}

}

func Test_TileParser_CheckWinningTiles_NoEyes_NotWin_OnlyTriplet(t *testing.T) {

	tileSets := [][]string{
		{"I1", "I1", "I2", "I2", "I3", "I3", "I4", "I4", "I4"},
		{"I1", "I1", "I2", "I2", "I3", "I3", "I4", "I4"},
	}

	for _, tiles := range tileSets {
		isWin := CheckWinningTiles(tiles, false, &ResolverRules{
			Triplet:  true,
			Straight: false,
		})
		assert.False(t, isWin, tiles)
	}

}

func Test_TileParser_FigureEyes(t *testing.T) {

	cases := []struct {
		Eyes      string
		LeftTiles []string
		Tiles     []string
	}{
		{
			"W9",
			[]string{"W1", "W1", "W2", "W2", "W3", "W3", "W4", "W4", "W4"},
			[]string{"W1", "W1", "W2", "W2", "W3", "W3", "W4", "W4", "W4", "W9", "W9"},
		},
		{
			"W3",
			[]string{"W1", "W1", "W2", "W2", "W3", "W3", "W4", "W4", "W4"},
			[]string{"W1", "W1", "W2", "W2", "W3", "W3", "W4", "W4", "W4", "W3", "W3"},
		},
		{
			"W1",
			[]string{"W2", "W2", "W3", "W3", "W4", "W4", "W5", "W2", "W2"},
			[]string{"W1", "W1", "W2", "W2", "W3", "W3", "W4", "W4", "W5", "W2", "W2"},
		},
		{
			"W2",
			[]string{"W3", "W3", "W4", "W4", "W5", "W5"},
			[]string{"W2", "W2", "W3", "W3", "W4", "W4", "W5", "W5"},
		},
		{
			"I2",
			[]string{},
			[]string{"I2", "I2"},
		},
	}

	rules := &ResolverRules{
		Triplet:  true,
		Straight: true,
	}

	for _, c := range cases {
		candidates := FigureEyesCandidates(c.Tiles, rules)
		tile, leftTiles := FigureEyesWithCandidates(c.Tiles, candidates, rules)
		assert.Equal(t, c.Eyes, tile, c.Tiles)
		assert.ElementsMatch(t, c.LeftTiles, leftTiles)
	}

}

func Test_TileParser_FigureEyes_HonorTiles(t *testing.T) {

	tileSets := map[string][]string{
		"I1": {"I1", "I1", "I2", "I2", "I2", "I4", "I4", "I4"},
	}

	rules := &ResolverRules{
		Triplet:  true,
		Straight: false,
	}

	for answer, tiles := range tileSets {
		candidates := FigureEyesCandidates(tiles, rules)
		tile, _ := FigureEyesWithCandidates(tiles, candidates, rules)
		assert.Equal(t, answer, tile, tiles)
	}

}

func Test_TileParser_ParseTileSegmentations(t *testing.T) {

	cases := []struct {
		Answer  [][]string
		HasEyes bool
		Tiles   []string
	}{
		{
			[][]string{
				{"W1", "W2", "W3"},
				{"W4", "W5", "W6"},
				{"W7", "W8", "W9"},
			},
			false,
			[]string{"W1", "W2", "W3", "W4", "W5", "W6", "W7", "W8", "W9"},
		},
		{
			[][]string{
				{"W1", "W2", "W3"},
				{"W4", "W5", "W6"},
				{"W8", "W8"},
			},
			true,
			[]string{"W1", "W2", "W3", "W4", "W5", "W6", "W8", "W8"},
		},
	}

	for _, c := range cases {
		segments, _ := ParseTileSegmentations(c.Tiles, c.HasEyes, &ResolverRules{
			Triplet:  true,
			Straight: true,
		})
		assert.ElementsMatch(t, c.Answer, segments, c.Tiles)
	}
}

func Test_TileParser_ResolveTileSegmentations(t *testing.T) {

	cases := []struct {
		Answer [][]string
		Tiles  []string
	}{
		{
			// Two suits
			[][]string{
				{"B1", "B2", "B3"},
				{"W4", "W5", "W6"},
				{"W8", "W8"},
			},
			[]string{"B1", "B2", "B3", "W4", "W5", "W6", "W8", "W8"},
		},
	}

	for _, c := range cases {
		segments := ResolveTileSegmentations(c.Tiles)
		assert.ElementsMatch(t, c.Answer, segments, c.Tiles)
	}
}
