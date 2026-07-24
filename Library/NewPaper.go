package Library

import "fmt"

type Newspaper struct {
	BasePublication
	name        string
	editionType string
	volume      int
	issueNumber int
	theme       string
	coverStory  string
	issn        string
	circulation int
	sections    []string
	price       string
}

func NewNewspaper(id, name, theme string, issueNumber int) Newspaper {
	return Newspaper{
		BasePublication: BasePublication{
			id:              id,
			title:           name,
			publisher:       "Демо-библиотека",
			publicationDate: "2026-05-05",
			language:        "ru",
			pageCount:       24,
			format:          "tabloid",
		},
		name:        name,
		editionType: "daily",
		volume:      1,
		issueNumber: issueNumber,
		theme:       theme,
		coverStory:  "Главная тема выпуска",
		price:       "50 RUB",
	}
}

func (n Newspaper) Summary() string {
	return fmt.Sprintf("%s — тема: %s, выпуск #%d", n.name, n.theme, n.issueNumber)
}
