package db

import (
	"fmt"
	"testing"
)

func TestGetSummaryByPage(t *testing.T) {
	info, page := GetSummaryByPage(1, 10)
	fmt.Println(page)
	fmt.Println(info)
}
