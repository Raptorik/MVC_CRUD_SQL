package config

import "database/sql"

func GetDB() (db *sql.DB, err error) {
	dbDriver := "mysql"
	dbName := "my-mvc"
	dbUser := "user"
	dbPass := "password"
	db, err = sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	return
}
