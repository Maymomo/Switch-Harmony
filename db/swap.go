package db

import (
	"fmt"
	"log"
)

const tmpName = "asdhiaosdjhasoingqe"

func SwapTable(table1, table2 string) {
	db, err := dbConnection()
	if err != nil {
		log.Fatal(err)
	}
	tx := db.Begin()
	if err = tx.Exec(fmt.Sprintf("RENAME table %s to %s, %s to %s, %s to %s",
		table1, tmpName, table2, table1, tmpName, table2)).Error; err != nil {
		tx.Rollback()
	}
	tx.Commit()
}
