package task

import (
	"fmt"
	"github.com/Maymomo/Switch-Harmony/db"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func GetGameDetailInfo(url string) (*db.GameDetail, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, FormatStatusCodeError(resp.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}
	var detail db.GameDetail
	detail.GameNameCN = doc.Find("div.game-name").Text()
	detail.GameNameEN = doc.Find("div.game-en").Text()
	//size := doc.Find("div.game-size")
	//lowest := doc.Find("div.first-game-detail")
	detail.Description = doc.Find("div.game-description").Text()
	imgUrl := doc.Find("div.carousel-inner").Find("img.game-image-s").Nodes
	for _, x := range imgUrl {
		detail.ImgUrl = append(detail.ImgUrl, x.Attr[1].Val)
	}
	tag1 := doc.Find("div.game-tag1-view").Find("span").First()
	for len(tag1.Nodes) != 0 {
		tagType, val := db.Str2Tag(tag1.Text())
		switch tagType {
		case db.LanguageType:
			detail.LanguageTag = append(detail.LanguageTag, val)
			detail.HasChinese = true
		case db.HasSolidEdition:
			detail.HasSolidEdition = true
		case db.Exclusive:
			detail.IsExclusive = true
		case db.HasDemo:
			detail.HasDemo = true
		default:
			log.Fatal(fmt.Errorf("tag %s parse error", tag1.Text()))
		}
		tag1 = tag1.Next()
	}
	score := doc.Find("div.game-ign-mc").Find("span").First()
	for {
		k, b := score.Attr("class")
		if !b || k == "score-button" {
			break
		}
		detail.GameScore = append(detail.GameScore, strings.TrimSpace(score.Text()))
		score = score.Next()
	}
	gameType := strings.Split(doc.Find("div.game-tag-view").Find("span").Text(), "/")
	for _, str := range gameType {
		detail.GameTypeTag = append(detail.GameTypeTag, strings.TrimSpace(str))
	}
	tmp := strings.Split(strings.TrimSpace(doc.Find("div.game-size").Text()), "|")
	gameSize := strings.TrimSpace(strings.Split(tmp[0], "：")[1])
	detail.GameSize = gameSize
	gamePlayer := strings.TrimSpace(strings.Split(tmp[1], "：")[1])
	detail.GamePlayers, _ = strconv.Atoi(gamePlayer[:len(gamePlayer)-len("人")])

	// parse Price
	priceInfo := doc.Find("div.col-md-12").Find("div.single-region-view").First()
	for {
		v, b := priceInfo.Attr("class")
		if len(priceInfo.Nodes) == 0 || !b || v == "more-region-info" {
			break
		}
		var language []string
		buffer := strings.Split(strings.TrimSpace(priceInfo.Find("div.game-language").Find("span").Text()), ":")
		if len(buffer) > 1 {
			language = strings.Split(buffer[1], "，")
		}
		price := db.Price{
			Region:        priceInfo.Find("div.region").Text(),
			HasChinese:    len(priceInfo.Find("div.chinese-view").Find("span.chinese").Nodes) != 0,
			IsLowestPrice: len(priceInfo.Find("div.chinese-view").Find("div.shidi").Nodes) != 0,
			GameLanguage:  language,
			RegionComment: strings.TrimSpace(priceInfo.Find("div.single-region-third-view").Find(".game-region-comment").Text()),
		}
		tmp := priceInfo.Find(".single-region-third-view").Find(".release-date.col-md-3").Find("span").Text()
		if len(tmp) != 0 {
			tmp = strings.Split(tmp, "：")[1]
			lowest, err := db.ParsePrice(tmp)
			if err != nil {
				log.Fatal(err)
			}
			price.HistoryLowestPrice = lowest
		}
		tmp = strings.TrimSpace(priceInfo.Find(".single-region-second-view").Find("div.release-date").Text())
		price.ReleaseDate = tmp[:len(tmp)-len("发售")]

		tmp = strings.TrimSpace(priceInfo.Find("span.sale").Text())
		if len(tmp) != 0 {
			tmp = tmp[:len(tmp)-len("折")]
			saleRate, err := strconv.ParseFloat(tmp, 10)
			if err != nil {
				log.Fatal(err)
			}
			price.SaleRate = saleRate
		}
		tmp = priceInfo.Find("div.end-date").Text()
		if len(tmp) != 0 {
			_, err = fmt.Sscanf(tmp, "剩余%d天", &price.EndDate)
			if err != nil {
				log.Fatal(err)
			}
			price.IsOnSale = true
		} else {
			price.IsOnSale = false
		}
		tmp = priceInfo.Find("span.price-cny").Text()
		priceCNY, err := db.ParsePrice(tmp)
		if err != nil {
			log.Fatal(err)
		}
		price.PriceCNY = priceCNY
		tmp = priceInfo.Find("div.price").Text()
		priceLocal, err := db.ParsePrice(tmp)
		if err != nil {
			log.Fatal(err)
		}
		price.PriceLocal = priceLocal
		detail.Price = append(detail.Price, price)

		priceInfo = priceInfo.Next()
	}
	return &detail, err
}
