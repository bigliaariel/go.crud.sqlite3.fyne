package config

import (
	"database/sql"
)

func GetDB() (db *sql.DB, err error) {
	db, err = sql.Open("sqlite3", "../db/data.db")
	return
}
