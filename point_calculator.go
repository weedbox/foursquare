package foursquare

type PointType int32

const (
	// 基本牌型台數
	MinimalPoints       PointType = iota // 平胡（自摸）
	PungHand                             // 碰碰胡
	HalfFlush                            // 混一色
	FullFlush                            // 清一色
	LittleThreeDragons                   // 小三元
	AllHonorsHand                        // 字一色
	BigThreeDragons                      // 大三元
	ThreeConcealedPungs                  // 三暗刻
	FourConcealedPungs                   // 四暗刻
	FiveConcealedPungs                   // 五暗刻
	SmallFourWinds                       // 小四喜
	BigFourWinds                         // 大四喜

	// 特殊牌型
	HeavenlyHand // 天胡
	EarthlyHand  // 地胡

	// 花牌
	FlowerTiles // 花牌

	// 其他條件
	AfterAKong     // 槓上開花
	LastTileDraw   // 海底撈月
	RobbingTheKong // 搶槓胡
	KongOnDiscard  // 杠上炮
	SingleWait     // 獨聽
	SelfDrawn      // 自摸加底

	// 番種
	MeldedKong    // 明槓
	ConcealedKong // 暗槓

	// 門前清
	ConcealedHand // 門前清
)

type PointRule struct {
	Type  PointType `json:"type"`
	Point int       `json:"point"`
}

type PointCalculator struct {
	Rules map[PointType]PointRule `json:"rules"`
}

var StandardRules map[PointType]PointRule = map[PointType]PointRule{
	MinimalPoints:       {Type: MinimalPoints, Point: 1},       // 平胡（自摸）
	PungHand:            {Type: PungHand, Point: 4},            // 碰碰胡
	HalfFlush:           {Type: HalfFlush, Point: 4},           // 混一色
	FullFlush:           {Type: FullFlush, Point: 8},           // 清一色
	AllHonorsHand:       {Type: AllHonorsHand, Point: 8},       // 字一色
	LittleThreeDragons:  {Type: LittleThreeDragons, Point: 4},  // 小三元
	BigThreeDragons:     {Type: BigThreeDragons, Point: 8},     // 大三元
	ThreeConcealedPungs: {Type: ThreeConcealedPungs, Point: 2}, // 三暗刻
	FourConcealedPungs:  {Type: FourConcealedPungs, Point: 5},  // 四暗刻
	FiveConcealedPungs:  {Type: FiveConcealedPungs, Point: 8},  // 五暗刻
	SmallFourWinds:      {Type: SmallFourWinds, Point: 8},      // 小四喜
	BigFourWinds:        {Type: BigFourWinds, Point: 16},       // 大四喜

	HeavenlyHand: {Type: HeavenlyHand, Point: 16}, // 天胡
	EarthlyHand:  {Type: EarthlyHand, Point: 16},  // 地胡

	FlowerTiles: {Type: FlowerTiles, Point: 1}, // 花牌

	AfterAKong:     {Type: AfterAKong, Point: 1},     // 槓上開花
	LastTileDraw:   {Type: LastTileDraw, Point: 1},   // 海底撈月
	RobbingTheKong: {Type: RobbingTheKong, Point: 1}, // 搶槓胡
	KongOnDiscard:  {Type: KongOnDiscard, Point: 1},  // 杠上炮
	SingleWait:     {Type: SingleWait, Point: 1},     // 獨聽
	SelfDrawn:      {Type: SelfDrawn, Point: 1},      // 自摸加底

	MeldedKong:    {Type: MeldedKong, Point: 1},    // 明槓
	ConcealedKong: {Type: ConcealedKong, Point: 1}, // 暗槓

	ConcealedHand: {Type: ConcealedHand, Point: 1}, // 門前清
}

func NewPointCalculator(rules map[PointType]PointRule) *PointCalculator {
	return &PointCalculator{
		Rules: rules,
	}
}

func (pc *PointCalculator) Calculate(g *Game, ps *PlayerState, hand *Hand) {
	//TODO
}

func (pc *PointCalculator) MinimalPoints(hand *Hand) {
	// 實現判斷平胡（自摸）的邏輯
}

func (pc *PointCalculator) PungHand(hand *Hand) int {

	// 實現判斷碰碰胡的邏輯

	if len(hand.Straight) > 0 || len(hand.Kong.Concealed) > 0 || len(hand.Kong.Open) > 0 {
		return 0
	}

	results := CountByTiles(hand.Tiles)

	foundEyes := false
	for _, r := range results {

		if r == 2 {

			// only one pair of eyes
			if foundEyes {
				return 0
			}

			foundEyes = true
			continue
		}

		if r != 3 {
			return 0
		}
	}

	return pc.Rules[PungHand].Point
}

