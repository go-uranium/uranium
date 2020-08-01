package controllor

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Init(url string) error {
	db, err := sql.Open("mysql", url)
	if err != nil {
		return err
	}
	DB = db
	return nil
}
