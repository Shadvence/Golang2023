package data

import (
	"Assignments/internal/validator"
	"time"
)

type Book struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"-"`
	Title     string    `json:"title"`
	Year      int32     `json:"year,omitempty"`
	Author    string    `json:"author,omitempty"`
	Genres    []string  `json:"genres,omitempty"`
	Version   int32     `json:"version"`
}

func ValidateBook(v *validator.Validator, book *Book) {
	v.Check(book.Title != "", "title", "must be provided")
	v.Check(len(book.Title) <= 500, "title", "must not be more than 500 bytes long")

	v.Check(book.Year != 0, "year", "must be provided")
	v.Check(book.Year >= 1888, "year", "must be greater than 1888")
	v.Check(book.Year <= int32(time.Now().Year()), "year", "must not be in the future")

	v.Check(book.Genres != nil, "genres", "must be provided")
	v.Check(len(book.Genres) >= 1, "genres", "must contain at least 1 genre")
	v.Check(len(book.Genres) <= 5, "genres", "must not contain more than 5 genres")
	v.Check(validator.Unique(book.Genres), "genres", "must not contain duplicate values")
}
