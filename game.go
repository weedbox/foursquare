package foursquare

type Game struct {
	opts *Options
	gs   *GameState
}

func NewGame(opts *Options) *Game {

	g := &Game{
		opts: opts,
		gs:   NewGameState(),
	}

	// Apply options
	g.gs.Meta.HandTileCount = opts.HandTileCount
	g.gs.Meta.Dices = opts.Dices
	g.gs.Meta.PlayerCount = opts.PlayerCount
	g.gs.Meta.Tiles = opts.Tiles

	return g
}

// StartGame 開始牌局
func (g *Game) StartGame() error {
	// 實現開始牌局的邏輯
	return g.triggerEvent(GameEvent_GameStarted, nil)
}

func (g *Game) InitializeGame() error {

	// Initializing positions for drawing tile
	g.gs.Status.CurrentTileSetPosition = 0
	g.gs.Status.CurrentSupplementPosition = len(g.gs.Meta.Tiles) - 1

	g.initializePlayers()

	return g.triggerEvent(GameEvent_GameInitialized, nil)
}

// WaitForAllPlayersReady 等待玩家準備完成
func (g *Game) WaitForAllPlayersReady() error {
	// 實現等待玩家準備完成的邏輯
	return g.triggerEvent(GameEvent_GameInitialized, nil) //TODO
}

func (g *Game) StartAtBanker() error {
	g.gs.Status.CurrentPlayer = 0
	return g.triggerEvent(GameEvent_PlayerSelected, nil)
}

// CheckPlayerContext 檢查玩家動作的情境
func (g *Game) CheckPlayerContext() error {
	// 實現檢查玩家動作的情境的邏輯
	return g.triggerEvent(GameEvent_Chow, nil)
}

// WaitForPlayerToDiscardTile 等待選一張打出或打出後聽牌
func (g *Game) WaitForPlayerToDiscardTile() error {
	// 實現等待選一張打出或打出後聽牌的邏輯
	return g.triggerEvent(GameEvent_GameInitialized, nil)
}

// DrawSupplementTIle 玩家補牌
func (g *Game) DrawSupplementTile() error {
	// 實現玩家補牌的邏輯
	return g.triggerEvent(GameEvent_GameInitialized, nil)
}

// Draw 玩家摸牌
func (g *Game) Draw() error {
	// 實現玩家摸牌的邏輯
	return g.triggerEvent(GameEvent_Drawn, nil)
}

// WaitForPlayerAction 等待當前玩家動作
func (g *Game) WaitForPlayerAction() error {
	// 實現等待當前玩家動作的邏輯
	return g.triggerEvent(GameEvent_GameInitialized, nil)
}

// WaitForReaction 等待其他玩家反應
func (g *Game) WaitForReaction() error {
	// 實現等待其他玩家反應的邏輯
	return g.triggerEvent(GameEvent_GameInitialized, nil)
}

// NextPlayer 決定下一家為可動作玩家
func (g *Game) NextPlayer() error {
	// 實現決定下一家為可動作玩家的邏輯
	return g.triggerEvent(GameEvent_PlayerSelected, nil)
}

// SelectPlayer 決定反應玩家為可動作玩家
func (g *Game) SelectPlayer() error {
	// 實現決定反應玩家為可動作玩家的邏輯
	return g.triggerEvent(GameEvent_PlayerSelected, nil)
}

// DrawGame 流局
func (g *Game) DrawGame() error {
	// 實現處理流局的邏輯
	return g.triggerEvent(GameEvent_GameDrawn, nil)
}

// DoSettlement 牌局結算
func (g *Game) DoSettlement() error {
	// 實現牌局結算的邏輯
	return g.triggerEvent(GameEvent_Settlement, nil)
}

// CloseGame 結束牌局
func (g *Game) CloseGame() error {
	// 實現結束牌局的邏輯
	return g.triggerEvent(GameEvent_GameClosed, nil)
}
