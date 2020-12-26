package database

import (
	"database/sql"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db   *sql.DB
	once sync.Once
)

func GetDatabase() *sql.DB {
	return db
}

func InitDatabase() {
	once.Do(
		func() {
			var err error
			db, err = sql.Open("mysql", "root:123@tcp(localhost:3306)/orderfood")
			if err != nil {
				panic(err)
			}
		})
}
