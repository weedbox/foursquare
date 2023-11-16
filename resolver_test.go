package foursquare

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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

	rules := &ResolverRules{
		Triplet:  true,
		Straight: true,
	}

	for _, c := range cases {
		isReadyHand, candidates := FigureReadyHandConditions(StandardSetOfTiles, TileSuitWan, c.Tiles, rules)
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

	for _, c := range cases {
		state := Resolve(StandardSetOfTiles, c.Tiles)
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
		{
			false,
			false,
			[]string{},
			[]string{"T2", "T3", "T4", "T4", "T4", "W1", "W1", "W2", "W3", "W4", "W5", "W6", "W7", "W7", "W7", "D1"},
		},
	}

	for _, c := range cases {
		state := Resolve(StandardSetOfTiles, c.Tiles)
		assert.Equal(t, c.IsWin, state.IsWin, c.Tiles)
		assert.Equal(t, c.IsReadyHand, state.IsReadyHand, c.Tiles)
		assert.ElementsMatch(t, c.Candidates, state.ReadyHandCandidates, c.Tiles)
	}

}
