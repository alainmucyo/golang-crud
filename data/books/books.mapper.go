package books

import "github.com/alainmucyo/crud/entity"

type BookResponse struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Isbn   string `json:"isbn"`
	Description string `json:"description"`
}

func toBooksResponse(book *entity.Book) *BookResponse{
	return &BookResponse{
		ID: book.ID,
		Title: book.Title,
		Isbn: book.Isbn,
		Description: book.Description,
	}
}
