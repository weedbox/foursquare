package foursquare

type GameEvent int32

const (
	GameEvent_GameStarted GameEvent = iota
	GameEvent_GameInitialized
	GameEvent_AllPlayersReady
	GameEvent_PlayerSelected
	GameEvent_Chow
	GameEvent_Pung
	GameEvent_DiscardActions
	GameEvent_Kong
	GameEvent_ConcealedKong
	GameEvent_NormalState
	GameEvent_Drawn
	GameEvent_FlowerTileDrawn
	GameEvent_TileDiscarded
	GameEvent_NoReactions
	GameEvent_PlayerReacted
	GameEvent_NoMoreTiles
	GameEvent_GameDrawn
	GameEvent_Win
	GameEvent_Settlement
	GameEvent_GameClosed

	// Events for waiting
	GameEvent_WaitForAllPlayersReady
	GameEvent_WaitForPlayerAction
	GameEvent_WaitForPlayerToDiscardTile
	GameEvent_WaitForReaction
)

var GameEventSymbols = map[GameEvent]string{
	GameEvent_GameStarted:                "GameStarted",
	GameEvent_GameInitialized:            "GameInitialized",
	GameEvent_AllPlayersReady:            "AllPlayersReady",
	GameEvent_PlayerSelected:             "PlayerSelected",
	GameEvent_Chow:                       "Chow",
	GameEvent_Pung:                       "Pung",
	GameEvent_DiscardActions:             "DiscardActions",
	GameEvent_Kong:                       "Kong",
	GameEvent_ConcealedKong:              "ConcealedKong",
	GameEvent_NormalState:                "NormalState",
	GameEvent_Drawn:                      "Drawn",
	GameEvent_FlowerTileDrawn:            "FlowerTileDrawn",
	GameEvent_TileDiscarded:              "TileDiscarded",
	GameEvent_NoReactions:                "NoReactions",
	GameEvent_PlayerReacted:              "PlayerReacted",
	GameEvent_NoMoreTiles:                "NoMoreTiles",
	GameEvent_GameDrawn:                  "GameDrawn",
	GameEvent_Win:                        "Win",
	GameEvent_Settlement:                 "Settlement",
	GameEvent_GameClosed:                 "GameClosed",
	GameEvent_WaitForAllPlayersReady:     "WaitForAllPlayersReady",
	GameEvent_WaitForPlayerAction:        "WaitForPlayerAction",
	GameEvent_WaitForPlayerToDiscardTile: "WaitForPlayerToDiscardTile",
	GameEvent_WaitForReaction:            "WaitForReaction",
}

var GameEventBySymbol = map[string]GameEvent{
	"GameStarted":                GameEvent_GameStarted,
	"GameInitialized":            GameEvent_GameInitialized,
	"AllPlayersReady":            GameEvent_AllPlayersReady,
	"PlayerSelected":             GameEvent_PlayerSelected,
	"Chow":                       GameEvent_Chow,
	"Pung":                       GameEvent_Pung,
	"DiscardActions":             GameEvent_DiscardActions,
	"Kong":                       GameEvent_Kong,
	"ConcealedKong":              GameEvent_ConcealedKong,
	"NormalState":                GameEvent_NormalState,
	"Drawn":                      GameEvent_Drawn,
	"FlowerTileDrawn":            GameEvent_FlowerTileDrawn,
	"TileDiscarded":              GameEvent_TileDiscarded,
	"NoReactions":                GameEvent_NoReactions,
	"PlayerReacted":              GameEvent_PlayerReacted,
	"NoMoreTiles":                GameEvent_NoMoreTiles,
	"GameDrawn":                  GameEvent_GameDrawn,
	"Win":                        GameEvent_Win,
	"Settlement":                 GameEvent_Settlement,
	"GameClosed":                 GameEvent_GameClosed,
	"WaitForAllPlayersReady":     GameEvent_WaitForAllPlayersReady,
	"WaitForPlayerAction":        GameEvent_WaitForPlayerAction,
	"WaitForPlayerToDiscardTile": GameEvent_WaitForPlayerToDiscardTile,
	"WaitForReaction":            GameEvent_WaitForReaction,
}

