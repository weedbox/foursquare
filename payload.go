package foursquare

type GameEventPayload_Win struct {
	DiscardingPlayer int    `json:"discarding_player"`
	WinningTile      string `json:"winning_tile"`
	Winners          []int  `json:"winners"`
}
