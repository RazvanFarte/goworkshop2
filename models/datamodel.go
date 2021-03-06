package models

import "fmt"

//BookDto - The DTO used to access books
type BookDto struct {
	UUID        string    `json:"uuid"`
	Title       string    `json:"title"`
	NoPages     int       `json:"noPages"`
	ReleaseDate string    `json:"releaseDate"`
	Author      AuthorDto `json:"author"`
}

//AuthorDto - The DTO used to access authors
type AuthorDto struct {
	UUID      string `json:"uuid"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Birthday  string `json:"birthday"`
	Death     string `json:"death"`
}

func (author AuthorDto) String() string {
	return fmt.Sprintf("AuthorDto{ uuid:%s, author:%s, lastName:%d, birthDay:%s, death:%s}\n", author.UUID,
		author.FirstName, author.LastName, author.Birthday, author.Death)
}

func (book BookDto) String() string {
	return fmt.Sprintf("BookDTO{ uuid:%s, title:%s, noPages:%d, releaseDate:%s, author:%s}\n", book.UUID,
		book.Title, book.NoPages, book.ReleaseDate, book.Author)
}


//Books - the list of available books
var Books []BookDto

// Authors - the list of available authors
var Authors []AuthorDto
