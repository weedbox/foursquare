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
)

var GameEventSymbols = map[GameEvent]string{
	GameEvent_GameStarted:     "GameStarted",
	GameEvent_GameInitialized: "GameInitialized",
	GameEvent_AllPlayersReady: "AllPlayersReady",
	GameEvent_PlayerSelected:  "PlayerSelected",
	GameEvent_Chow:            "Chow",
	GameEvent_Pung:            "Pung",
	GameEvent_DiscardActions:  "DiscardActions",
	GameEvent_Kong:            "Kong",
	GameEvent_ConcealedKong:   "ConcealedKong",
	GameEvent_NormalState:     "NormalState",
	GameEvent_Drawn:           "Drawn",
	GameEvent_FlowerTileDrawn: "FlowerTileDrawn",
	GameEvent_TileDiscarded:   "TileDiscarded",
	GameEvent_NoReactions:     "NoReactions",
	GameEvent_PlayerReacted:   "PlayerReacted",
	GameEvent_NoMoreTiles:     "NoMoreTiles",
	GameEvent_GameDrawn:       "GameDrawn",
	GameEvent_Win:             "Win",
	GameEvent_Settlement:      "Settlement",
}

var GameEventBySymbol = map[string]GameEvent{
	"GameStarted":     GameEvent_GameStarted,
	"GameInitialized": GameEvent_GameInitialized,
	"AllPlayersReady": GameEvent_AllPlayersReady,
	"PlayerSelected":  GameEvent_PlayerSelected,
	"Chow":            GameEvent_Chow,
	"Pung":            GameEvent_Pung,
	"DiscardActions":  GameEvent_DiscardActions,
	"Kong":            GameEvent_Kong,
	"ConcealedKong":   GameEvent_ConcealedKong,
	"NormalState":     GameEvent_NormalState,
	"Drawn":           GameEvent_Drawn,
	"FlowerTileDrawn": GameEvent_FlowerTileDrawn,
	"TileDiscarded":   GameEvent_TileDiscarded,
	"NoReactions":     GameEvent_NoReactions,
	"PlayerReacted":   GameEvent_PlayerReacted,
	"NoMoreTiles":     GameEvent_NoMoreTiles,
	"GameDrawn":       GameEvent_GameDrawn,
	"Win":             GameEvent_Win,
	"Settlement":      GameEvent_Settlement,
}
