package foursquare

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_PointRule_PungHand(t *testing.T) {

	cases := []struct {
		Answer bool
		Hand   *Hand
	}{
		{
			true,
			&Hand{
				Flowers:  []string{"F1", "F2"},
				Triplet:  []string{"T1", "W2"},
				Straight: [][]string{},
				Kong: Kong{
					Open:      []string{},
					Concealed: []string{},
				},
				Tiles: []string{"T6", "T6", "T6", "B9", "B9", "W8", "W8", "W8", "W5", "W5"},
				Draw:  []string{"B9"},
			},
		},
		{
			true,
			&Hand{
				Flowers:  []string{"F1", "F2"},
				Triplet:  []string{"T1", "W2"},
				Straight: [][]string{},
				Kong: Kong{
					Open:      []string{},
					Concealed: []string{},
				},
				Tiles: []string{"T6", "T6", "T6", "B9", "B9", "W8", "W8", "W8", "W5", "W5", "B9"},
				Draw:  []string{},
			},
		},
		{
			false,
			&Hand{
				Flowers: []string{"F1", "F2"},
				Triplet: []string{"T1", "W2"},
				Straight: [][]string{
					{"B3", "B4", "B5"},
				},
				Kong: Kong{
					Open:      []string{},
					Concealed: []string{},
				},
				Tiles: []string{"T6", "T6", "T6", "B9", "B9", "W5", "W5"},
				Draw:  []string{"B9"},
			},
		},
		{
			false,
			&Hand{
				Flowers:  []string{"F1", "F2"},
				Triplet:  []string{"T1", "W2", "B1"},
				Straight: [][]string{},
				Kong: Kong{
					Open:      []string{"B3"},
					Concealed: []string{},
				},
				Tiles: []string{"T6", "T6", "T6", "B9", "B9", "W5", "W5"},
				Draw:  []string{"B9"},
			},
		},
	}

	for i, c := range cases {
		assert.Equal(t, c.Answer, StandardPointRule.PungHand(c.Hand), i)
	}
}

func Test_PointRule_FullFlush(t *testing.T) {

	cases := []struct {
		Answer bool
		Hand   *Hand
	}{
		{
			true,
			&Hand{
				Flowers: []string{"F1", "F2"},
				Triplet: []string{"T1"},
				Straight: [][]string{
					{"T1", "T2", "T3"},
				},
				Kong: Kong{
					Open:      []string{},
					Concealed: []string{},
				},
				Tiles: []string{"T3", "T4", "T5", "T6", "T6", "T6", "T7", "T8", "T9", "T9"},
				Draw:  []string{"T9"},
			},
		},
		{
			false,
			&Hand{
				Flowers: []string{"F1", "F2"},
				Triplet: []string{"B1"},
				Straight: [][]string{
					{"T1", "T2", "T3"},
				},
				Kong: Kong{
					Open:      []string{},
					Concealed: []string{},
				},
				Tiles: []string{"T3", "T4", "T5", "T6", "T6", "T6", "T7", "T8", "T9", "T9"},
				Draw:  []string{"T9"},
			},
		},
	}

	for i, c := range cases {
		assert.Equal(t, c.Answer, StandardPointRule.FullFlush(c.Hand), i)
	}
}

func Test_PointRule_HalfFlush(t *testing.T) {

	cases := []struct {
		Answer bool
		Hand   *Hand
	}{
		{
			false,
			&Hand{
				Flowers: []string{"F1", "F2"},
				Triplet: []string{"T1"},
				Straight: [][]string{
					{"T1", "T2", "T3"},
				},
				Kong: Kong{
					Open:      []string{},
					Concealed: []string{},
				},
				Tiles: []string{"T3", "T4", "T5", "T6", "T6", "T6", "T7", "T8", "T9", "T9"},
				Draw:  []string{"T9"},
			},
		},
		{
			false,
			&Hand{
				Flowers: []string{"F1", "F2"},
				Triplet: []string{"B1"},
				Straight: [][]string{
					{"T1", "T2", "T3"},
				},
				Kong: Kong{
					Open:      []string{},
					Concealed: []string{},
				},
				Tiles: []string{"T3", "T4", "T5", "T6", "T6", "T6", "T7", "T8", "T9", "T9"},
				Draw:  []string{"T9"},
			},
		},
		{
			true,
			&Hand{
				Flowers: []string{"F1", "F2"},
				Triplet: []string{"I1", "D1"},
				Straight: [][]string{
					{"T1", "T2", "T3"},
				},
				Kong: Kong{
					Open:      []string{},
					Concealed: []string{},
				},
				Tiles: []string{"T3", "T4", "T5", "T6", "T6", "T6", "T7", "T8", "T9", "T9"},
				Draw:  []string{"T9"},
			},
		},
	}

	for i, c := range cases {
		assert.Equal(t, c.Answer, StandardPointRule.HalfFlush(c.Hand), i)
	}
}
