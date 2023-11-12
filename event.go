package foursquare

import "fmt"

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
	GameEvent_GameClosed:      "GameClosed",
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
	"GameClosed":      GameEvent_GameClosed,
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
	}

	return nil
}

// onGameStarted 處理遊戲開始的事件
func (g *Game) onGameStarted(payload interface{}) error {
	fmt.Println("遊戲開始:", payload)
	return nil
}

// onGameInitialized 處理遊戲初始化的事件
func (g *Game) onGameInitialized(payload interface{}) error {
	fmt.Println("遊戲初始化完成:", payload)
	return nil
}

// onAllPlayersReady 處理所有玩家準備完成的事件
func (g *Game) onAllPlayersReady(payload interface{}) error {
	fmt.Println("所有玩家已準備好:", payload)
	return nil
}

// onPlayerSelected 處理玩家被選定的事件
func (g *Game) onPlayerSelected(payload interface{}) error {
	fmt.Println("玩家已選定:", payload)
	return nil
}

// onChow 處理吃牌動作的事件
func (g *Game) onChow(payload interface{}) error {
	fmt.Println("執行了吃牌動作:", payload)
	return nil
}

// onPung 處理碰牌動作的事件
func (g *Game) onPung(payload interface{}) error {
	fmt.Println("執行了碰牌動作:", payload)
	return nil
}

// onDiscardActions 處理放棄動作的事件
func (g *Game) onDiscardActions(payload interface{}) error {
	fmt.Println("放棄動作:", payload)
	return nil
}

// onKong 處理槓牌動作的事件
func (g *Game) onKong(payload interface{}) error {
	fmt.Println("執行了槓牌動作:", payload)
	return nil
}

// onConcealedKong 處理暗槓動作的事件
func (g *Game) onConcealedKong(payload interface{}) error {
	fmt.Println("執行了暗槓動作:", payload)
	return nil
}

// onNormalState 處理正常情境的事件
func (g *Game) onNormalState(payload interface{}) error {
	fmt.Println("正常情境:", payload)
	return nil
}

// onDrawn 處理玩家摸牌的事件
func (g *Game) onDrawn(payload interface{}) error {
	fmt.Println("玩家已摸牌:", payload)
	return nil
}

// onFlowerTileDrawn 處理玩家抓到花牌的事件
func (g *Game) onFlowerTileDrawn(payload interface{}) error {
	fmt.Println("玩家抓到花牌:", payload)
	return nil
}

// onTileDiscarded 處理牌被打出的事件
func (g *Game) onTileDiscarded(payload interface{}) error {
	fmt.Println("牌已被打出:", payload)
	return nil
}

// onNoReactions 處理其他玩家無反應的事件
func (g *Game) onNoReactions(payload interface{}) error {
	fmt.Println("其他玩家無反應:", payload)
	return nil
}

// onPlayerReacted 處理玩家反應的事件
func (g *Game) onPlayerReacted(payload interface{}) error {
	fmt.Println("玩家已反應:", payload)
	return nil
}

// onNoMoreTiles 處理摸不到牌的事件
func (g *Game) onNoMoreTiles(payload interface{}) error {
	fmt.Println("玩家摸不到牌:", payload)
	return nil
}

// onGameDrawn 處理遊戲流局的事件
func (g *Game) onGameDrawn(payload interface{}) error {
	fmt.Println("遊戲流局:", payload)
	return nil
}

// onWin 處理胡牌的事件
func (g *Game) onWin(payload interface{}) error {
	fmt.Println("玩家胡牌:", payload)
	return nil
}

// onSettlement 處理遊戲結算的事件
func (g *Game) onSettlement(payload interface{}) error {
	fmt.Println("進行遊戲結算:", payload)
	return nil
}
