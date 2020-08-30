package controller

import (
	"encoding/json"
	"fmt"
	"github.com/alainmucyo/crud/model"
	"github.com/alainmucyo/crud/views"
	"github.com/gorilla/mux"
	"net/http"
)

var Books []views.Book

func GetBooks(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	data,err :=model.GetBooks()
	if err!=nil {
		writer.Write([]byte("Error Occurred"))
		return
	}
	json.NewEncoder(writer).Encode(data)
}

func ShowBook(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	for _, book := range Books {
		if book.ID == params["id"] {
			json.NewEncoder(writer).Encode(book)
			return
		}
	}
}

func DeleteBook(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	for index, book := range Books {
		if book.ID == params["id"] {
			Books = append(Books[:index], Books[index+1:]...)
		}
	}
	json.NewEncoder(writer).Encode("Ok")
}

func UpdateBook(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	for index, book := range Books {
		if book.ID == params["id"] {
			Books = append(Books[:index], Books[index+1:]...)
		}
	}
	var book views.Book
	_ = json.NewDecoder(request.Body).Decode(&book)
	book.ID = params["id"]
	Books = append(Books, book)
	json.NewEncoder(writer).Encode(book)
}

func StoreBook(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var book views.Book
	_ = json.NewDecoder(request.Body).Decode(&book)
	err := model.CreateBook(book.Title, book.Isbn, book.Description)
	if err != nil {
		writer.Write([]byte("Some errors"))
		fmt.Println(err)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	json.NewEncoder(writer).Encode(book)
}
