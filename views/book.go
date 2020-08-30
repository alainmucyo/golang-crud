package views

type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Isbn   string `json:"isbn"`
	Description string `json:"description"`
}