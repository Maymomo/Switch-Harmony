package db

import (
	"log"
	"math"
)

func (r *GameSummary) InsertToDB() {
	db, err := dbConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	db.Table(SummaryTempTable).Create(r.toDB())
}

func GetSummaryByPage(page, size int) (ret []*GameSummary, totalPages int) {
	db, err := dbConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	var tmp []GameSummaryDB
	db.Table(SummaryTable).Offset(page * size).Limit(size).Find(&GameSummaryDB{}).Find(&tmp)
	aff := db.Table(SummaryTable).Find(&GameSummaryDB{}).RowsAffected
	totalPages = int(math.Ceil(float64(aff) / float64(size)))
	for _, v := range tmp {
		ret = append(ret, v.toInfo())
	}
	return
}
