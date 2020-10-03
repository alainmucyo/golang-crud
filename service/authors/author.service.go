package authors

import (
	"github.com/alainmucyo/crud/entity"
	"github.com/alainmucyo/crud/repository/authors"
)

type AuthorService interface {
	CreateAuthor(author *entity.Author) (*entity.Author, error)
	ListAuthor() ([]entity.Author, error)
}

type service struct {
	authors.AuthorRepository
}

func NewAuthorService(repository authors.AuthorRepository) AuthorService {
	return &service{repository}
}