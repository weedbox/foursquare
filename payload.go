package foursquare

type PlayerReacted struct {
	PlayerIdx int    `json:"playerIdx"`
	Reaction  string `json:"reaction"` // chow, pung, kong and win
}
