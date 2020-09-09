package main

import (
	"github.com/alainmucyo/crud/data/database"
	"github.com/alainmucyo/crud/routes"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

func main() {
	//fmt.Println("Hello")
	db := database.Connect()
	defer db.Close()
	println("Server started....")
	log.Fatal(http.ListenAndServe(":3000", routes.Register()))
}