func (pc *PointCalculator) HalfFlush(hand *Hand) int {

	// 實現判斷混一色的邏輯

	tiles := append(hand.Tiles, hand.Triplet...)
	tiles = append(tiles, hand.Kong.Open...)
	tiles = append(tiles, hand.Kong.Concealed...)

	for _, s := range hand.Straight {
		tiles = append(tiles, s...)
	}

	results := CountBySuits(tiles)

	// Dragon, Winds and one suit
	if len(results) > 3 || len(results) == 1 {
		return 0
	}

	var foundSuit TileSuit
	for suit, _ := range results {
		if suit == TileSuitDragon || suit == TileSuitWind {
			continue
		}

		if len(foundSuit) != 0 {
			return 0
		}

		foundSuit = suit
	}

	return pc.Rules[HalfFlush].Point
}

func (pc *PointCalculator) FullFlush(hand *Hand) int {

	// 實現判斷清一色的邏輯

	tiles := append(hand.Tiles, hand.Triplet...)
	tiles = append(tiles, hand.Kong.Open...)
	tiles = append(tiles, hand.Kong.Concealed...)

	for _, s := range hand.Straight {
		tiles = append(tiles, s...)
	}

	results := CountBySuits(tiles)
	if len(results) != 1 {
		return 0
	}

	for suit, _ := range results {
		// No dragon and wind
		if suit == TileSuitDragon || suit == TileSuitWind {
			return 0
		}
	}

	return pc.Rules[FullFlush].Point
}

func (pc *PointCalculator) LittleThreeDragons(hand *Hand) int {

	// 實現判斷小三元的邏輯

	tiles := hand.Tiles

	for _, t := range hand.Triplet {
		tiles = append(tiles, t, t, t)
	}

	results := CountByTiles(tiles)

	tripletDragons := 0
	pairDragons := 0
	for tile, c := range results {

		if tile == "D1" || tile == "D2" || tile == "D3" {
			if c == 3 {
				tripletDragons++
			} else if c == 2 {
				pairDragons++
			} else {
				return 0
			}
		}
	}

	if tripletDragons != 2 || pairDragons != 1 {
		return 0
	}

	return pc.Rules[LittleThreeDragons].Point
}

func (pc *PointCalculator) AllHonorsHand(hand *Hand) int {

	// 實現判斷字一色的邏輯

	tiles := append(hand.Tiles, hand.Triplet...)
	tiles = append(tiles, hand.Kong.Open...)
	tiles = append(tiles, hand.Kong.Concealed...)

	results := CountBySuits(tiles)
	for suit, _ := range results {
		if suit != TileSuitDragon && suit != TileSuitWind {
			return 0
		}
	}

	return pc.Rules[AllHonorsHand].Point
}

func (pc *PointCalculator) BigThreeDragons(hand *Hand) int {

	// 實現判斷大三元的邏輯

	tiles := hand.Tiles

	for _, t := range hand.Triplet {
		tiles = append(tiles, t, t, t)
	}

	results := CountByTiles(tiles)

	// 中
	count, ok := results["D1"]
	if !ok || count != 3 {
		return 0
	}

	// 發
	count, ok = results["D2"]
	if !ok || count != 3 {
		return 0
	}

	// 白
	count, ok = results["D3"]
	if !ok || count != 3 {
		return 0
	}

	return pc.Rules[BigThreeDragons].Point
}

func (pc *PointCalculator) ThreeConcealedPungs(hand *Hand) int {

	// 實現判斷三暗刻的邏輯

	count := 0
	count += len(hand.Triplet)
	count += len(hand.Kong.Concealed)

	segments := ResolveTileSegmentations(hand.Tiles)
	for _, s := range segments {
		if IsTriplet(s) {
			count++
		}
	}

	if count != 3 {
		return 0
	}

	return pc.Rules[ThreeConcealedPungs].Point
}

func (pc *PointCalculator) FourConcealedPungs(hand *Hand) int {

	// 實現判斷四暗刻的邏輯

	count := 0
	count += len(hand.Triplet)
	count += len(hand.Kong.Concealed)

	segments := ResolveTileSegmentations(hand.Tiles)

	for _, s := range segments {
		if IsTriplet(s) {
			count++
		}
	}

	if count != 4 {
		return 0
	}

	return pc.Rules[FourConcealedPungs].Point
}

func (pc *PointCalculator) FiveConcealedPungs(hand *Hand) int {

	// 實現判斷五暗刻的邏輯

	count := 0
	count += len(hand.Triplet)
	count += len(hand.Kong.Concealed)

	segments := ResolveTileSegmentations(hand.Tiles)

	for _, s := range segments {
		if IsTriplet(s) {
			count++
		}
	}

	if count != 5 {
		return 0
	}

	return pc.Rules[FiveConcealedPungs].Point
}

