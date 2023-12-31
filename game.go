package foursquare

import (
	"encoding/json"
	"errors"
	"fmt"
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
	ErrInvalidGameStatus         = errors.New("game: invalid game status")
)

type Game struct {
	initialHand map[int]*Hand
	gs          *GameState
}

func NewGame(opts *Options) *Game {

	g := &Game{
		gs: NewGameState(),
	}

	// Apply options
	g.initialHand = opts.InitialHand
	g.gs.Meta.TileSetDef = opts.TileSetDef
	g.gs.Meta.HandTileCount = opts.HandTileCount
	g.gs.Meta.Dices = opts.Dices
	g.gs.Meta.PlayerCount = opts.PlayerCount
	g.gs.Meta.WinningStreak = opts.WinningStreak
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

	return &g.gs.Players[playerIdx]
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

func (g *Game) Ready() error {
	return g.triggerEvent(GameEvent_Ready, nil)
}

func (g *Game) StartAtBanker() error {
	g.gs.Status.CurrentPlayer = 0
	return g.triggerEvent(GameEvent_PlayerSelected, nil)
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

func (g *Game) Act(action string) error {

	ps := g.GetCurrentPlayer()

	if !ps.IsAllowedAction(action) {
		return ErrInvalidAction
	}

	switch action {
	case "win":

		payload := &GameEventPayload_Win{
			DiscardingPlayer: ps.Idx,
			WinningTile:      ps.Hand.Draw[0],
			Winners:          []int{ps.Idx},
		}

		return g.triggerEvent(GameEvent_Win, payload)
	case "kong":
		err := ps.Hand.DoKong(ps.Hand.Draw[0], true)
		if err != nil {
			return err
		}

		return g.triggerEvent(GameEvent_ConcealedKong, nil)
	}

	return g.triggerEvent(GameEvent_Cancel, nil)
}

func (g *Game) React(playerIdx int, reaction string, selectedTiles []string) error {

	// No one has any reactions
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

	discardingPlayer := g.gs.Status.CurrentPlayer

	g.gs.Status.CurrentPlayer = playerIdx

	// Reset allowed actions
	ps.ResetAllowedActions()

	// Take the last discarded tile
	discardedTile := g.gs.Status.DiscardArea[len(g.gs.Status.DiscardArea)-1]
	g.gs.Status.DiscardArea = g.gs.Status.DiscardArea[:len(g.gs.Status.DiscardArea)-1]

	// do reaction
	switch reaction {
	case "win":

		payload := &GameEventPayload_Win{
			DiscardingPlayer: discardingPlayer,
			WinningTile:      discardedTile,
			Winners:          []int{playerIdx},
		}

		return g.triggerEvent(GameEvent_Win, payload)
	case "kong":

		err := ps.Hand.DoKong(discardedTile, false)
		if err != nil {
			return err
		}

		return g.triggerEvent(GameEvent_Kong, discardedTile)

	case "pung":

		err := ps.Hand.DoPung(discardedTile)
		if err != nil {
			return err
		}

		return g.triggerEvent(GameEvent_Pung, discardedTile)

	case "chow":

		err := ps.Hand.DoChow(discardedTile, selectedTiles)
		if err != nil {
			return err
		}

		return g.triggerEvent(GameEvent_Chow, discardedTile)
	}

	return g.triggerEvent(GameEvent_NoReactions, nil)
}

func (g *Game) DiscardTile(tile string) error {

	ps := g.GetCurrentPlayer()

	if !ps.IsAllowedAction("discard") {
		return ErrInvalidAction
	}

	if !ps.Hand.Discard(tile) {
		return ErrPlayerHasNoSuchTile
	}

	g.gs.Status.DiscardArea = append(g.gs.Status.DiscardArea, tile)

	return g.triggerEvent(GameEvent_TileDiscarded, nil)
}

func (g *Game) ReadyHand(tile string) error {

	ps := g.GetCurrentPlayer()

	if !ps.IsAllowedAction("readyhand") {
		return ErrInvalidAction
	}

	if !ps.Hand.Discard(tile) {
		return ErrPlayerHasNoSuchTile
	}

	ps.IsReadyHand = true

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
func (g *Game) WaitForReady() error {
	return g.triggerEvent(GameEvent_WaitForReady, nil)
}

func (g *Game) WaitForPlayerToDiscardTile() error {

	// Preparing allowed actions for player
	ps := g.GetCurrentPlayer()
	ps.ResetAllowedActions()

	// Discard tile directly if player stay in ready hand condition
	if ps.IsReadyHand {
		return g.DiscardTile(ps.Hand.Draw[0])
	}

	ps.AllowAction(&Action{
		Name: "discard",
	})

	// Figure discard candidates for readyhand condition
	candidates := FigureDiscardCandidatesForReadyHand(g.gs.Meta.TileSetDef, ps.Hand.Tiles)
	if len(candidates) > 0 {
		ps.AllowAction(&Action{
			Name:                "readyhand",
			ReadyHandCandidates: candidates,
		})
	}

	return g.triggerEvent(GameEvent_WaitForPlayerToDiscardTile, nil)
}

func (g *Game) WaitForPlayerAction() error {

	ps := g.GetCurrentPlayer()
	ps.ResetAllowedActions()

	// Figure out actions
	actions := ps.Hand.FigureActions()
	if len(actions) == 0 {

		// No Actions
		return g.triggerEvent(GameEvent_Cancel, nil)
	}

	// Assign allowed actions for player
	ps.AllowActions(actions)

	return g.triggerEvent(GameEvent_WaitForPlayerAction, nil)
}

func (g *Game) WaitForReaction() error {

	discardedTile := g.gs.Status.DiscardArea[len(g.gs.Status.DiscardArea)-1]

	ps := g.GetCurrentPlayer()

	players := g.getPlayersStartingFrom(ps.Idx)
	if len(players) != len(g.gs.Players) {
		return ErrInvalidGameStatus
	}

	// Figure out reactions that player can do
	hasReactors := false
	for i, p := range players {

		p.ResetAllowedActions()

		if p.Idx == ps.Idx {
			continue
		}

		// Assign allowed actions for player
		actions := p.Hand.FigureReactions(g.gs.Meta.TileSetDef, discardedTile, i)
		if len(actions) > 0 {
			hasReactors = true
			p.AllowActions(actions)
		}
	}

	if hasReactors {

		// All wins?
		var winners []int
		for _, p := range players {
			if p.IsAllowedAction("win") {
				winners = append(winners, p.Idx)
			}
		}

		// Oops!
		if len(winners) == len(players)-1 {

			g.resetAllowedActions()

			payload := &GameEventPayload_Win{
				DiscardingPlayer: ps.Idx,
				WinningTile:      discardedTile,
				Winners:          winners,
			}

			return g.triggerEvent(GameEvent_Win, payload)
		}

		return g.triggerEvent(GameEvent_WaitForReaction, nil)
	}

	return g.triggerEvent(GameEvent_NoReactions, nil)
}

func (g *Game) GetState() *GameState {
	return g.gs
}

func (g *Game) PrintState() {
	data, _ := json.Marshal(g.gs)
	fmt.Println(string(data))
}
