package db

import (
	"github.com/jinzhu/gorm"
)

func GormConnect() *gorm.DB {
	DBMS := "mysql"
	USER := "developer"
	PASS := ""
	PROTOCOL := "tcp(localhost:3306)"
	DBNAME := "dev_test"
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=True&loc=Local"

	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic(err.Error())
	}

	return db
}
