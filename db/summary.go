package db

import (
	"fmt"
	"log"
	"math"
	"strconv"
)

const MaxPagingLimit = 50

func (r *GameSummary) InsertToDB() {
	db, err := dbConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	tmp := r.toDB()
	// check weather exit
	b := db.Table(SummaryTempTable).NewRecord(tmp)
	if !b {
		db.Table(SummaryTempTable).Save(tmp)
	} else {
		db.Table(SummaryTempTable).Create(tmp)
	}
}

func GetSummaryByPage(str string, offset, limit int) (ret []*GameSummary, totalPages int) {
	if offset < 0 || limit < 0 {
		return nil, 0
	}
	if limit > MaxPagingLimit {
		limit = MaxPagingLimit
	}
	db, err := dbConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	var tmp []GameSummaryDB
	if str == NullStr {
		db.Table(SummaryTable).Offset(offset).Limit(limit).Find(&GameSummaryDB{}).Find(&tmp)
		aff := db.Table(SummaryTable).Find(&GameSummaryDB{}).RowsAffected
		totalPages = int(math.Ceil(float64(aff) / float64(limit)))
	} else {
		aff := db.Table(SummaryTable).Offset(offset).Limit(limit).Where("name_cn LIKE ?", fmt.Sprintf("%%%s%%",str)).Find(&tmp).RowsAffected
		totalPages = int(math.Ceil(float64(aff) / float64(limit)))
	}
	for _, v := range tmp {
		ret = append(ret, v.toInfo())
	}

	return
}

func GetSummaryByID(ID int) (*GameSummary, error) {
	db, err := dbConnection()
	if err != nil {
		log.Printf("[ERROR] GetSummaryByID connect db errror.")
		return nil, err
	}
	defer db.Close()
	var ret GameSummary
	err = db.Table(SummaryTable).Where("id = ?", strconv.Itoa(ID)).First(&ret).Error
	return &ret, err
}
