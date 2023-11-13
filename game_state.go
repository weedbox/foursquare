package foursquare

type GameState struct {
	GameID    string        `json:"game_id"`
	CreatedAt int64         `json:"created_at"`
	UpdatedAt int64         `json:"updated_at"`
	Meta      Meta          `json:"meta"`
	Players   []PlayerState `json:"players"`
	Status    Status        `json:"status"`
}

type Meta struct {
	HandTileCount int      `json:"handtile_count"`
	PlayerCount   int      `json:"player_count"`
	Dices         []int    `json:"dices"`
	Tiles         []string `json:"tiles"`
}

type PlayerState struct {
	Idx         int      `json:"idx"`
	IsBanker    bool     `json:"is_banker"`
	IsReadyHand bool     `json:"is_ready_hand"`
	Hand        *Hand    `json:"hand"`
	Actions     []string `json:"actions"`
}

type Status struct {
	CurrentTileSetPosition    int      `json:"cur_tpos"`
	CurrentSupplementPosition int      `json:"cur_spos"`
	CurrentPlayer             int      `json:"cur_player"`
	DiscardArea               []string `json:"discard_area"`
}

func NewGameState() *GameState {
	return &GameState{}
}
