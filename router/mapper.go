package router

import (
	"github.com/alainmucyo/crud/entity"
)

func toResponseModel(entity *entity.Book) *BookResponse {
	return &BookResponse{
		ID:          entity.ID,
		Title:       entity.Title,
		Description: entity.Description,
		Isbn:        entity.Isbn,
	}
}
