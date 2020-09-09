package database

import (
	"database/sql"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var con *sql.DB
func Connect() *sql.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Can't load .env")
	}
	DB_PASSWORD:= os.Getenv("DB_PASSWORD")
	DB_USERNAME := os.Getenv("DB_USERNAME")
	DB_DATABASE := os.Getenv("DB_DATABASE")
	DB_TYPE := os.Getenv("DB_TYPE")
	db,err := sql.Open(DB_TYPE,DB_USERNAME+":"+DB_PASSWORD+"@/"+DB_DATABASE)
	if err !=nil {
		os.Exit(2)
	}
	con = db
	return db
}

func Conn() *sql.DB {
	return con
}
