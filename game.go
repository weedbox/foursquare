package foursquare

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrNoTiles                   = errors.New("game: no tiles")
	ErrInsufficientNumberOfDices = errors.New("game: insufficient number of dices")
	ErrPlayerHasNoSuchTile       = errors.New("game: player has no such tile")
)

type Game struct {
	gs *GameState
}

func NewGame(opts *Options) *Game {

	g := &Game{
		gs: NewGameState(),
	}

	// Apply options
	g.gs.Meta.HandTileCount = opts.HandTileCount
	g.gs.Meta.Dices = opts.Dices
	g.gs.Meta.PlayerCount = opts.PlayerCount
	g.gs.Meta.Tiles = opts.Tiles

	return g
}

func NewGameWithState(gs *GameState) *Game {
	return &Game{
		gs: gs,
	}
}

func (g *Game) GetCurrentPlayer() *PlayerState {
	return &g.gs.Players[g.gs.Status.CurrentPlayer]
}

func (g *Game) StartGame() error {

	if len(g.gs.Meta.Dices) != 2 {
		return ErrInsufficientNumberOfDices
	}

	if len(g.gs.Meta.Tiles) == 0 {
		return ErrNoTiles
	}

	g.gs.GameID = uuid.New().String()
	g.gs.CreatedAt = time.Now().Unix()
	g.gs.UpdatedAt = time.Now().Unix()

	return g.triggerEvent(GameEvent_GameStarted, nil)
}

func (g *Game) InitializeGame() error {

	// Initializing positions for drawing tile
	g.gs.Status.CurrentTileSetPosition = 0
	g.gs.Status.CurrentSupplementPosition = len(g.gs.Meta.Tiles) - 1

	g.initializePlayers()

	return g.triggerEvent(GameEvent_GameInitialized, nil)
}

func (g *Game) StartAtBanker() error {
	g.gs.Status.CurrentPlayer = 0
	return g.triggerEvent(GameEvent_PlayerSelected, nil)
}

func (g *Game) CheckPlayerContext(ctx string) error {

	switch ctx {
	case "chow":
		return g.triggerEvent(GameEvent_Chow, nil)
	case "pung":
		return g.triggerEvent(GameEvent_Kong, nil)
	case "kong":
		return g.triggerEvent(GameEvent_Kong, nil)
	case "win":
		return g.triggerEvent(GameEvent_Win, nil)
	}

	return g.triggerEvent(GameEvent_NormalState, nil)
}

func (g *Game) DrawSupplementTile() error {

	ps := g.GetCurrentPlayer()

	tile, flowerTiles := g.drawSupplementTile()

	if tile == "" {
		return g.triggerEvent(GameEvent_NoMoreTiles, nil)
	}

	ps.Hand.Flowers = append(ps.Hand.Flowers, flowerTiles...)
	ps.Hand.Deal([]string{tile})

	return g.triggerEvent(GameEvent_Drawn, nil)
}

func (g *Game) Draw() error {

	tiles := g.dealTiles(1)
	if len(tiles) == 0 {
		return g.triggerEvent(GameEvent_NoMoreTiles, nil)
	}

	ps := g.GetCurrentPlayer()

	if TileSuit(tiles[0][0:1]) == TileSuitFlower {
		ps.Hand.Flowers = append(ps.Hand.Flowers, tiles...)
		return g.triggerEvent(GameEvent_FlowerTileDrawn, nil)
	}

	ps.Hand.Deal(tiles)

	return g.triggerEvent(GameEvent_Drawn, nil)
}

func (g *Game) NextPlayer() error {

	if g.gs.Status.CurrentPlayer == len(g.gs.Players)-1 {
		g.gs.Status.CurrentPlayer = 0
	} else {
		g.gs.Status.CurrentPlayer++
	}

	return g.triggerEvent(GameEvent_PlayerSelected, "normal")
}

// SelectPlayer 決定反應玩家為可動作玩家
func (g *Game) SelectPlayer() error {
	// 實現決定反應玩家為可動作玩家的邏輯
	return g.triggerEvent(GameEvent_PlayerSelected, nil)
}

func (g *Game) DiscardTile(tile string, isReadyHand bool) error {

	ps := g.GetCurrentPlayer()

	if ps.IsReadyHand {
		ps.Hand.DiscardDrawTile()
	} else {

		if !ps.Hand.Discard(tile) {
			return ErrPlayerHasNoSuchTile
		}

		if isReadyHand {
			ps.IsReadyHand = true
		}
	}

	g.gs.Status.DiscardArea = append(g.gs.Status.DiscardArea, tile)

	return g.triggerEvent(GameEvent_TileDiscarded, nil)
}

func (g *Game) DrawGame() error {
	return g.triggerEvent(GameEvent_GameDrawn, nil)
}

// DoSettlement 牌局結算
func (g *Game) DoSettlement() error {
	// 實現牌局結算的邏輯
	return g.triggerEvent(GameEvent_Settlement, nil)
}

func (g *Game) CloseGame() error {
	return g.triggerEvent(GameEvent_GameClosed, nil)
}

// Wait for external input
func (g *Game) WaitForAllPlayersReady() error {
	return g.triggerEvent(GameEvent_WaitForAllPlayersReady, nil)
}

func (g *Game) WaitForPlayerToDiscardTile() error {
	return g.triggerEvent(GameEvent_WaitForPlayerToDiscardTile, nil)
}

func (g *Game) WaitForPlayerAction() error {
	return g.triggerEvent(GameEvent_WaitForPlayerAction, nil)
}

func (g *Game) WaitForReaction() error {
	return g.triggerEvent(GameEvent_WaitForReaction, nil)
}
