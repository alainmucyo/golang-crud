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
	db := database.Connect()
	booksRep := bookRepository.NewBookRepository(db)
	service := bookService.NewBookService(booksRep)
	controller:=bookController.NewBooksController(service)
	httpHandler := routes.RouteRegister(controller)
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Can't load .env")
	}
	port := os.Getenv("PORT")
	url := os.Getenv("URL")
	defer db.Close()
	println("Server started at " + url + ":" + port)
	log.Fatal(http.ListenAndServe(":"+port, httpHandler))
}
