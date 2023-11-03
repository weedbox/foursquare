package foursquare

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Resolver_ResolveGroup_Win(t *testing.T) {

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

	r := NewResolver()

	for _, tiles := range tileSets {
		isWin := r.resolveGroup(tiles, &ResolverRules{
			Triplet:  true,
			Straight: true,
		})
		assert.True(t, isWin, tiles)
	}

}

func Test_Resolver_ResolveGroup_Win_OnlyTriplet(t *testing.T) {

	tileSets := [][]string{
		{"I1", "I1", "I1", "I2", "I2", "I2", "I3", "I3", "I3"},
	}

	r := NewResolver()

	for _, tiles := range tileSets {
		isWin := r.resolveGroup(tiles, &ResolverRules{
			Triplet:  true,
			Straight: false,
		})
		assert.True(t, isWin, tiles)
	}

}

func Test_Resolver_ResolveGroup_NotWin(t *testing.T) {

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

	r := NewResolver()

	for _, tiles := range tileSets {
		isWin := r.resolveGroup(tiles, &ResolverRules{
			Triplet:  true,
			Straight: true,
		})
		assert.False(t, isWin, tiles)
	}

}

func Test_Resolver_ResolveGroup_NotWin_OnlyTriplet(t *testing.T) {

	tileSets := [][]string{
		{"I1", "I1", "I2", "I2", "I3", "I3", "I4", "I4", "I4"},
		{"I1", "I1", "I2", "I2", "I3", "I3", "I4", "I4"},
	}

	r := NewResolver()

	for _, tiles := range tileSets {
		isWin := r.resolveGroup(tiles, &ResolverRules{
			Triplet:  true,
			Straight: false,
		})
		assert.False(t, isWin, tiles)
	}

}

func Test_Resolver_FigureEyesOfGroup(t *testing.T) {

	tileSets := map[string][]string{
		"W9": {"W1", "W1", "W2", "W2", "W3", "W3", "W4", "W4", "W4", "W9", "W9"},
		"W3": {"W1", "W1", "W2", "W2", "W3", "W3", "W4", "W4", "W4", "W3", "W3"},
		"W1": {"W1", "W1", "W2", "W2", "W3", "W3", "W4", "W4", "W5", "W2", "W2"},
		"W2": {"W2", "W2", "W3", "W3", "W4", "W4", "W5", "W5"},
	}

	r := NewResolver()

	rules := &ResolverRules{
		Triplet:  true,
		Straight: true,
	}

	for answer, tiles := range tileSets {
		candidates := r.figureEyesCandidatesOfGroup(tiles, rules)
		tile := r.figureEyesWithCandidates(tiles, candidates, rules)
		assert.Equal(t, answer, tile, tiles)
	}

}

func Test_Resolver_FigureEyesOfGroup_HonorTiles(t *testing.T) {

	tileSets := map[string][]string{
		"I1": {"I1", "I1", "I2", "I2", "I2", "I4", "I4", "I4"},
	}

	r := NewResolver()

	rules := &ResolverRules{
		Triplet:  true,
		Straight: false,
	}

	for answer, tiles := range tileSets {
		candidates := r.figureEyesCandidatesOfGroup(tiles, rules)
		tile := r.figureEyesWithCandidates(tiles, candidates, rules)
		assert.Equal(t, answer, tile, tiles)
	}

}
