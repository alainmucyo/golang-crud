package model

import (
	"fmt"
	"github.com/alainmucyo/crud/views"
)

func CreateBook(title, isbn, description string) error {
	insertQ, err := con.Query("INSERT INTO books(title,isbn,description) VALUES(?,?,?)", title, isbn, description)
	defer insertQ.Close()
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func GetBooks() ([]views.Book, error) {
	rows, err := con.Query("SELECT * FROM books")
	defer rows.Close()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Println("Rows: ",rows)
	 books := []views.Book{}
	for rows.Next() {
		data := views.Book{}
		err :=rows.Scan(&data.ID,&data.Title,&data.Isbn,&data.Description)
		if err != nil {
			fmt.Println("Error",err)
		}
		fmt.Println("data",data)
		books = append(books, data)
	}
	return books, nil

}
