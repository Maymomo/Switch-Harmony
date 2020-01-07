package db

import (
	"fmt"
	"testing"
)

func TestGetSummaryByPage(t *testing.T) {
	info, page := GetSummaryByPage(NullStr,1, 10)
	fmt.Println(page)
	fmt.Println(info)
}
func TestGetSummaryByPage2(t *testing.T) {
	// test search
	info, page := GetSummaryByPage("塞尔达",0,10)
	fmt.Println(page)
	fmt.Println(info[0])
}

