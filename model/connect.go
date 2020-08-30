package model

import (
	"database/sql"
	"os"
)
var con *sql.DB
func Connect() *sql.DB {
	db,err := sql.Open("mysql","root:password@/go_learn")
	if err !=nil {
		os.Exit(2)
	}
	con = db
	return db
}
