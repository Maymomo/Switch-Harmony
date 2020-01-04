package task

import (
	"github.com/Maymomo/Switch-Harmony/db"
	"log"
)

const maxPageCount = 300
const maxRetryCount = 5

func Work() {
	now := 1
	for now <= maxPageCount {
		tryCount := 0
	retry:
		log.Printf("[TASK] CrawlOvewrPage %d", now)
		pages, err := CrawlOverPage(now)
		if err != nil {
			v, b := err.(StatusCodeError)
			if b && tryCount < maxRetryCount {
				tryCount++
				log.Printf("[TASK] CrawlOvewrPage %d error. Status code is %d. retry it.", now, v.code)
				goto retry
			} else {
				log.Fatal(err)
			}
		}
		if len(pages) == 0 {
			break
		}
		for _, x := range pages {
			x.Summary.InsertToDB()
			x.Detail.InsertToDB()
		}
		now++
	}
	log.Printf("[TASK] swap table begin")
	db.SwapTable(db.DetailTable, db.DetailTempTable)
	db.SwapTable(db.SummaryTempTable, db.SummaryTable)
	log.Printf("[TASK] swap table end")
}
