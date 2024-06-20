package structs

import "time"

type Book struct {
	Isbn        int64     `json:"isbn"`
	Title       string    `json:"title"`
	Subtitle    string    `json:"subtitle"`
	Author      string    `json:"author"`
	Published   time.Time `json:"published"`
	Publisher   string    `json:"publisher"`
	Pages       int64     `json:"pages"`
	Description string    `json:"description"`
	Website     string    `json:"website"`
}

func NewBook(isbn int64, title, subtitle, author string, published time.Time, publisher string, pages int64, description, website string) *Book {
	// TODO: Validate Book Fields BEFORE Creating a new Book

	return &Book{
		Isbn:        isbn,
		Title:       title,
		Subtitle:    subtitle,
		Author:      author,
		Published:   published,
		Publisher:   publisher,
		Pages:       pages,
		Description: description,
		Website:     website,
	}
}

func MakeBook(isbn int64, title, subtitle, author string, published time.Time, publisher string, pages int64, description, website string) Book {
	// TODO: Validate Book Fields BEFORE Creating a new Book

	return Book{
		Isbn:        isbn,
		Title:       title,
		Subtitle:    subtitle,
		Author:      author,
		Published:   published,
		Publisher:   publisher,
		Pages:       pages,
		Description: description,
		Website:     website,
	}
}
