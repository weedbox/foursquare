package foursquare

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Play_Around_In_A_Circle(t *testing.T) {

	opts := NewOptions()
	opts.Dices = RollDices()
	opts.Tiles = NewTileSet(StandardSetOfTiles)

	g := NewGame(opts)
	assert.Nil(t, g.InitializeGame())
	assert.Nil(t, g.Ready())

	// Banker has to dicard tile
	assert.Equal(t, g.gs.Status.CurrentEvent, GetGameEventSymbols(GameEvent_WaitForPlayerToDiscardTile))

	player := g.GetCurrentPlayer()
	assert.True(t, player.Hand.Exists("T8"))
	assert.Equal(t, player.Idx, 0)
	assert.True(t, player.IsBanker)
	assert.Equal(t, player.AllowedActions[0].Name, "discard")

	// Dicard tile
	assert.Equal(t, 17, len(player.Hand.Tiles))
	assert.Nil(t, g.DiscardTile("W4"))
	assert.False(t, player.Hand.Exists("W4"))
	assert.Equal(t, 16, len(player.Hand.Tiles))
	assert.True(t, ContainsTile(g.GetState().Status.DiscardArea, "W4"))

	// Waiting for reactions
	assert.Equal(t, g.gs.Status.CurrentEvent, GetGameEventSymbols(GameEvent_WaitForReaction))

	// Second player can do chow
	player = g.GetPlayer(1)
	assert.True(t, player.IsAllowedAction("chow"))
	assert.Equal(t, 3, len(player.AllowedActions[0].Candidates))
	assert.NotNil(t, g.React(2, "chow", []string{})) // Not allowed
	assert.Nil(t, g.React(1, "chow", []string{"W3", "W5"}))

	assert.Equal(t, g.gs.Status.CurrentEvent, GetGameEventSymbols(GameEvent_WaitForPlayerToDiscardTile))
	assert.Equal(t, player.AllowedActions[0].Name, "discard")
	assert.ElementsMatch(t, []string{"W3", "W4", "W5"}, player.Hand.Straight[0])
	assert.Equal(t, 14, len(player.Hand.Tiles))
	assert.False(t, ContainsTile(g.GetState().Status.DiscardArea, "W4")) // W4 was removed from discard area

	// Dicard tile
	assert.Nil(t, g.DiscardTile("W1"))
	assert.False(t, player.Hand.Exists("W1"))
	assert.Equal(t, 13, len(player.Hand.Tiles))
	assert.True(t, ContainsTile(g.GetState().Status.DiscardArea, "W1"))

	// Waiting for reactions
	assert.Equal(t, g.gs.Status.CurrentEvent, GetGameEventSymbols(GameEvent_WaitForReaction))

	// Third player can do chow
	player = g.GetPlayer(2)
	assert.True(t, player.IsAllowedAction("chow"))
	assert.Equal(t, 1, len(player.AllowedActions[0].Candidates)) // W2, W3
	assert.ElementsMatch(t, []string{"W2", "W3"}, player.AllowedActions[0].Candidates[0])
	assert.Nil(t, g.React(2, "chow", []string{"W2", "W3"}))

	assert.Equal(t, g.gs.Status.CurrentEvent, GetGameEventSymbols(GameEvent_WaitForPlayerToDiscardTile))
	assert.Equal(t, player.AllowedActions[0].Name, "discard")
	assert.ElementsMatch(t, []string{"W1", "W2", "W3"}, player.Hand.Straight[0])
	assert.Equal(t, 14, len(player.Hand.Tiles))
	assert.False(t, ContainsTile(g.GetState().Status.DiscardArea, "W1")) // W1 was removed from discard area

	// Dicard tile
	assert.NotNil(t, g.DiscardTile("T8")) // No such tile
	assert.Nil(t, g.DiscardTile("T6"))
	assert.False(t, player.Hand.Exists("T6"))
	assert.Equal(t, 13, len(player.Hand.Tiles))
	assert.True(t, ContainsTile(g.GetState().Status.DiscardArea, "T6"))

	// Waiting for reactions
	assert.Equal(t, g.gs.Status.CurrentEvent, GetGameEventSymbols(GameEvent_WaitForReaction))
	assert.Nil(t, g.React(-1, "", []string{})) // No reactions

	// fourth player can do draw
	assert.Equal(t, g.gs.Status.CurrentEvent, GetGameEventSymbols(GameEvent_WaitForPlayerToDiscardTile))
	player = g.GetPlayer(3)
	assert.True(t, player.Hand.Exists("T8"))
	assert.True(t, player.IsAllowedAction("discard"))
	assert.Equal(t, 17, len(player.Hand.Tiles))
	assert.Nil(t, g.DiscardTile("T8"))
	assert.Equal(t, 0, len(player.Hand.Draw))
	assert.True(t, ContainsTile(g.GetState().Status.DiscardArea, "T8"))

	// Waiting for reactions
	assert.Equal(t, g.gs.Status.CurrentEvent, GetGameEventSymbols(GameEvent_WaitForReaction))
	assert.Nil(t, g.React(-1, "", []string{})) // No reactions

	// Banker's turn
	player = g.GetPlayer(0)
	assert.Equal(t, 1, len(player.Hand.Draw))
	assert.True(t, player.IsAllowedAction("discard"))

	//g.PrintState()
}
