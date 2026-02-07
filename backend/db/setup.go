package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func Open(filename string) (database Database, err error) {
	database.sqldb, err = sql.Open("sqlite3", filename)
	if err != nil {
		return Database{}, err
	}

	_, err = database.sqldb.Exec(sqlInitTables)
	if err != nil {
		return Database{}, err
	}

	return database, nil
}
