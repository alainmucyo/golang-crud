package books

import (
	"fmt"
	"github.com/alainmucyo/crud/entity"
	domain "github.com/alainmucyo/crud/repository/books"

)

type BookService interface {
	StoreBook(book *entity.Book) (*entity.Book, error)
	ListBooks() ([]entity.Book, error)
}

type Service struct{
	domain.BookRepository
}

func NewBookService(repository domain.BookRepository) BookService {
	return &Service{
		repository,
	}
}

func (s *Service) StoreBook(book *entity.Book) (*entity.Book, error) {
	return s.BookRepository.CreateBook(book)
}

func (s *Service) ListBooks() ([]entity.Book, error) {
	fmt.Println("Listing books")
	return s.BookRepository.ListBooks()
}
