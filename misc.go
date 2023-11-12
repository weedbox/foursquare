package foursquare

import (
	"math/rand"
	"time"
)

func RollDices() []int {

	rand.Seed(time.Now().UnixNano())

	return []int{
		rand.Intn(6) + 1,
		rand.Intn(6) + 1,
	}
}
