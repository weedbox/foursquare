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
	ErrInvalidPlayer             = errors.New("game: invalid player")
	ErrInvalidReaction           = errors.New("game: invalid reaction")
	ErrInvalidAction             = errors.New("game: invalid action")
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

func (g *Game) GetPlayer(playerIdx int) *PlayerState {

	if playerIdx < 0 || playerIdx >= len(g.gs.Players) {
		return nil
	}

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

	// Reset allowed actions for previous player
	ps := g.GetCurrentPlayer()
	ps.ResetAllowedActions()

	// Next player
	if g.gs.Status.CurrentPlayer == len(g.gs.Players)-1 {
		g.gs.Status.CurrentPlayer = 0
	} else {
		g.gs.Status.CurrentPlayer++
	}

	return g.triggerEvent(GameEvent_PlayerSelected, "normal")
}

func (g *Game) SelectPlayer(playerIdx int, ctx string) error {

	if playerIdx < 0 || playerIdx >= g.gs.Meta.PlayerCount {
		return ErrInvalidPlayer
	}

	// Reset allowed actions for previous player
	ps := g.GetCurrentPlayer()
	ps.ResetAllowedActions()

	// New player
	g.gs.Status.CurrentPlayer = playerIdx

	return g.triggerEvent(GameEvent_PlayerSelected, ctx)
}

func (g *Game) React(playerIdx int, reaction string) error {

	// No any reactions
	if playerIdx == -1 {
		return g.triggerEvent(GameEvent_NoReactions, nil)
	}

	// check if reaction is valid
	ps := g.GetPlayer(playerIdx)
	if ps == nil {
		return ErrInvalidPlayer
	}

	if !ps.IsAllowedAction(reaction) {
		return ErrInvalidReaction
	}

	// Reset allowed actions
	ps.ResetAllowedActions()

	return g.triggerEvent(GameEvent_PlayerReacted, &PlayerReacted{
		PlayerIdx: playerIdx,
		Reaction:  reaction,
	})
}

func (g *Game) DiscardTile(tile string, isReadyHand bool) error {

	ps := g.GetCurrentPlayer()

	if !ps.IsAllowedAction("discard") {
		return ErrInvalidAction
	}

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

	// Preparing allowed actions for player
	ps := g.GetCurrentPlayer()
	ps.AllowedActions = []string{"discard"}

	return g.triggerEvent(GameEvent_WaitForPlayerToDiscardTile, nil)
}

func (g *Game) WaitForPlayerAction() error {

	ps := g.GetCurrentPlayer()

	// Preparing allowed actions for player
	actions := ps.Hand.FigureActions()
	if len(actions) > 0 {
		return g.triggerEvent(GameEvent_DiscardActions, nil)
	}

	return g.triggerEvent(GameEvent_WaitForPlayerAction, nil)
}

func (g *Game) WaitForReaction() error {
	//TODO: Preparing allowed actions for the rest of players
	return g.triggerEvent(GameEvent_WaitForReaction, nil)
}
