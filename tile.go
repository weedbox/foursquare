package foursquare

import (
	"fmt"
	"math/rand"
	"time"
)

type TileSuit string

const (
	TileSuitWan    TileSuit = "W"
	TileSuitTong            = "T"
	TileSuitBamboo          = "B"
	TileSuitWind            = "I"
	TileSuitDragon          = "D"
	TileSuitFlower          = "F"
	TileSuitSeason          = "S"
)

type TileDef struct {
	Suit    TileSuit `json:"suit"`
	Numbers int      `json:"numbers"`
	Count   int      `json:"count"`
}

type TilesOptions struct {
	Wan    TileDef `json:"wan"`    // 萬
	Tong   TileDef `json:"tong"`   // 筒
	Bamboo TileDef `json:"bamboo"` // 條
	Wind   TileDef `json:"wind"`   // 東南西北
	Dragon TileDef `json:"dragon"` // 中發白
	Flower TileDef `json:"flower"` // 梅蘭竹菊
	Season TileDef `json:"season"` // 春夏秋冬
}

var StandardSetOfTiles = &TilesOptions{
	Wan:    TileDef{TileSuitWan, 9, 4},
	Tong:   TileDef{TileSuitTong, 9, 4},
	Bamboo: TileDef{TileSuitBamboo, 9, 4},
	Wind:   TileDef{TileSuitWind, 4, 4},
	Dragon: TileDef{TileSuitDragon, 3, 4},
	Flower: TileDef{TileSuitFlower, 4, 1},
	Season: TileDef{TileSuitSeason, 4, 1},
}

func GenTiles(suit TileSuit, numbers int, count int) []string {

	tiles := make([]string, numbers*count)

	k := 0
	for i := 1; i <= numbers; i++ {

		for j := 0; j < count; j++ {
			tiles[k] = fmt.Sprintf("%s%d", suit, i)
			k++
		}
	}

	return tiles
}

func NewTileSet(opt *TilesOptions) []string {

	tiles := make([]string, 0)
	tiles = append(tiles, GenTiles(opt.Wan.Suit, opt.Wan.Numbers, opt.Wan.Count)...)
	tiles = append(tiles, GenTiles(opt.Tong.Suit, opt.Tong.Numbers, opt.Tong.Count)...)
	tiles = append(tiles, GenTiles(opt.Bamboo.Suit, opt.Bamboo.Numbers, opt.Bamboo.Count)...)
	tiles = append(tiles, GenTiles(opt.Wind.Suit, opt.Wind.Numbers, opt.Wind.Count)...)
	tiles = append(tiles, GenTiles(opt.Dragon.Suit, opt.Dragon.Numbers, opt.Dragon.Count)...)
	tiles = append(tiles, GenTiles(opt.Flower.Suit, opt.Flower.Numbers, opt.Flower.Count)...)
	tiles = append(tiles, GenTiles(opt.Season.Suit, opt.Season.Numbers, opt.Season.Count)...)

	return tiles
}

func ShuffleTiles(tiles []string) []string {

	rand.Seed(time.Now().UnixNano())

	rand.Shuffle(len(tiles), func(i, j int) {
		tiles[i], tiles[j] = tiles[j], tiles[i]
	})

	return tiles
}
