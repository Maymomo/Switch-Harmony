package task

import (
	"fmt"
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

func StatusCodeError(status int) error {
	return fmt.Errorf("status code is %d", status)
}

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
