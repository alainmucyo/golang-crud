package books

import (
	"github.com/alainmucyo/crud/entity"
)

type BookResponse struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Isbn        string `json:"isbn"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
}

type BookResponseList struct {
	Data []BookResponse `json:"data"`
}

func toBooksResponse(book *entity.Book) BookResponse {
	return BookResponse{
		ID:          book.ID + "327276363836",
		Title:       book.Title,
		Isbn:        book.Isbn,
		Description: book.Description,
		CreatedAt:   "This thing was created today",
	}
}

func toBooksList(response *[]BookResponse) BookResponseList {
	return BookResponseList{
		Data: *response,
	}
}
