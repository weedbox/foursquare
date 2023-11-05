package foursquare

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Resolver_CheckTiles_NoEyes_Win(t *testing.T) {

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

	r := NewResolver(StandardSetOfTiles)

	for _, tiles := range tileSets {
		isWin := r.checkTiles(tiles, false, &ResolverRules{
			Triplet:  true,
			Straight: true,
		})
		assert.True(t, isWin, tiles)
	}

}

func Test_Resolver_CheckTiles_HasEyes_Win(t *testing.T) {

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

	r := NewResolver(StandardSetOfTiles)

	for _, tiles := range tileSets {
		isWin := r.checkTiles(tiles, true, &ResolverRules{
			Triplet:  true,
			Straight: true,
		})
		assert.True(t, isWin, tiles)
	}

}

func Test_Resolver_CheckTiles_NoEyes_Win_OnlyTriplet(t *testing.T) {

	tileSets := [][]string{
		{"I1", "I1", "I1", "I2", "I2", "I2", "I3", "I3", "I3"},
	}

	r := NewResolver(StandardSetOfTiles)

	for _, tiles := range tileSets {
		isWin := r.checkTiles(tiles, false, &ResolverRules{
			Triplet:  true,
			Straight: false,
		})
		assert.True(t, isWin, tiles)
	}

}

func Test_Resolver_CheckTiles_HasEyes_Win_OnlyTriplet(t *testing.T) {

	tileSets := [][]string{
		{"I1", "I1", "I1", "I2", "I2", "I2", "I3", "I3", "I3", "I4", "I4"},
	}

	r := NewResolver(StandardSetOfTiles)

	for _, tiles := range tileSets {
		isWin := r.checkTiles(tiles, true, &ResolverRules{
			Triplet:  true,
			Straight: false,
		})
		assert.True(t, isWin, tiles)
	}

}

func Test_Resolver_CheckTiles_NoEyes_NotWin(t *testing.T) {

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

	r := NewResolver(StandardSetOfTiles)

	for _, tiles := range tileSets {
		isWin := r.checkTiles(tiles, false, &ResolverRules{
			Triplet:  true,
			Straight: true,
		})
		assert.False(t, isWin, tiles)
	}

}

func Test_Resolver_CheckTiles_HasEyes_NotWin(t *testing.T) {

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

	r := NewResolver(StandardSetOfTiles)

	for _, tiles := range tileSets {
		isWin := r.checkTiles(tiles, true, &ResolverRules{
			Triplet:  true,
			Straight: true,
		})
		assert.False(t, isWin, tiles)
	}

}

func Test_Resolver_CheckTiles_NoEyes_NotWin_OnlyTriplet(t *testing.T) {

	tileSets := [][]string{
		{"I1", "I1", "I2", "I2", "I3", "I3", "I4", "I4", "I4"},
		{"I1", "I1", "I2", "I2", "I3", "I3", "I4", "I4"},
	}

	r := NewResolver(StandardSetOfTiles)

	for _, tiles := range tileSets {
		isWin := r.checkTiles(tiles, false, &ResolverRules{
			Triplet:  true,
			Straight: false,
		})
		assert.False(t, isWin, tiles)
	}

}

func Test_Resolver_FigureEyes(t *testing.T) {

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

	r := NewResolver(StandardSetOfTiles)

	rules := &ResolverRules{
		Triplet:  true,
		Straight: true,
	}

	for _, c := range cases {
		candidates := r.figureEyesCandidates(c.Tiles, rules)
		tile, leftTiles := r.figureEyesWithCandidates(c.Tiles, candidates, rules)
		assert.Equal(t, c.Eyes, tile, c.Tiles)
		assert.ElementsMatch(t, c.LeftTiles, leftTiles)
	}

}

func Test_Resolver_FigureEyes_HonorTiles(t *testing.T) {

	tileSets := map[string][]string{
		"I1": {"I1", "I1", "I2", "I2", "I2", "I4", "I4", "I4"},
	}

	r := NewResolver(StandardSetOfTiles)

	rules := &ResolverRules{
		Triplet:  true,
		Straight: false,
	}

	for answer, tiles := range tileSets {
		candidates := r.figureEyesCandidates(tiles, rules)
		tile, _ := r.figureEyesWithCandidates(tiles, candidates, rules)
		assert.Equal(t, answer, tile, tiles)
	}

}

