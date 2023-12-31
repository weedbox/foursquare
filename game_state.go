package foursquare

type GameState struct {
	GameID    string        `json:"game_id"`
	CreatedAt int64         `json:"created_at"`
	UpdatedAt int64         `json:"updated_at"`
	Meta      Meta          `json:"meta"`
	Players   []PlayerState `json:"players"`
	Status    Status        `json:"status"`
	Result    *Result       `json:"result,omitempty"`
}

type Meta struct {
	TileSetDef    *TileSetDef `json:"tileset_def"`
	HandTileCount int         `json:"handtile_count"`
	PlayerCount   int         `json:"player_count"`
	WinningStreak int         `json:"winning_streak"`
	Dices         []int       `json:"dices"`
	Tiles         []string    `json:"tiles"`
}

type PlayerState struct {
	Idx            int       `json:"idx"`
	IsBanker       bool      `json:"is_banker"`
	IsReadyHand    bool      `json:"is_ready_hand"`
	Hand           *Hand     `json:"hand"`
	AllowedActions []*Action `json:"allowed_actions"`
}

type Status struct {
	CurrentEvent              string   `json:"cur_event"`
	CurrentTileSetPosition    int      `json:"cur_tpos"`
	CurrentSupplementPosition int      `json:"cur_spos"`
	CurrentPlayer             int      `json:"cur_player"`
	DiscardArea               []string `json:"discard_area"`
}

type Result struct {
	IsDrawnGame      bool                 `json:"is_drawn_game"`
	DiscardingPlayer int                  `json:"discarding_player,omitempty"`
	WinningTile      string               `json:"winning_tile,omitempty"`
	Winners          map[int]WinnerResult `json:"winners,omitempty"`
}

type WinnerResult struct {
	Points     int               `json:"points"`
	Conditions map[PointType]int `json:"conditions"`
}

func NewGameState() *GameState {
	return &GameState{}
}

func (ps *PlayerState) IsAllowedAction(action string) bool {

	for _, a := range ps.AllowedActions {
		if a.Name == action {
			return true
		}
	}

	return false
}

func (ps *PlayerState) ResetAllowedActions() {
	ps.AllowedActions = make([]*Action, 0)
}

func (ps *PlayerState) AllowAction(a *Action) {
	ps.AllowedActions = append(ps.AllowedActions, a)
}

func (ps *PlayerState) AllowActions(actions []*Action) {
	for _, a := range actions {
		ps.AllowedActions = append(ps.AllowedActions, a)
	}
}
