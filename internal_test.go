package foursquare

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Internal_DrawTiles(t *testing.T) {

	opts := NewOptions()
	g := NewGame(opts)

	g.gs.Meta.PlayerCount = 4
	g.gs.Meta.Tiles = []string{
		"W1", "F1", "W3", "W3", "F2",
	}
	g.gs.Status.CurrentTileSetPosition = 0
	g.gs.Status.CurrentSupplementPosition = len(g.gs.Meta.Tiles) - 1

	tiles := g.dealTiles(3)

	assert.ElementsMatch(t, []string{"W1", "F1", "W3"}, tiles)
}

func Test_Internal_DrawSupplementTile(t *testing.T) {

	opts := NewOptions()
	g := NewGame(opts)

	g.gs.Meta.PlayerCount = 4
	g.gs.Meta.Tiles = []string{
		"W1", "W2", "W3", "F1", "F2",
	}
	g.gs.Status.CurrentTileSetPosition = 0
	g.gs.Status.CurrentSupplementPosition = len(g.gs.Meta.Tiles) - 1

	tile, flowerTiles := g.drawSupplementTile()

	assert.Equal(t, "W3", tile)
	assert.ElementsMatch(t, []string{"F2", "F1"}, flowerTiles)
}

func Test_Internal_DrawSupplementTiles(t *testing.T) {

	opts := NewOptions()
	g := NewGame(opts)

	g.gs.Meta.PlayerCount = 4
	g.gs.Meta.Tiles = []string{
		"W1", "W2", "W3", "F1", "F2",
	}
	g.gs.Status.CurrentTileSetPosition = 0
	g.gs.Status.CurrentSupplementPosition = len(g.gs.Meta.Tiles) - 1

	tileCount := 3
	tiles, flowerTiles := g.drawSupplementTiles(tileCount)

	assert.Equal(t, tileCount, len(tiles))
	assert.ElementsMatch(t, []string{"W3", "W2", "W1"}, tiles)
	assert.ElementsMatch(t, []string{"F2", "F1"}, flowerTiles)
}
