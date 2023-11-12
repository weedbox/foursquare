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
func (g *Game) StartGame() {
	// 實現開始牌局的邏輯
}

func (g *Game) InitializeGame() {

	// Initializing positions for drawing tile
	g.gs.Status.CurrentTileSetPosition = 0
	g.gs.Status.CurrentSupplementPosition = len(g.gs.Meta.Tiles) - 1

	g.initializePlayers()
}

// WaitForAllPlayersReady 等待玩家準備完成
func (g *Game) WaitForAllPlayersReady() {
	// 實現等待玩家準備完成的邏輯
}

// StartAtBanker 決定莊家開始動作
func (g *Game) StartAtBanker() {
	// 實現決定莊家開始動作的邏輯
}

// CheckPlayerContext 檢查玩家動作的情境
func (g *Game) CheckPlayerContext() {
	// 實現檢查玩家動作的情境的邏輯
}

// WaitForPlayerToDiscardTile 等待選一張打出或打出後聽牌
func (g *Game) WaitForPlayerToDiscardTile() {
	// 實現等待選一張打出或打出後聽牌的邏輯
}

// DrawSupplementTIle 玩家補牌
func (g *Game) DrawSupplementTile() {
	// 實現玩家補牌的邏輯
}

// Draw 玩家摸牌
func (g *Game) Draw() {
	// 實現玩家摸牌的邏輯
}

// WaitForPlayerAction 等待當前玩家動作
func (g *Game) WaitForPlayerAction() {
	// 實現等待當前玩家動作的邏輯
}

// WaitForReaction 等待其他玩家反應
func (g *Game) WaitForReaction() {
	// 實現等待其他玩家反應的邏輯
}

// NextPlayer 決定下一家為可動作玩家
func (g *Game) NextPlayer() {
	// 實現決定下一家為可動作玩家的邏輯
}

// SelectPlayer 決定反應玩家為可動作玩家
func (g *Game) SelectPlayer() {
	// 實現決定反應玩家為可動作玩家的邏輯
}

// DrawGame 流局
func (g *Game) DrawGame() {
	// 實現處理流局的邏輯
}

// DoSettlement 牌局結算
func (g *Game) DoSettlement() {
	// 實現牌局結算的邏輯
}

// CloseGame 結束牌局
func (g *Game) CloseGame() {
	// 實現結束牌局的邏輯
}
