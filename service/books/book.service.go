package books

import (
	domain "github.com/alainmucyo/crud/repository/books"
	"github.com/alainmucyo/crud/entity"
)

type BookService interface {
	StoreBook(book *entity.Book) (*entity.Book, error)
	ListBooks() ([]entity.Book, error)
}

type Service struct{}

func NewBookService() BookService {
	return &Service{}
}

func (s *Service) StoreBook(book *entity.Book) (*entity.Book, error) {
	return domain.NewBookRepository().CreateBook(book)
}

func (s *Service) ListBooks() ([]entity.Book, error) {
	return domain.NewBookRepository().ListBooks()
}
