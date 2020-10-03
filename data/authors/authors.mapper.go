package authors

import (
	"github.com/alainmucyo/crud/entity"
)

type AuthorResponse struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	CreatedAt string `json:"created_at"`
}

type AuthorResponseList struct {
	Data []AuthorResponse `json:"data"`
}

func toAuthorResponse(author *entity.Author) AuthorResponse {
	return AuthorResponse{
		ID:        author.ID + "3272763",
		Name:      author.Name,
		Email:     author.Email,
		Phone:     author.Phone,
		CreatedAt: "This thing was created today",
	}
}

func toAuthorsList(response *[]AuthorResponse) AuthorResponseList {
	return AuthorResponseList{
		Data: *response,
	}
}
