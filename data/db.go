package data

import (
	"database/sql"
	"github.com/spf13/viper"
	"log"
)

var mysql *MyDB

type MyDB struct {
	DB *sql.DB
}

func NewMyDB(vip *viper.Viper) *MyDB {
	if mysql == nil {

		dsn := vip.GetString("dsn")
		db, err := sql.Open("mysql", dsn)
		if err != nil {
			log.Fatal(err)
		}
		mysql = &MyDB{
			DB: db,
		}
	}
	return mysql
}
