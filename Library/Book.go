package Library

import "fmt"

type Book struct {
	BasePublication
	size        int
	author      string
	authors     []string
	isbn        string
	genre       string
	edition     int
	series      string
	seriesIndex int
}

func NewBook(id, title, author string, pageCount int) Book {
	return Book{
		BasePublication: BasePublication{
			id:              id,
			title:           title,
			publisher:       "Демо-библиотека",
			publicationDate: "2026-05-05",
			language:        "ru",
			pageCount:       pageCount,
			format:          "hardcover",
		},
		author:  author,
		authors: []string{author},
		genre:   "Художественная литература",
		edition: 1,
	}
}

func (b Book) Summary() string {
	return fmt.Sprintf("%s — %s", b.title, b.author)
}
