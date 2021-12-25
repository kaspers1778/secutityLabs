package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func NewDB(dataSourceName string) *sql.DB {
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		panic(fmt.Errorf("db cannot open :%v", err))
	}
	return db
}

