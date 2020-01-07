package db

import (
	"log"
	"strconv"
)

func (r *GameDetail) InsertToDB() {
	db, err := dbConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	tmp := r.toDB()
	b := db.Table(DetailTempTable).NewRecord(tmp)
	if !b {
		db.Table(DetailTempTable).Save(tmp)
	} else {
		db.Table(DetailTempTable).Create(tmp)
	}
}

func QueryDetailByID(id int) (*GameDetail, error) {
	db, err := dbConnection()
	if err != nil {
		log.Printf("[ERROR] QueryDetailByID connect db error.")
		return nil, err
	}
	defer db.Close()
	var tmp GameDetailDB
	err = db.Table(DetailTable).Where("id = ?", strconv.Itoa(id)).First(&tmp).Error
	return tmp.toInfo(), err
}
