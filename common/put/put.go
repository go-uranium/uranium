package put

import "database/sql"

type Putter interface {
	Put(...interface{}) (sql.Result, error)
}

type dbPutter struct {
	db    *sql.DB
	query string
}

func (dbp *dbPutter) Put(args ...interface{}) (sql.Result, error) {
	return dbp.db.Exec(dbp.query, args...)
}

func PutterFromDBExec(db *sql.DB, query string) Putter {
	return &dbPutter{
		db:    db,
		query: query,
	}
}
