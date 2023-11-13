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
	Idx            int      `json:"idx"`
	IsBanker       bool     `json:"is_banker"`
	IsReadyHand    bool     `json:"is_ready_hand"`
	Hand           *Hand    `json:"hand"`
	AllowedActions []string `json:"allowed_actions"`
}

type Status struct {
	CurrentEvent              string   `json:"cur_event"`
	CurrentTileSetPosition    int      `json:"cur_tpos"`
	CurrentSupplementPosition int      `json:"cur_spos"`
	CurrentPlayer             int      `json:"cur_player"`
	DiscardArea               []string `json:"discard_area"`
}

func NewGameState() *GameState {
	return &GameState{}
}

func (ps *PlayerState) IsAllowedAction(action string) bool {

	for _, aa := range ps.AllowedActions {
		if aa == action {
			return true
		}
	}

	return false
}
