package foursquare

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Game_InitializeGame_StandardSetTiles(t *testing.T) {

	opts := NewOptions()
	opts.Dices = RollDices()
	opts.Tiles = NewTileSet(StandardSetOfTiles)

	g := NewGame(opts)

	g.InitializeGame()

	tiles := []string{"W1", "W2", "W3", "W4", "W5", "W6", "W7", "W8", "W9", "T1", "T2", "T3", "T4", "T5", "T6", "T7"}
	for _, ps := range g.gs.Players {
		assert.ElementsMatch(t, tiles, ps.Hand.Tiles)
	}
}

func Test_Game_InitializeGame_WithFlowerTiles(t *testing.T) {

	cases := []struct {
		Answer [][]string
		Tiles  []string
	}{
		{
			Answer: [][]string{
				[]string{"W1", "T1"},
				[]string{"W2", "T2"},
				[]string{"W3", "T3"},
				[]string{"T5", "T4"},
			},
			Tiles: []string{"W1", "W2", "W3", "F1", "T1", "T2", "T3", "T4", "T5"},
		},
	}

	for _, c := range cases {

		opts := NewOptions()
		opts.HandTileCount = 2
		opts.Dices = RollDices()
		opts.Tiles = []string{"W1", "W2", "W3", "F1", "T1", "T2", "T3", "T4", "T5"}
		g := NewGame(opts)
		g.InitializeGame()

		for i, ans := range c.Answer {
			ps := g.gs.Players[i]
			assert.ElementsMatch(t, ans, ps.Hand.Tiles)
		}
	}
}

func Test_Game_WaitForReady(t *testing.T) {

	opts := NewOptions()
	opts.Dices = RollDices()
	opts.Tiles = NewTileSet(StandardSetOfTiles)

	g := NewGame(opts)

	g.WaitForReady()

	assert.Equal(t, g.gs.Status.CurrentEvent, GetGameEventSymbols(GameEvent_WaitForReady))
}
