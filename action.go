package foursquare

type Action struct {
	Name       string     `json:"name"`
	Candidates [][]string `json:"candidates,omitempty"`
}

func (a *Action) AddCandidate(c []string) {
	a.Candidates = append(a.Candidates, c)
}
