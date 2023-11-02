package foursquare

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewTileSet(t *testing.T) {

	tiles := NewTileSet(StandardSetOfTiles)

	assert.Equal(t, 144, len(tiles))
}

func Test_ShuffleTiles(t *testing.T) {

	tiles := NewTileSet(StandardSetOfTiles)

	tiles = ShuffleTiles(tiles)

	assert.Equal(t, 144, len(tiles))
}
