package main

import (
	"github.com/alainmucyo/crud/model"
	"github.com/alainmucyo/crud/routes"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

func main() {
	db := model.Connect()
	defer db.Close()
	println("Server started....")
	log.Fatal(http.ListenAndServe(":8080", routes.Register()))
}
