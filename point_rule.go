package foursquare

type PointType int32

const (
	// 基本牌型台數
	MinimalPoints      PointType = iota // 平胡（自摸）
	PungHand                            // 碰碰胡
	HalfFlush                           // 混一色
	FullFlush                           // 清一色
	LittleThreeDragons                  // 小三元
	AllHonorsHand                       // 字一色
	BigThreeDragons                     // 大三元
	FourConcealedPungs                  // 四暗刻
	SmallFourWinds                      // 小四喜
	BigFourWinds                        // 大四喜

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

type PointRuleCondition struct {
	Type  PointType `json:"type"`
	Point int       `json:"point"`
}

type PointRule struct {
	Conditions map[PointType]PointRuleCondition `json:"conditions"`
}

var StandardPointRule = &PointRule{
	Conditions: map[PointType]PointRuleCondition{
		MinimalPoints:      {Type: MinimalPoints, Point: 1},       // 平胡（自摸）
		PungHand:           {Type: PungHand, Point: 2},            // 碰碰胡
		HalfFlush:          {Type: HalfFlush, Point: 3},           // 混一色
		FullFlush:          {Type: FullFlush, Point: 6},           // 清一色
		LittleThreeDragons: {Type: LittleThreeDragons, Point: 4},  // 小三元
		AllHonorsHand:      {Type: AllHonorsHand, Point: 5},       // 字一色
		BigThreeDragons:    {Type: BigThreeDragons, Point: 8},     // 大三元
		FourConcealedPungs: {Type: FourConcealedPungs, Point: 10}, // 四暗刻
		SmallFourWinds:     {Type: SmallFourWinds, Point: 10},     // 小四喜
		BigFourWinds:       {Type: BigFourWinds, Point: 13},       // 大四喜

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
	},
}

func (pr *PointRule) MinimalPoints(hand *Hand) {
	// 實現判斷平胡（自摸）的邏輯
}

func (pr *PointRule) PungHand(hand *Hand) bool {

	// 實現判斷碰碰胡的邏輯

	if len(hand.Straight) > 0 || len(hand.Kong.Concealed) > 0 || len(hand.Kong.Open) > 0 {
		return false
	}

	tiles := append(hand.Tiles, hand.Draw...)
	results := CountByTiles(tiles)

	foundEyes := false
	for _, r := range results {

		if r == 2 {

			// only one pair of eyes
			if foundEyes {
				return false
			}

			foundEyes = true
			continue
		}

		if r != 3 {
			return false
		}
	}

	return true
}

func (pr *PointRule) HalfFlush(hand *Hand) bool {

	// 實現判斷混一色的邏輯

	tiles := append(hand.Tiles, hand.Draw...)
	tiles = append(tiles, hand.Triplet...)
	tiles = append(tiles, hand.Kong.Open...)
	tiles = append(tiles, hand.Kong.Concealed...)

	for _, s := range hand.Straight {
		tiles = append(tiles, s...)
	}

	results := CountBySuits(tiles)

	// Dragon, Winds and one suit
	if len(results) > 3 || len(results) == 1 {
		return false
	}

	var foundSuit TileSuit
	for suit, _ := range results {
		if suit == TileSuitDragon || suit == TileSuitWind {
			continue
		}

		if len(foundSuit) != 0 {
			return false
		}

		foundSuit = suit
	}

	return true
}

func (pr *PointRule) FullFlush(hand *Hand) bool {

	// 實現判斷清一色的邏輯

	tiles := append(hand.Tiles, hand.Draw...)
	tiles = append(tiles, hand.Triplet...)
	tiles = append(tiles, hand.Kong.Open...)
	tiles = append(tiles, hand.Kong.Concealed...)

	for _, s := range hand.Straight {
		tiles = append(tiles, s...)
	}

	results := CountBySuits(tiles)
	if len(results) != 1 {
		return false
	}

	for suit, _ := range results {
		// No dragon and wind
		if suit == TileSuitDragon || suit == TileSuitWind {
			return false
		}
	}

	return true
}

func (pr *PointRule) LittleThreeDragons(hand *Hand) {
	// 實現判斷小三元的邏輯
}

func (pr *PointRule) AllHonorsHand(hand *Hand) bool {

	// 實現判斷字一色的邏輯

	tiles := append(hand.Tiles, hand.Draw...)
	tiles = append(tiles, hand.Triplet...)
	tiles = append(tiles, hand.Kong.Open...)
	tiles = append(tiles, hand.Kong.Concealed...)

	results := CountBySuits(tiles)
	for suit, _ := range results {
		if suit != TileSuitDragon && suit != TileSuitWind {
			return false
		}
	}

	return true
}

func (pr *PointRule) BigThreeDragons(hand *Hand) bool {

	// 實現判斷大三元的邏輯

	tiles := append(hand.Tiles, hand.Draw...)

	for _, t := range hand.Triplet {
		tiles = append(tiles, t, t, t)
	}

	results := CountByTiles(tiles)

	// 中
	count, ok := results["D1"]
	if !ok || count != 3 {
		return false
	}

	// 發
	count, ok = results["D2"]
	if !ok || count != 3 {
		return false
	}

	// 白
	count, ok = results["D3"]
	if !ok || count != 3 {
		return false
	}

	return true
}

func (pr *PointRule) FourConcealedPungs(hand *Hand) {

	// 實現判斷四暗刻的邏輯
}

func (pr *PointRule) SmallFourWinds(hand *Hand) {
	// 實現判斷小四喜的邏輯
}

func (pr *PointRule) BigFourWinds(hand *Hand) bool {

	// 實現判斷大四喜的邏輯

	tiles := append(hand.Tiles, hand.Draw...)

	for _, t := range hand.Triplet {
		tiles = append(tiles, t, t, t)
	}

	results := CountByTiles(tiles)

	// 東
	count, ok := results["I1"]
	if !ok || count != 3 {
		return false
	}

	// 南
	count, ok = results["I2"]
	if !ok || count != 3 {
		return false
	}

	// 西
	count, ok = results["I3"]
	if !ok || count != 3 {
		return false
	}

	// 北
	count, ok = results["I3"]
	if !ok || count != 3 {
		return false
	}

	return true
}

func (pr *PointRule) HeavenlyHand(hand *Hand) {
	// 實現判斷天胡的邏輯
}

func (pr *PointRule) EarthlyHand(hand *Hand) {
	// 實現判斷地胡的邏輯
}

func (pr *PointRule) FlowerTiles(hand *Hand) {
	// 實現判斷花牌的邏輯
}

func (pr *PointRule) AfterAKong(hand *Hand) {
	// 實現判斷槓上開花的邏輯
}

func (pr *PointRule) LastTileDraw(hand *Hand) {
	// 實現判斷海底撈月的邏輯
}

func (pr *PointRule) RobbingTheKong(hand *Hand) {
	// 實現判斷搶槓胡的邏輯
}

func (pr *PointRule) KongOnDiscard(hand *Hand) {
	// 實現判斷杠上炮的邏輯
}

func (pr *PointRule) SingleWait(hand *Hand) {
	// 實現判斷獨聽的邏輯
}

func (pr *PointRule) SelfDrawn(hand *Hand) {
	// 實現判斷自摸加底的邏輯
}

func (pr *PointRule) MeldedKong(hand *Hand) {
	// 實現判斷明槓的邏輯
}

func (pr *PointRule) ConcealedKong(hand *Hand) {
	// 實現判斷暗槓的邏輯
}

func (pr *PointRule) ConcealedHand(hand *Hand) {
	// 實現判斷門前清的邏輯
}
