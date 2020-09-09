package main

import (
	"github.com/alainmucyo/crud/data/database"
	"github.com/alainmucyo/crud/routes"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {
	//fmt.Println("Hello")
	db := database.Connect()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Can't load .env")
	}
	port:= os.Getenv("PORT")
	url := os.Getenv("URL")
	defer db.Close()
	println("Server started at "+url+":"+port)
	log.Fatal(http.ListenAndServe(":"+port, routes.Register()))
}
