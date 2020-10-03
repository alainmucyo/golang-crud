package authors

import (
	"database/sql"
	"fmt"
	"github.com/alainmucyo/crud/entity"
)

type AuthorRepository interface {
	CreateAuthor(author *entity.Author) (*entity.Author, error)
	ListAuthor() ([]entity.Author, error)
}

type authorRepository struct {
	*sql.DB
}

func NewAuthorityRepository(db *sql.DB) AuthorRepository {
	return &authorRepository{db}
}
func (a *authorRepository) CreateAuthor(author *entity.Author) (*entity.Author, error) {
	con := a.DB
	insertQ, err := con.Query("INSERT INTO authors(name,email,phone) VALUES(?,?,?)", author.Name, author.Email, author.Phone)
	defer insertQ.Close()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return nil, nil
}

func (a *authorRepository) ListAuthor() ([]entity.Author, error) {
	con := a.DB
	rows, err := con.Query("SELECT * FROM authors")
	defer rows.Close()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	selectedAuthors := []entity.Author{}
	for rows.Next() {
		data := entity.Author{}
		err := rows.Scan(&data.ID, &data.Name, &data.Email, &data.Phone)
		if err != nil {
			fmt.Println("Error", err)
		}
		fmt.Println("data", data)
		selectedAuthors = append(selectedAuthors, data)
	}
	return selectedAuthors, nil
}
