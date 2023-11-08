package foursquare

type Kong struct {
	Open      []string `json:"open"`
	Concealed []string `json:"concealed"`
}

type Hand struct {
	Flowers  []string   `json:"flowers"`
	Triplet  []string   `json:"triplets"`
	Straight [][]string `json:"straight"`
	Kong     Kong       `json:"kong"`
	Tiles    []string   `json:"tiles"`
	Draw     []string   `json:"draw"`
}

func NewHand() *Hand {
	return &Hand{
		Flowers:  make([]string, 0),
		Triplet:  make([]string, 0),
		Straight: make([][]string, 0),
		Kong: Kong{
			Open:      make([]string, 0),
			Concealed: make([]string, 0),
		},
		Tiles: make([]string, 0),
		Draw:  make([]string, 0),
	}
}
