package data

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func Init(driverName, dataSourceName string) error {
	var err error
	db, err = sql.Open(driverName, dataSourceName)
	if err != nil {
		return err
	}

	return nil
}

func Quit() error {
	return db.Close()
}
