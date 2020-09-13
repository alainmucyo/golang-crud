package books

import (
	"encoding/json"
	"fmt"
	"github.com/alainmucyo/crud/entity"
	"github.com/alainmucyo/crud/service/books"
	"net/http"
)

type BookController interface {
	GetBooks(writer http.ResponseWriter, request *http.Request)
	CreateBook(writer http.ResponseWriter, request *http.Request)
}

type Controller struct {
	books.BookService
}


func NewBooksController(service books.BookService) BookController {
	return &Controller{service}
}

func (c *Controller) GetBooks(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	data, err := c.BookService.ListBooks()
	if err != nil {
		writer.Write([]byte("Error Occurred"))
		return
	}
	booksResponse := make([]BookResponse, len(data))
	for i, elem := range data {
		booksResponse[i] = toBooksResponse(&elem)
	}
	booksList := toBooksList(&booksResponse)
	json.NewEncoder(writer).Encode(booksList)
}

func (c *Controller) CreateBook(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var book entity.Book
	_ = json.NewDecoder(request.Body).Decode(&book)
	_, err := c.BookService.StoreBook(&book)
	if err != nil {
		writer.Write([]byte("Some errors"))
		fmt.Println(err)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	json.NewEncoder(writer).Encode(book)
}

/*func Show(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	for _, book := range Books {
		if book.ID == params["id"] {
			json.NewEncoder(writer).Encode(book)
			return
		}
	}
}*/

/*func DeleteBook(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	for index, book := range Books {
		if book.ID == params["id"] {
			Books = append(Books[:index], Books[index+1:]...)
		}
	}
	json.NewEncoder(writer).Encode("Ok")
}*/

/*func UpdateBook(writer http.ResponseWriter, request *http.Request) {
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
}*/
