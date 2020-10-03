package authors

import (
	"encoding/json"
	"github.com/alainmucyo/crud/entity"
	"github.com/alainmucyo/crud/service/authors"
	"net/http"
)

type AuthorController interface {
	Index(writer http.ResponseWriter, request *http.Request)
	Store(writer http.ResponseWriter, request *http.Request)
}

type controller struct {
	authors.AuthorService
}

func NewAuthController(service authors.AuthorService) AuthorController {
	return &controller{service}
}
func (c *controller) Index(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	data, err := c.ListAuthor()
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("Server error"))
		return
	}
	authorResponse := make([]AuthorResponse, len(data))
	for i, author := range data {
		authorResponse[i] = toAuthorResponse(&author)
	}
	authorsList := toAuthorsList(&authorResponse)
	json.NewEncoder(writer).Encode(authorsList)

}

func (c *controller) Store(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var author entity.Author
	_ = json.NewDecoder(request.Body).Decode(&author)
	_, err := c.CreateAuthor(&author)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode("Server errors")
		return
	}
	json.NewEncoder(writer).Encode("Author Created successfully!")
}
