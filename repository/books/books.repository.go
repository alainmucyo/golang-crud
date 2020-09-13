package books

import (
	"database/sql"
	"fmt"
	"github.com/alainmucyo/crud/entity"
)

type BookRepository interface {
	CreateBook(book *entity.Book) (*entity.Book, error)
	ShowBook(Id int) (*entity.Book, error)
	ListBooks() ([]entity.Book, error)
	DeleteBook(Id int) (string, error)
}

type repository struct {
	*sql.DB
}
func NewBookRepository(db *sql.DB) BookRepository{
	return &repository{db}
}
func (rep *repository) ShowBook(Id int) (*entity.Book, error) {
	return nil,nil
}


func (rep *repository) DeleteBook(Id int) (string, error) {
	panic("implement me")
}

func (rep *repository) CreateBook(book *entity.Book) (*entity.Book, error) {
	con := rep.DB
	insertQ, err := con.Query("INSERT INTO books(title,isbn,description) VALUES(?,?,?)", book.Title, book.Isbn, book.Description)
	defer insertQ.Close()
	if err != nil {
		fmt.Println(err)
		return nil,err
	}

	return nil, nil
}
func (rep *repository) ListBooks() ([]entity.Book, error) {
	con := rep.DB
	rows, err := con.Query("SELECT * FROM books")
	defer rows.Close()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	selectedBooks := []entity.Book{}
	for rows.Next() {
		data := entity.Book{}
		err := rows.Scan(&data.ID, &data.Title, &data.Isbn, &data.Description)
		if err != nil {
			fmt.Println("Error", err)
		}
		fmt.Println("data", data)
		selectedBooks = append(selectedBooks, data)
	}
	return selectedBooks, nil
}

