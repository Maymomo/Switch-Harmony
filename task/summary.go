package task

const PageURL = "http://www.eshop-switch.com/game/queryGame"

type QueryGameSummary struct {
	ID          int     `json:"ID"`
	GAME        string  `json:"GAME"`
	GAME_EN     string  `json:"GAME_EN"`
	SALE        int     `json:"SALE"`
	IMAGE_M     string  `json:"IMAGE_M"`
	PRICE       float64 `json:"PRICE"`
	REGION_NAME string  `json:"REGION_NAME"`
	TAGS        string  `json:"TAGS"`
}

func CrawlOverPage(number int) {

}
