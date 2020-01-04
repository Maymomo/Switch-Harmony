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
	db.Table(DetailTempTable).Create(r.toDB())
}

func QueryDetailByID(id int) *GameDetail {
	db, err := dbConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	var tmp GameDetailDB
	db.Table(DetailTable).Where("id = ?", strconv.Itoa(id)).First(&tmp)
	return tmp.toInfo()
}
