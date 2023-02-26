package models

// Book struct
type Book struct {
	Id          int    `json:"id"`
	BookName    string `json:"book_name"`
	Author      string `json:"author"`
	Description string `json:"description"`
}