func (g *Game) triggerEvent(ge GameEvent, payload interface{}) error {

	switch ge {
	case GameEvent_GameStarted:
		return g.onGameStarted(payload)
	case GameEvent_GameInitialized:
		return g.onGameInitialized(payload)
	case GameEvent_AllPlayersReady:
		return g.onAllPlayersReady(payload)
	case GameEvent_PlayerSelected:
		return g.onPlayerSelected(payload)
	case GameEvent_Chow:
		return g.onChow(payload)
	case GameEvent_Pung:
		return g.onPung(payload)
	case GameEvent_DiscardActions:
		return g.onDiscardActions(payload)
	case GameEvent_Kong:
		return g.onKong(payload)
	case GameEvent_ConcealedKong:
		return g.onConcealedKong(payload)
	case GameEvent_NormalState:
		return g.onNormalState(payload)
	case GameEvent_Drawn:
		return g.onDrawn(payload)
	case GameEvent_FlowerTileDrawn:
		return g.onFlowerTileDrawn(payload)
	case GameEvent_TileDiscarded:
		return g.onTileDiscarded(payload)
	case GameEvent_NoReactions:
		return g.onNoReactions(payload)
	case GameEvent_PlayerReacted:
		return g.onPlayerReacted(payload)
	case GameEvent_NoMoreTiles:
		return g.onNoMoreTiles(payload)
	case GameEvent_GameDrawn:
		return g.onGameDrawn(payload)
	case GameEvent_Win:
		return g.onWin(payload)
	case GameEvent_Settlement:
		return g.onSettlement(payload)

	// Wait
	case GameEvent_WaitForAllPlayersReady:
	case GameEvent_WaitForPlayerAction:
	case GameEvent_WaitForPlayerToDiscardTile:
	case GameEvent_WaitForReaction:
	}

	return nil
}

func (g *Game) onGameStarted(payload interface{}) error {
	return g.InitializeGame()
}

func (g *Game) onGameInitialized(payload interface{}) error {
	return g.WaitForAllPlayersReady()
}

func (g *Game) onAllPlayersReady(payload interface{}) error {
	return g.StartAtBanker()
}

// ctx: normal, chow, pung, kong, win
func (g *Game) onPlayerSelected(ctx interface{}) error {

	if ctx == nil {
		return g.CheckPlayerContext("normal")
	}

	return g.CheckPlayerContext(ctx.(string))
}

func (g *Game) onChow(payload interface{}) error {
	return g.WaitForPlayerToDiscardTile()
}

func (g *Game) onPung(payload interface{}) error {
	return g.WaitForPlayerToDiscardTile()
}

func (g *Game) onDiscardActions(payload interface{}) error {
	return g.WaitForPlayerToDiscardTile()
}

func (g *Game) onKong(payload interface{}) error {
	return g.DrawSupplementTile()
}

func (g *Game) onConcealedKong(payload interface{}) error {
	return g.DrawSupplementTile()
}

func (g *Game) onNormalState(payload interface{}) error {
	return g.Draw()
}

func (g *Game) onDrawn(payload interface{}) error {
	return g.WaitForPlayerAction()
}

func (g *Game) onFlowerTileDrawn(payload interface{}) error {
	return g.DrawSupplementTile()
}

func (g *Game) onTileDiscarded(tile interface{}) error {
	return g.WaitForPlayerAction()
}

func (g *Game) onNoReactions(payload interface{}) error {
	return g.NextPlayer()
}

func (g *Game) onPlayerReacted(payload interface{}) error {
	//TODO
	return g.SelectPlayer()
}

func (g *Game) onNoMoreTiles(payload interface{}) error {
	return g.DrawGame()
}

func (g *Game) onGameDrawn(payload interface{}) error {
	return g.DoSettlement()
}

func (g *Game) onWin(payload interface{}) error {
	return g.DoSettlement()
}

func (g *Game) onSettlement(payload interface{}) error {
	return g.CloseGame()
}