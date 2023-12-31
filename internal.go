package foursquare

func (g *Game) resetAllowedActions() {

	for _, p := range g.gs.Players {
		p.ResetAllowedActions()
	}
}

func (g *Game) getPlayersStartingFrom(idx int) []*PlayerState {

	var players []*PlayerState
	cur := idx
	for i := 0; i < len(g.gs.Players); i++ {

		if cur >= len(g.gs.Players) {
			cur = 0
		}

		p := g.GetPlayer(cur)
		if p == nil {
			continue
		}

		players = append(players, p)

		cur++
	}

	return players
}

func (g *Game) drawTile() (string, []string) {

	var tile string
	var flowerTiles []string

	tile = g.gs.Meta.Tiles[g.gs.Status.CurrentTileSetPosition]
	g.gs.Status.CurrentTileSetPosition++

	if TileSuit(tile[0:1]) == TileSuitFlower {
		flowerTiles = append(flowerTiles, tile)
	}

	t, fts := g.drawSupplementTile()

	tile = t
	flowerTiles = append(flowerTiles, fts...)

	return tile, flowerTiles
}

func (g *Game) dealTiles(count int) []string {

	tiles := make([]string, 0, count)

	finalPos := g.gs.Status.CurrentTileSetPosition + count
	for i := g.gs.Status.CurrentTileSetPosition; i < finalPos; i++ {
		tiles = append(tiles, g.gs.Meta.Tiles[i])
		g.gs.Status.CurrentTileSetPosition++
	}

	return tiles
}

func (g *Game) drawSupplementTile() (string, []string) {

	var tile string
	var flowerTiles []string

	for g.gs.Status.CurrentSupplementPosition >= g.gs.Status.CurrentTileSetPosition {

		t := g.gs.Meta.Tiles[g.gs.Status.CurrentSupplementPosition]

		// Check if it is not flower tile
		if TileSuit(t[0:1]) != TileSuitFlower {
			tile = t
			g.gs.Status.CurrentSupplementPosition--
			break
		}

		flowerTiles = append(flowerTiles, t)
		g.gs.Status.CurrentSupplementPosition--
	}

	return tile, flowerTiles
}

func (g *Game) drawSupplementTiles(count int) ([]string, []string) {

	tiles := make([]string, 0)
	flowerTiles := make([]string, 0)

	for i := 0; i < count; i++ {
		t, fts := g.drawSupplementTile()
		tiles = append(tiles, t)
		flowerTiles = append(flowerTiles, fts...)

	}

	return tiles, flowerTiles
}

func (g *Game) initializeHandTiles() {

	if g.initialHand == nil {
		for i := 0; i < g.gs.Meta.HandTileCount; i++ {

			for _, ps := range g.gs.Players {
				ps.Hand.Tiles = append(ps.Hand.Tiles, g.dealTiles(1)...)
			}
		}
	}

	// Draw supplemement tile for flowers
	for _, ps := range g.gs.Players {

		var newTiles []string
		for _, tile := range ps.Hand.Tiles {

			// Check if it is flower tile
			if TileSuit(tile[0:1]) == TileSuitFlower {
				ps.Hand.Flowers = append(ps.Hand.Flowers, tile)
				continue
			}

			newTiles = append(newTiles, tile)
		}

		ps.Hand.Tiles = newTiles

		// Draw supplement tiles
		tiles, flowerTiles := g.drawSupplementTiles(len(ps.Hand.Flowers))
		ps.Hand.Tiles = append(ps.Hand.Tiles, tiles...)
		ps.Hand.Flowers = append(ps.Hand.Flowers, flowerTiles...)
	}
}

func (g *Game) initializePlayers() {

	// Initializing players
	for i := 0; i < g.gs.Meta.PlayerCount; i++ {

		ps := PlayerState{
			Idx: i,
		}

		ps.ResetAllowedActions()

		if i == 0 {
			ps.IsBanker = true
		}

		if g.initialHand != nil {
			ps.Hand = g.initialHand[i]
		} else {
			// Empty hand
			ps.Hand = NewHand()
		}

		g.gs.Players = append(g.gs.Players, ps)
	}

	if g.initialHand != nil {
		var initialTiles []string
		for _, h := range g.initialHand {

			// Remove tiles from pool
			tiles := h.GetAllTiles()
			newTiles, _ := RemoveTiles(g.gs.Meta.Tiles, tiles)
			g.gs.Meta.Tiles = newTiles

			initialTiles = append(initialTiles, tiles...)
		}

		g.gs.Meta.Tiles = append(initialTiles, g.gs.Meta.Tiles...)
	}

	g.initializeHandTiles()
}
