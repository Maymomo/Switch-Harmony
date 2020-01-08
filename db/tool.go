package db

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"
)

const (
	ALL = iota
	HK
	UK
	EU
	USA
	JP
)

const (
	LanguageType = iota
	HasSolidEdition
	Exclusive
	HasDemo
)

func Str2Tag(tag string) (int, int) {
	tag = strings.TrimSpace(tag)
	switch tag {
	case "全区中文":
		return LanguageType, ALL
	case "港区中文":
		return LanguageType, HK
	case "美区中文":
		return LanguageType, USA
	case "日区中文":
		return LanguageType, JP
	case "试玩":
		return HasDemo, 0
	case "实体":
		return HasSolidEdition, 0
	case "独占":
		return Exclusive, 0
	default:
		return 0, 0
	}
}

func ParsePrice(str string) (float64, error) {
	start := 0
	end := len(str)
	for {
		if start >= len(str) {
			return 0.0, fmt.Errorf("parse price %s error", str)
		}
		if !(str[start] >= '0' && str[start] <= '9') {
			start++
		} else {
			break
		}
	}
	for {
		if end <= start {
			return 0.0, fmt.Errorf("parse Price %s error", str)
		}
		if !(str[end-1] >= '0' && str[end-1] <= '9') {
			end--
		} else {
			break
		}
	}
	return strconv.ParseFloat(str[start:end], 10)
}

func Tag2Str(tagType, tag int) string {
	return ""
}

func (r *GameDetailDB) toInfo() *GameDetail {
	if r == nil {
		return nil
	}
	str := r.Json
	var buf GameDetailJson
	err := json.Unmarshal([]byte(str), &buf)
	if err != nil {
		log.Printf("[ERROR] Toinfo Parse json error %e\n", err)
		return nil
	}
	return &GameDetail{
		ID:              r.ID,
		GameNameCN:      r.GameNameCN,
		GameNameEN:      r.GameNameEN,
		GameSize:        r.GameSize,
		ImgUrl:          buf.ImgUrl,
		Price:           buf.Price,
		LowestPrice:     buf.LowestPrice,
		LanguageTag:     buf.LanguageTag,
		HasSolidEdition: r.HasSolidEdition,
		HasDemo:         r.HasDemo,
		IsExclusive:     r.IsExclusive,
		GameTypeTag:     buf.GameTypeTag,
		HasChinese:      r.HasChinese,
		GamePlayers:     r.GamePlayers,
		Description:     r.Description,
		GameScore:       buf.GameScore,
	}
}

func (r *GameSummaryDB) toInfo() *GameSummary {
	if r == nil {
		return nil
	}
	str := r.Json
	var buf GameSummaryJson
	if err := json.Unmarshal([]byte(str), &buf); err != nil {
		log.Printf("[ERROR] Toinfo Parse json error %e\n", err)
		return nil
	}
	return &GameSummary{
		ID:              r.ID,
		NameCN:          r.NameCN,
		NameEN:          r.NameEN,
		ImgUrl:          r.ImgUrl,
		Region:          r.Region,
		PriceCNY:        r.PriceCNY,
		SaleRate:        r.SaleRate,
		LanguageTag:     buf.LanguageTag,
		HasSolidEdition: r.HasSolidEdition,
		HasDemo:         r.HasDemo,
		IsExclusive:     r.IsExclusive,
	}
}

func (r *GameSummary) toDB() *GameSummaryDB {
	if r == nil {
		return nil
	}
	ret := GameSummaryDB{
		ID:              r.ID,
		NameCN:          r.NameCN,
		NameEN:          r.NameEN,
		ImgUrl:          r.ImgUrl,
		Region:          r.Region,
		PriceCNY:        r.PriceCNY,
		SaleRate:        r.SaleRate,
		HasSolidEdition: r.HasSolidEdition,
		HasDemo:         r.HasDemo,
		IsExclusive:     r.IsExclusive,
	}
	tmp := GameSummaryJson{
		LanguageTag: r.LanguageTag,
	}
	buffer, err := json.Marshal(tmp)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	ret.Json = string(buffer)
	return &ret
}

func (r *GameDetail) toDB() *GameDetailDB {
	if r == nil {
		return nil
	}
	ret := GameDetailDB{
		ID:              r.ID,
		GameNameCN:      r.GameNameCN,
		GameNameEN:      r.GameNameEN,
		GameSize:        r.GameSize,
		HasSolidEdition: r.HasSolidEdition,
		HasDemo:         r.HasDemo,
		IsExclusive:     r.IsExclusive,
		HasChinese:      r.HasChinese,
		GamePlayers:     r.GamePlayers,
		Description:     r.Description,
	}
	tmp := GameDetailJson{
		ImgUrl:      r.ImgUrl,
		Price:       r.Price,
		LowestPrice: r.LowestPrice,
		LanguageTag: r.LanguageTag,
		GameScore:   r.GameScore,
		GameTypeTag: r.GameTypeTag,
	}
	buffer, err := json.Marshal(tmp)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	ret.Json = string(buffer)
	return &ret
}
