package main

import (
	bookController "github.com/alainmucyo/crud/data/books"
	"github.com/alainmucyo/crud/data/database"
	bookRepository "github.com/alainmucyo/crud/repository/books"
	"github.com/alainmucyo/crud/routes"
	bookService "github.com/alainmucyo/crud/service/books"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {
	//fmt.Println("Hello")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Can't load .env")
	}
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	DB_USERNAME := os.Getenv("DB_USERNAME")
	DB_DATABASE := os.Getenv("DB_DATABASE")
	DB_TYPE := os.Getenv("DB_TYPE")
	db := database.Connect(DB_PASSWORD,DB_USERNAME,DB_DATABASE,DB_TYPE)
	booksRep := bookRepository.NewBookRepository(db)
	service := bookService.NewBookService(booksRep)
	controller:=bookController.NewBooksController(service)
	httpHandler := routes.RouteRegister(controller)
	PORT := os.Getenv("PORT")
	URL := os.Getenv("URL")
	defer db.Close()
	println("Server started at " + URL + ":" + PORT)
	log.Fatal(http.ListenAndServe(":"+PORT, httpHandler))
}
