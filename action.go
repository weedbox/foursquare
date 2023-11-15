package foursquare

type Action struct {
	Name                string              `json:"name"`
	Candidates          [][]string          `json:"candidates,omitempty"`
	ReadyHandCandidates []*DiscardCandidate `json:"ready_hand_candidates,omitempty"`
}

func (a *Action) AddCandidate(c []string) {
	a.Candidates = append(a.Candidates, c)
}
