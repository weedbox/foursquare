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
			// 只有碰，手牌剩下眼睛
			true,
			&Hand{
				Flowers:  []string{"F1", "F2"},
				Triplet:  []string{"T1", "W2", "D3", "B4", "B6"},
				Straight: [][]string{},
				Kong: Kong{
					Open:      []string{},
					Concealed: []string{},
				},
				Tiles: []string{"T6", "T6"},
				Draw:  []string{},
			},
		},
		{
			// 只有吃，手牌只剩眼睛
			false,
			&Hand{
				Flowers: []string{"F1", "F2"},
				Triplet: []string{},
				Straight: [][]string{
					{"B3", "B4", "B5"},
					{"W3", "W4", "W5"},
					{"T3", "T4", "T5"},
					{"T6", "T7", "T8"},
					{"W1", "W2", "W3"},
				},
				Kong: Kong{
					Open:      []string{},
					Concealed: []string{},
				},
				Tiles: []string{"W9"},
				Draw:  []string{"W9"},
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
				Flowers:  []string{"F1", "F2"},
				Triplet:  []string{"D1"},
				Straight: [][]string{},
				Kong: Kong{
					Open:      []string{},
					Concealed: []string{},
				},
				Tiles: []string{"D2", "D2", "D2", "D3", "D3", "D3", "D4", "D4", "D4", "I2"},
				Draw:  []string{"I2"},
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

func Test_PointRule_AllHonorsHand(t *testing.T) {

	cases := []struct {
		Answer bool
		Hand   *Hand
	}{
		{
			true,
			&Hand{
				Flowers:  []string{"F1", "F2"},
				Triplet:  []string{"D1"},
				Straight: [][]string{},
				Kong: Kong{
					Open:      []string{},
					Concealed: []string{},
				},
				Tiles: []string{"D2", "D2", "D2", "D3", "D3", "D3", "D4", "D4", "D4", "I2"},
				Draw:  []string{"I2"},
			},
		},
	}

	for i, c := range cases {
		assert.Equal(t, c.Answer, StandardPointRule.AllHonorsHand(c.Hand), i)
	}
}

func Test_PointRule_BigThreeDragons(t *testing.T) {

	cases := []struct {
		Answer bool
		Hand   *Hand
	}{
		{
			true,
			&Hand{
				Flowers:  []string{"F1", "F2"},
				Triplet:  []string{"D1", "D2", "D3"},
				Straight: [][]string{},
				Kong: Kong{
					Open:      []string{},
					Concealed: []string{},
				},
				Tiles: []string{"T1", "T2", "T3", "B1", "B1", "B1", "I1"},
				Draw:  []string{"I1"},
			},
		},
		{
			true,
			&Hand{
				Flowers:  []string{"F1", "F2"},
				Triplet:  []string{"D1", "D2"},
				Straight: [][]string{},
				Kong: Kong{
					Open:      []string{},
					Concealed: []string{},
				},
				Tiles: []string{"D3", "D3", "D3", "T1", "T2", "T3", "B1", "B1", "B1", "I1"},
				Draw:  []string{"I1"},
			},
		},
	}

	for i, c := range cases {
		assert.Equal(t, c.Answer, StandardPointRule.BigThreeDragons(c.Hand), i)
	}
}

func Test_PointRule_BigFourWinds(t *testing.T) {

	cases := []struct {
		Answer bool
		Hand   *Hand
	}{
		{
			true,
			&Hand{
				Flowers:  []string{"F1", "F2"},
				Triplet:  []string{"I1", "I2", "I3", "I4"},
				Straight: [][]string{},
				Kong: Kong{
					Open:      []string{},
					Concealed: []string{},
				},
				Tiles: []string{"B1", "B1", "B1", "T1"},
				Draw:  []string{"T1"},
			},
		},
		{
			true,
			&Hand{
				Flowers:  []string{"F1", "F2"},
				Triplet:  []string{"I1", "I2"},
				Straight: [][]string{},
				Kong: Kong{
					Open:      []string{},
					Concealed: []string{},
				},
				Tiles: []string{"I3", "I3", "I3", "I4", "I4", "I4", "B1", "B1", "B1", "T1"},
				Draw:  []string{"T1"},
			},
		},
	}

	for i, c := range cases {
		assert.Equal(t, c.Answer, StandardPointRule.BigFourWinds(c.Hand), i)
	}
}
