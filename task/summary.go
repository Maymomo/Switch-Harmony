package task

import (
	"encoding/json"
	"fmt"
	"github.com/Maymomo/Switch-Harmony/db"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"
)

const PageURL = "http://www.eshop-switch.com/game/queryGame"
const DetailURL = "http://www.eshop-switch.com/game/"

type QueryGameSummary struct {
	ID          int     `json:"ID"`
	GAME        string  `json:"GAME"`
	GAME_EN     string  `json:"GAME_EN"`
	SALE        string  `json:"SALE"`
	IMAGE_M     string  `json:"IMAGE_M"`
	PRICE       float64 `json:"PRICE"`
	REGION_NAME string  `json:"REGION_NAME"`
	TAGS        string  `json:"TAGS"`
}
type QueryResp struct {
	CurrentPage    int                `json:"currentPage"`
	List           []QueryGameSummary `json:"list"`
	PageSize       int                `json:"pageSize"`
	TotalPageCount int                `json:"totalPageCount"`
}

func crawlGameDetail(x QueryGameSummary, save *[]*db.GameInfo, group *sync.WaitGroup) {
	defer group.Done()
	u := DetailURL + strconv.Itoa(x.ID) + ".html"
	count := 0
retry:
	detail, err := GetGameDetailInfo(u)
	detail.ID = x.ID
	if err != nil {
		count++
		log.Printf("[TASK] Get game %d detail info error. Error is %s. Retry it.", x.ID, err.Error())
		goto retry
	}
	sale := 100
	if x.SALE != "" {
		sale, err = strconv.Atoi(x.SALE)
		if err != nil {
			fmt.Printf("parse game %d sale rate err\n", x.ID)
			log.Fatal(err)
		}
	}
	summary := db.GameSummary{
		ID:       x.ID,
		NameCN:   x.GAME,
		NameEN:   x.GAME_EN,
		ImgUrl:   x.IMAGE_M,
		Region:   x.REGION_NAME,
		PriceCNY: x.PRICE,
		SaleRate: sale,
	}
	tags := strings.Split(x.TAGS, " ")
	for _, x := range tags {
		t, v := db.Str2Tag(x)
		switch t {
		case db.LanguageType:
			summary.LanguageTag = append(summary.LanguageTag, v)
		case db.Exclusive:
			summary.IsExclusive = true
		case db.HasDemo:
			summary.HasDemo = true
		case db.HasSolidEdition:
			summary.HasSolidEdition = true
		}
	}
	*save = append(*save, &db.GameInfo{
		Detail:  *detail,
		Summary: summary,
	})
	return
}

func CrawlOverPage(number int) ([]*db.GameInfo, error) {
	values := url.Values{}
	values.Set("current_page", strconv.Itoa(number))
	values.Set("page_size", "24")
	values.Set("order_by", "0")
	resp, err := http.PostForm(PageURL, values)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, FormatStatusCodeError(resp.StatusCode)
	}
	buffer, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var queryResp QueryResp
	err = json.Unmarshal(buffer, &queryResp)
	if err != nil {
		return nil, err
	}
	var ret []*db.GameInfo
	waitGroup := sync.WaitGroup{}
	for _, x := range queryResp.List {
		waitGroup.Add(1)
		go crawlGameDetail(x, &ret, &waitGroup)
	}
	waitGroup.Wait()
	return ret, nil
}
