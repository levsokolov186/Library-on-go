package Library

import "fmt"

type Comics struct {
	BasePublication
	issueNumber     int
	volume          int
	series          string
	writers         []string
	artists         []string
	inkers          []string
	colorists       []string
	letterers       []string
	coverArtists    []string
	genre           string
	ageRating       string
	synopsis        string
	hasVariantCover bool
}

func NewComics(id, title, series string, issueNumber int) Comics {
	return Comics{
		BasePublication: BasePublication{
			id:              id,
			title:           title,
			publisher:       "Демо-библиотека",
			publicationDate: "2026-05-05",
			language:        "ru",
			pageCount:       32,
			format:          "softcover",
		},
		issueNumber: issueNumber,
		volume:      1,
		series:      series,
		writers:     []string{"Редакция"},
		artists:     []string{"Художник"},
		genre:       "Приключения",
		ageRating:   "12+",
	}
}

func (c Comics) Summary() string {
	return fmt.Sprintf("%s — %s #%d", c.title, c.series, c.issueNumber)
}
