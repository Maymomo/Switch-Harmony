package task

import (
	"fmt"
	"testing"
)

func TestCrawlOverPage(t *testing.T) {
	ret, err := CrawlOverPage(1)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(ret)
	for _, x := range ret {
		x.Summary.InsertToDB()
		x.Detail.InsertToDB()
	}
}
