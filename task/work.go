package task

import (
	"github.com/Maymomo/Switch-Harmony/db"
	"log"
	"time"
)

const maxPageCount = 300
const maxRetryCount = 5

func needCrawlDirectly() bool {
	_, npage := db.GetSummaryByPage(db.NullStr, 0, 1)
	return npage == 0
}

func BackgroundWork() {
	go func() {
		if needCrawlDirectly() {
			Work()
		}
		for _ = range time.Tick(time.Hour * 2) {
			Work()
		}
	}()
}

func Work() {
	now := 1
	for now <= maxPageCount {
		tryCount := 0
	retry:
		log.Printf("[TASK] CrawlOvewrPage %d", now)
		pages, err := CrawlOverPage(now)
		if err != nil {
			if tryCount < maxRetryCount {
				tryCount++
				log.Printf("[TASK] CrawlOvewrPage %d error. err is %e. retry it.\n", now, err)
				goto retry
			} else {
				log.Printf("[TASK][ERROR] CrawlOverPage error %e. Return and retry after two hours.\n", err)
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
