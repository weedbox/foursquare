package foursquare

type Options struct {
	TileSetDef    *TilesOptions `json:"tileset_def"`
	HandTileCount int           `json:"handtile_count"`
	PlayerCount   int           `json:"player_count"`
	Dices         []int         `json:"dices"`
	Tiles         []string      `json:"tiles"`
}

func NewOptions() *Options {
	return &Options{
		TileSetDef:    StandardSetOfTiles,
		HandTileCount: 16,
		PlayerCount:   4,
		Dices:         make([]int, 0),
		Tiles:         make([]string, 0),
	}
}
