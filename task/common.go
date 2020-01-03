package task

type Price struct {
	Region             string
	HasChinese         bool
	IsLowestPrice      bool
	IsOnSale           bool
	SaleRate           float64
	EndDate            int
	PriceCNY           float64
	PriceLocal         float64
	GameLanguage       []string
	ReleaseDate        string
	RegionComment      string
	HistoryLowestPrice float64
}
type GameDetail struct {
	ID              int
	GameNameCN      string
	GameNameEN      string
	GameSize        string
	ImgUrl          []string
	Price           []Price
	LowestPrice     Price
	LanguageTag     []int
	HasSolidEdition bool
	HasDemo         bool
	IsExclusive     bool
	GameTypeTag     []string
	HasChinese      bool
	GamePlayers     int
	Description     string
	GameScore       []string
}

type GameSummary struct {
	ID              int
	NameCN          string
	NameEN          string
	ImgUrl          string
	Region          string
	PriceCNY        string
	SaleRate        string
	LanguageTag     []int
	HasSolidEdition bool
	HasDemo         bool
	IsExclusive     bool
}

type GameInfo struct {
	Summary GameSummary
	Detail  GameDetail
}
