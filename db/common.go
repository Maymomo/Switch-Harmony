package db

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
	PriceCNY        float64
	SaleRate        int
	LanguageTag     []int
	HasSolidEdition bool
	HasDemo         bool
	IsExclusive     bool
}

type GameInfo struct {
	Summary GameSummary
	Detail  GameDetail
}

type GameSummaryJson struct {
	LanguageTag []int
}
type GameSummaryDB struct {
	ID              int
	NameCN          string
	NameEN          string
	ImgUrl          string
	Region          string
	PriceCNY        float64
	SaleRate        int
	HasSolidEdition bool
	HasDemo         bool
	IsExclusive     bool
	Json            string `gorm:"type:TEXT"`
}

type GameDetailJson struct {
	ImgUrl      []string
	Price       []Price
	LowestPrice Price
	LanguageTag []int
	GameScore   []string
	GameTypeTag []string
}
type GameDetailDB struct {
	ID              int
	GameNameCN      string
	GameNameEN      string
	GameSize        string
	HasSolidEdition bool
	HasDemo         bool
	IsExclusive     bool
	HasChinese      bool
	GamePlayers     int
	Description     string `gorm:"type:TEXT"`
	Json            string `gorm:"type:TEXT"`
}

type GameSummaryDBTemp GameSummaryDB
type GameDetailDBTemp GameDetailDB

const (
	DetailTable      = "game_detail_dbs"
	DetailTempTable  = "game_detail_db_temps"
	SummaryTable     = "game_summary_dbs"
	SummaryTempTable = "game_summary_db_temps"
)

const (
	NullStr = ""
)