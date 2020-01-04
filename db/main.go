package db

import (
	"fmt"
	"github.com/Maymomo/Switch-Harmony/common"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

func DataBaseInit() {
	config := common.GetConfig().MysqlConfig
	argsStr := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=true&loc=Local",
		config.User, config.Password, config.Address, config.Port, config.DB)
	db, err := gorm.Open("mysql", argsStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	db.DB().SetMaxOpenConns(10)
	db.AutoMigrate(&GameDetailDBTemp{})
	db.AutoMigrate(&GameSummaryDBTemp{})
	db.AutoMigrate(&GameDetailDB{})
	db.AutoMigrate(&GameSummaryDB{})
}

func dbConnection() (*gorm.DB, error) {
	config := common.GetConfig().MysqlConfig
	argsStr := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=true&loc=Local",
		config.User, config.Password, config.Address, config.Port, config.DB)
	return gorm.Open("mysql", argsStr)
}
