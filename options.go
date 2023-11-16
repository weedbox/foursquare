package foursquare

type Options struct {
	TileSetDef    *TileSetDef `json:"tileset_def"`
	HandTileCount int         `json:"handtile_count"`
	PlayerCount   int         `json:"player_count"`
	WinningStreak int         `json:"winning_streak"`
	Dices         []int       `json:"dices"`
	Tiles         []string    `json:"tiles"`

	InitialHand map[int]*Hand `json:"initial_hand,omitempty"`
}

func NewOptions() *Options {
	return &Options{
		TileSetDef:    StandardSetOfTiles,
		HandTileCount: 16,
		PlayerCount:   4,
		WinningStreak: 0,
		Dices:         make([]int, 0),
		Tiles:         make([]string, 0),
		InitialHand:   nil,
	}
}