func (pc *PointCalculator) SmallFourWinds(hand *Hand) int {

	// 實現判斷小四喜的邏輯

	tiles := hand.Tiles

	for _, t := range hand.Triplet {
		tiles = append(tiles, t, t, t)
	}

	results := CountByTiles(tiles)

	tripletWinds := 0
	pairWinds := 0
	for tile, c := range results {

		if tile == "I1" || tile == "I2" || tile == "I3" || tile == "I4" {
			if c == 3 {
				tripletWinds++
			} else if c == 2 {
				pairWinds++
			} else {
				return 0
			}
		}
	}

	if tripletWinds != 3 || pairWinds != 1 {
		return 0
	}

	return pc.Rules[SmallFourWinds].Point
}

func (pc *PointCalculator) BigFourWinds(hand *Hand) int {

	// 實現判斷大四喜的邏輯

	tiles := hand.Tiles

	for _, t := range hand.Triplet {
		tiles = append(tiles, t, t, t)
	}

	results := CountByTiles(tiles)

	tripletCount := 0

	// 東
	count, ok := results["I1"]
	if !ok || count != 3 {
		return 0
	}

	tripletCount++

	// 南
	count, ok = results["I2"]
	if !ok || count != 3 {
		return 0
	}

	tripletCount++

	// 西
	count, ok = results["I3"]
	if !ok || count != 3 {
		return 0
	}

	tripletCount++

	// 北
	count, ok = results["I4"]
	if !ok || count != 3 {
		return 0
	}

	tripletCount++

	// It should be 4 triplets
	if tripletCount > 4 {
		return 0
	}

	return pc.Rules[BigFourWinds].Point
}

func (pc *PointCalculator) HeavenlyHand(g *Game, ps *PlayerState, hand *Hand) int {

	// Not banker
	if !ps.IsBanker {
		return 0
	}

	// Not draw by self
	if len(ps.Hand.Draw) == 0 {
		return 0
	}

	if len(ps.Hand.Kong.Open) > 0 || len(ps.Hand.Kong.Concealed) > 0 {
		return 0
	}

	// Not the first tile
	if len(g.gs.Status.DiscardArea) != 0 {
		return 0
	}

	return pc.Rules[HeavenlyHand].Point
}

// TODO: TBD
func (pc *PointCalculator) EarthlyHand(g *Game, ps *PlayerState, hand *Hand) int {

	// Should not be banker
	if ps.IsBanker {
		return 0
	}

	// No one do special action
	for _, p := range g.gs.Players {
		if len(p.Hand.Kong.Open) > 0 || len(p.Hand.Triplet) > 0 || len(p.Hand.Straight) > 0 {
			return 0
		}
	}

	// Not draw by self
	if len(ps.Hand.Draw) == 0 {
		return 0
	}

	return pc.Rules[EarthlyHand].Point
}

func (pc *PointCalculator) FlowerTiles(hand *Hand) {
	// 實現判斷花牌的邏輯
}

func (pc *PointCalculator) AfterAKong(hand *Hand) {
	// 實現判斷槓上開花的邏輯
}

func (pc *PointCalculator) LastTileDraw(g *Game, hand *Hand) int {

	// 海底撈月

	if len(hand.Draw) == 0 {
		return 0
	}

	if len(g.gs.Meta.Tiles) != g.gs.Status.CurrentTileSetPosition {
		return 0
	}

	return pc.Rules[LastTileDraw].Point
}

func (pc *PointCalculator) RobbingTheKong(hand *Hand) {
	// 實現判斷搶槓胡的邏輯
}

func (pc *PointCalculator) KongOnDiscard(hand *Hand) {
	// 實現判斷杠上炮的邏輯
}

func (pc *PointCalculator) SingleWait(hand *Hand) {
	// 實現判斷獨聽的邏輯
}

func (pc *PointCalculator) SelfDrawn(hand *Hand) int {

	if len(hand.Draw) == 0 {
		return 0
	}

	return pc.Rules[SelfDrawn].Point
}

func (pc *PointCalculator) MeldedKong(hand *Hand) {
	// 實現判斷明槓的邏輯
}

func (pc *PointCalculator) ConcealedKong(hand *Hand) {
	// 實現判斷暗槓的邏輯
}

func (pc *PointCalculator) ConcealedHand(hand *Hand) int {

	// 實現判斷門前清的邏輯

	if len(hand.Triplet) > 0 || len(hand.Straight) > 0 || len(hand.Kong.Concealed) > 0 || len(hand.Kong.Open) > 0 {
		return 0
	}

	return pc.Rules[ConcealedHand].Point
}