func Test_Resolver_FigureReadyHandConditions(t *testing.T) {

	cases := []struct {
		IsReadyHand bool
		Candidates  []string
		Tiles       []string
	}{
		{
			true,
			[]string{"W4", "W6", "W7", "W9"},
			[]string{"W5", "W6", "W7", "W7", "W8", "W8", "W8"},
		},
		{
			true,
			[]string{"W1", "W2", "W3", "W4"},
			[]string{"W1", "W1", "W1", "W2", "W2", "W3", "W3"},
		},
		{
			true,
			[]string{"W3", "W6", "W9"},
			[]string{"W4", "W5", "W6", "W7", "W8", "W9", "W9", "W9"},
		},
		{
			true,
			[]string{"W3", "W6", "W7", "W8", "W9"},
			[]string{"W4", "W5", "W6", "W7", "W7", "W7", "W8", "W9", "W9", "W9"},
		},
		{
			true,
			[]string{"W1", "W4", "W7"},
			[]string{"W2", "W3", "W4", "W5", "W6"},
		},
		{
			true,
			[]string{"W1", "W4", "W7"},
			[]string{"W1", "W2", "W3", "W4", "W5", "W6", "W7"},
		},
		{
			true,
			[]string{"W1", "W2", "W3", "W4"},
			[]string{"W2", "W3", "W3", "W3", "W4", "W4", "W4"},
		},
		{
			true,
			[]string{"W3", "W4", "W5", "W6"},
			[]string{"W3", "W3", "W3", "W4", "W4", "W4", "W5"},
		},
		{
			true,
			[]string{"W3", "W5", "W6", "W8"},
			[]string{"W4", "W4", "W4", "W5", "W5", "W6", "W7"},
		},
		{
			true,
			[]string{"W2", "W3", "W5", "W8"},
			[]string{"W3", "W4", "W4", "W4", "W5", "W6", "W7"},
		},
		{
			true,
			[]string{"W4", "W5", "W6", "W7"},
			[]string{"W4", "W4", "W4", "W5", "W5", "W6", "W6"},
		},
		{
			true,
			[]string{"W3", "W5", "W6", "W8", "W9"},
			[]string{"W4", "W4", "W4", "W5", "W6", "W7", "W8"},
		},
		{
			true,
			[]string{"W3", "W4", "W5", "W6", "W7"},
			[]string{"W4", "W4", "W4", "W5", "W6", "W6", "W6"},
		},
		{
			true,
			[]string{"W2", "W3", "W4", "W5", "W6", "W9"},
			[]string{"W3", "W3", "W3", "W4", "W5", "W5", "W5", "W6", "W7", "W8"},
		},
		{
			true,
			[]string{"W2", "W3", "W5", "W6", "W8", "W9"},
			[]string{"W1", "W1", "W1", "W2", "W3", "W4", "W5", "W6", "W7", "W8"},
		},
		{
			true,
			[]string{"W1", "W2", "W3", "W4", "W5", "W6", "W7"},
			[]string{"W1", "W1", "W1", "W2", "W3", "W4", "W5", "W6", "W6", "W6"},
		},
		{
			true,
			[]string{"W2", "W3", "W4", "W5", "W6", "W7", "W8"},
			[]string{"W3", "W3", "W4", "W4", "W5", "W5", "W6", "W6", "W7", "W7", "W8", "W8", "W8"},
		},
		{
			true,
			[]string{"W1", "W2", "W3", "W4", "W5", "W6", "W7", "W8"},
			[]string{"W2", "W2", "W2", "W3", "W4", "W5", "W6", "W7", "W7", "W7"},
		},
		{
			true,
			[]string{"W1", "W2", "W3", "W4", "W5", "W6", "W7", "W8", "W9"},
			[]string{"W1", "W1", "W1", "W2", "W3", "W4", "W5", "W6", "W7", "W8", "W9", "W9", "W9"},
		},
	}

	r := NewResolver(StandardSetOfTiles)

	rules := &ResolverRules{
		Triplet:  true,
		Straight: true,
	}

	for _, c := range cases {
		isReadyHand, candidates := r.figureReadyHandConditions(TileSuitWan, c.Tiles, rules)
		assert.Equal(t, c.IsReadyHand, isReadyHand, c.Tiles)
		assert.ElementsMatch(t, c.Candidates, candidates)
	}

}

func Test_Resolver_Resolve_Win(t *testing.T) {

	cases := []struct {
		IsWin       bool
		IsReadyHand bool
		Candidates  []string
		Tiles       []string
	}{
		{
			true,
			false,
			[]string{},
			[]string{"T9", "T9", "T9", "I3", "I3", "I3", "D1", "D1", "D1", "T1", "T1", "T1", "T5", "T5", "T5", "T7", "T7"},
		},
		{
			true,
			false,
			[]string{},
			[]string{"T1", "T2", "T3", "T4", "T5", "T6", "B1", "B2", "B3", "B4", "B5", "B6", "W1", "W2", "W3", "I2", "I2"},
		},
	}

	r := NewResolver(StandardSetOfTiles)

	for _, c := range cases {
		state := r.Resolve(c.Tiles)
		assert.Equal(t, c.IsWin, state.IsWin, c.Tiles)
		assert.Equal(t, c.IsReadyHand, state.IsReadyHand, c.Tiles)
		assert.ElementsMatch(t, c.Candidates, state.ReadyHandCandidates)
	}

}

func Test_Resolver_Resolve_ReadyHand(t *testing.T) {

	cases := []struct {
		IsWin       bool
		IsReadyHand bool
		Candidates  []string
		Tiles       []string
	}{
		{
			false,
			true,
			[]string{"T6", "T7", "T8"},
			[]string{"T9", "T9", "T9", "I3", "I3", "I3", "D1", "D1", "D1", "T1", "T1", "T1", "T5", "T5", "T5", "T7"},
		},
		{
			// Two candidates for eyes
			false,
			true,
			[]string{"W3", "I2"},
			[]string{"T1", "T2", "T3", "T4", "T5", "T6", "B1", "B2", "B3", "B4", "B5", "B6", "W3", "W3", "I2", "I2"},
		},
		{
			false,
			true,
			[]string{"W3"},
			[]string{"T1", "T2", "T3", "T4", "T5", "T6", "B1", "B2", "B3", "B4", "B5", "B6", "W3", "I2", "I2", "I2"},
		},
		{
			false,
			true,
			[]string{"I2"},
			[]string{"T1", "T2", "T3", "T4", "T5", "T6", "B1", "B2", "B3", "B4", "B5", "B6", "W3", "W3", "W3", "I2"},
		},
	}

	r := NewResolver(StandardSetOfTiles)

	for _, c := range cases {
		state := r.Resolve(c.Tiles)
		assert.Equal(t, c.IsWin, state.IsWin, c.Tiles)
		assert.Equal(t, c.IsReadyHand, state.IsReadyHand, c.Tiles)
		assert.ElementsMatch(t, c.Candidates, state.ReadyHandCandidates, c.Tiles)
	}

}
